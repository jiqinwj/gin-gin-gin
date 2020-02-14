package annotations

import (
	"fmt"
	"go/ast"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)
var CodeDir string
func init()  {
	dir,err:=os.Getwd()
	if err!=nil{
		log.Fatal(err)
	}
	CodeDir=strings.Replace(dir,"\\","/",-1)+"/Coder"
}
//根据分组名 把键值映射到Map中
func MapGroup(list []string,names []string)   map[string]interface{}{
	ret:=make(map[string]interface{})

	if len(list) <=1 {
		return nil
	}
	for index,name:=range names{
		if index==0 || name==""{
			continue
		}
		ret[name]=ExtractValue(list[index])
	}
	return ret
}
//判断value是string还是int
func ExtractValue(str string) interface{}  {
	if regexp.MustCompile(`^\d+$`).MatchString(str){
		i,_:= strconv.Atoi(str)
		return i
	}
	if regexp.MustCompile(`^".*?"$`).MatchString(str){
		return  strings.Trim(str,"\"")
	}
	if regexp.MustCompile(`^{.*}$`).MatchString(str){
		 strList:=[]byte(str)
		 return string(strList[1:len(strList)-1])
	}
	return str
}

func CommentToAnno(comment string,tag string,v interface{}) error {
	pattern:=tag+`\((?P<params>.*?)\)`
	reg,err:= regexp.Compile(pattern)
	if err!=nil{
		return err
	}
	if !reg.MatchString(comment){
		return fmt.Errorf("not match")
	}

	ret:=MapGroup(reg.FindStringSubmatch(comment),reg.SubexpNames())

	if ret==nil || strings.Trim(ret["params"].(string)," ")==""{ //如果没有参数 譬如@gen()
		return nil
	}
	param_patter:=`(?P<pName>[a-zA-Z]+)=(?P<pValue>(\"(.+?)\")|(\d+)|(\{.*\}))`
	reg,_= regexp.Compile(param_patter)

	if !reg.MatchString(ret["params"].(string)){  //参数格式不对
		return fmt.Errorf("params error")
	}
	lists:=reg.FindAllStringSubmatch(ret["params"].(string),-1)
	vv:=reflect.ValueOf(v).Elem()
	vt:=reflect.TypeOf(v).Elem()

	for _,list:=range lists{
		mp:=MapGroup(list,reg.SubexpNames())
		for i:=0;i<vv.NumField();i++{
			if strings.ToLower(vt.Field(i).Name)==strings.ToLower(mp["pName"].(string)){
				if vv.Field(i).Kind().String() == "slice" { //切片
					list:=strings.Split(mp["pValue"].(string),",")
					vv.FieldByName(vt.Field(i).Name).Set(reflect.ValueOf(list))
				}else {
					vv.FieldByName(vt.Field(i).Name).Set(reflect.ValueOf(mp["pValue"]))
				}
			}
		}
	}
	return nil
}

var (
	zeroType=map[string]string{
		"int":"0",
		"string":"\"\"",
	}
)
//首字母大写
func UcFirst(str string) string{
	var ret string
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {
				vv[i] -= 32
				ret += string(vv[i])
			} else {
				return str
			}
		} else {
			ret += string(vv[i])
		}
	}
	return ret
}
//获取返回值类型，是一个切片
func GetResultType(list []*ast.Field) []string  {
	ret:=make([]string,0)
	for _,f:=range list{
		//fmt.Printf("%T\n",f.Type)
		if nf,ok:=f.Type.(*ast.SelectorExpr);ok{
			ret=append(ret,fmt.Sprintf("%s.%s",nf.X,nf.Sel))
		}
		if nf,ok:=f.Type.(*ast.Ident);ok{
			ret=append(ret,fmt.Sprintf("%s",nf.Name))
		}
		if _,ok:=f.Type.(*ast.FuncType);ok{
			ret=append(ret,fmt.Sprintf("%s","fun"))
		}
	}
	return ret
}

//优化变量名
func GetVarName(v string,fn string) string   {
	fn=strings.ToLower(fn)
	vList:=strings.Split(v,".")
	if len(vList)>1 {
		return fn+"_"+strings.ToLower(vList[len(vList)-1])
	}
	return fn+"_"+strings.ToLower(v)
}

//拼接指定参数
func GetExistsParams(params []string) string {
	ret:=[]string{}
	idreg:=regexp.MustCompile(`^id:(\w+)`)
	for _,p:=range params{
		if idreg.MatchString(p){  //代表这个参数 是某个注解的返回值
			ret=append(ret,"#{"+idreg.FindStringSubmatch(p)[1]+"}")
		}else {
			ret=append(ret,p)
		}
	}
	return strings.Join(ret,",")
}
//处理函数参数的生成
func GetParamsType(p []*ast.Field,pkg string) string  {
	ret:=make([]string,0)
	for _,param:=range p{
		if pp,ok:=param.Type.(*ast.Ident);ok{  //譬如string int
			if getp,ok:=zeroType[pp.Name];ok{
				ret=append(ret,getp)
			}else{
				ret=append(ret,"nil")
			}
		}
		if p,ok:=param.Type.(*ast.SelectorExpr);ok{

			ret=append(ret,fmt.Sprintf("%s.%s{}",p.X,p.Sel))
		}
		if p,ok:=param.Type.(*ast.StarExpr);ok{

			if pf,ok:=p.X.(*ast.SelectorExpr);ok{
				ret=append(ret,fmt.Sprintf("&%s.%s{}",pf.X,pf.Sel))
			}else if pp,ok:=p.X.(*ast.Ident);ok{//做了改动，如果前面没有包名，也是Ident
				if getp,ok:=zeroType[pp.Name];ok{
					ret=append(ret,getp)
				}else{
					ret=append(ret,fmt.Sprintf("&%s.%s{}",pkg,p.X))
				}

			}

		}
	}
	return strings.Join(ret,",")
}