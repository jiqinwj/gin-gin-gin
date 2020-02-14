package Coder

import (
	"bytes"
	"fmt"
	"github.com/deckarep/golang-set"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"gin-gin-gin/Coder/annotations"
	"regexp"
	"sort"
	"strings"
	"text/template"
)


//根据 go.mod获取module名称
func GetModuleName() string{
	dir, _ := os.Getwd()
	mod_dir:=strings.Replace(dir,"\\","/",-1)+"/go.mod"
	mod_content,err:=ioutil.ReadFile(mod_dir)
	if err!=nil{
		log.Fatal(err)
	}
	pattern:=`module\s*(.*?)\n`
	ret:=regexp.MustCompile(pattern).FindStringSubmatch(string(mod_content))
	if len(ret)!=2{
		log.Fatal("go.mod设置不正确")
	}
	return ret[1]
}
func init()  {

	//这里要做的事:在当前目录下生成一个autocode目录
	dir, _ := os.Getwd()
	dir=strings.Replace(dir,"\\","/",-1)+"/autocode"
	_, err := os.Stat(dir)
	if err!=nil && os.IsNotExist(err){
		err:=os.Mkdir(dir,0666)
		if err!=nil{
			log.Fatal("初始化文件夹失败")
		}
	}
}



//不使用filepath的walk方法 递归读取文件夹 .如果是为了获取文件夹名
func WalkDir(dir string,dset mapset.Set,prefix string){

	dset.Add(prefix+strings.Replace(dir,"./","/",-1))
	files, err := ioutil.ReadDir(dir)//读取目录下文件
	if err != nil{
		return
	}
	for _, file := range files{
		if file.IsDir(){
			WalkDir(dir + "/" + file.Name(),dset,prefix)
			continue
		}
	}
}

//递归获取所有文件，
func GetGoFiles(fpath string) []string {
	ret:=make([]string,0)
	 _= filepath.Walk(fpath,
		func(file string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() || path.Ext(info.Name())!=".go" ||  info.Name()=="service.auto.go" {
				return nil
			}
			file=strings.Replace(file,"\\","/",-1)
			ret=append(ret,file)
			return nil
		})
	 return ret
}

func getGenAnnoByID(id string,annos Annotations) Annotation{
	for _,anno:=range annos{
		if anno.(*annotations.Gen).Anno.Id==id{
			return anno
		}
	}
	return nil
}
//这一步的作用是把 ID参数给替换掉
func SetParamByID(code *string ,annos Annotations){
	idreg:=regexp.MustCompile(`#{(.*?)}`)
	allMatches:=idreg.FindAllStringSubmatch(*code,-1)
	for _,match:=range allMatches{
		for index,item:=range match{
			if index==0{
				continue
			}
			getAnno:=getGenAnnoByID(item,annos)
			pVar:="nil"
			if getAnno!=nil{
				pVar=getAnno.GetProp("VarName").(string)
			}
			*code=strings.Replace(*code,`#{`+item+`}`,pVar,-1)
		}
	}
}


//排序
func SortGen(gens *[]Annotation)  {
	sort.SliceStable(*gens, func(i, j int) bool {  //这一步做的是排序
		anno1:=(*gens)[i].(*annotations.Gen)
		anno2:=(*gens)[j].(*annotations.Gen)
		return anno1.Anno.Order<anno2.Anno.Order
	})
}

//解析模板，临时函数
func parseTpl(tplName string,data  interface{}) string  {
	tplContent,err:=ioutil.ReadFile(annotations.CodeDir+"/annotations/templates/"+tplName+".tpl")
	if err!=nil {
		log.Println(err)
		return ""
	}
	tmpl, _ := template.New(tplName).Parse(string(tplContent))
	buf := bytes.Buffer{}
	err=tmpl.Execute(&buf,data)
	if err!=nil{
		log.Println(err)
		return ""
	}
	return buf.String()
}
func WriteImport(f *os.File,dir string)  {
	dset:=mapset.NewSet()
	 WalkDir(dir,dset,GetModuleName())
	 cnt:=parseTpl("import",dset.ToSlice())
	fmt.Fprint(f,cnt)
}
func WriteCode(f *os.File,genAnnos []Annotation){
	cods:=[]string{}
	for _,g:=range genAnnos{//到这一步需要 替换参数
		code:=g.GetProp("Code").(string)
	    SetParamByID(&code,genAnnos)
		cods=append(cods,code)
	}
	cnt:=parseTpl("autocode",cods)
	fmt.Fprint(f,cnt)

}
//路由注解的写入代码函数
func WriteCodeByRouter(f *os.File,annSet []Annotation, genAnnos []Annotation){
	cods:=[]string{}
	groups:=mapset.NewSet()
	for _,g:=range annSet{//到这一步需要 替换参数
	    route:=g.(*annotations.Router)
		code:=route.Code
		groups.Add(route.Anno.Group) //保存分组
		SetParamByID(&code,genAnnos)
		cods=append(cods,code)
	}
	data:= struct {
		Cods []string
		Groups []interface{}
	}{cods,groups.ToSlice()}
	cnt:=parseTpl("router",data)
	fmt.Fprint(f,cnt)
}

func GenCode(file string)  Annotations  {
	fset:=token.NewFileSet()
	ast_file,err:=parser.ParseFile(fset,file,nil,0 | parser.ParseComments )
	if err!=nil{
		log.Fatal(err)
	}
	 ret:=Annotations{}
	for _,decl:=range ast_file.Decls {
		if fn,ok:=decl.(*ast.FuncDecl);ok{
			annos:=Parse(fn.Doc)
			if annos.Contains("gen") && fn.Type.Results!=nil{  //1、判断注解中 是否需要进行生成 2 //必须有返回值，不然没有意义去生成
					genList:=annos.Get("gen") //这里变成切片了
					for index,gen:=range genList{
						gen.SetProp("PkgName",ast_file.Name.String()) //这一步把package灌入
						gen.Invoke(fn,index) //生成代码
						ret=append(ret,gen)
					}
					routerList:=annos.Get("router") //这里是处理路由
					for index,router:=range routerList{
						router.SetProp("PkgName",ast_file.Name.String()) //这一步把package灌入
						router.Invoke(fn,index) //生成代码
						ret=append(ret,router)
					}
			}
		}
	}

	return ret
}