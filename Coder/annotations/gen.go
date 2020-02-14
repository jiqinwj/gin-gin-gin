package annotations

import (
	"fmt"
	"go/ast"
	"strings"
)

type Gen struct {//注解 作用是：判断是否要生成代码
	Str string //注释中 字符串值，用来做正则判断
	Comment string  //注释内容
	Hit bool //是否命中
	Anno *struct{
		Order int
		Id string
		Params []string
	}
	AbstractAnnotation
}

func NewGen() *Gen {
	return &Gen{Str:"@gen",Hit:false}
}
//解析方法
func(this *Gen) Parse(comment string)  {
	this.Comment=comment
	this.Anno=&struct {
		Order int
		Id string
		Params []string
	}{Params:make([]string,0)}
	err:=CommentToAnno(this.Comment,this.Str,this.Anno)
	if err!=nil {
		this.Hit=false
	}else {
		this.Hit=true
	}
}
//是否命中
func(this *Gen) IsHit() bool  {
	return this.Hit
}
func(this *Gen) GetAnno() interface{} {
	return this.Anno
}
func(this *Gen)Invoke (f interface{},findex int)  {
	fn:=f.(*ast.FuncDecl)
	pkgName:=this.GetProp("PkgName") //这里是包名，很重要
	rList:= GetResultType(fn.Type.Results.List) //2、""返回值变量
	code:=""
	varName:="" //保存 返回值变量
	isMutilRet:=false //是否多值返回
	for index,r:=range rList{
		if findex>0{ //由于支持了 一个函数多注解生成，因此需要根据index区分 变量值
			varName+=fmt.Sprintf(" %s_%d",UcFirst(GetVarName(r,fn.Name.String())),findex)
		}else {
			varName+=fmt.Sprintf(" %s",UcFirst(GetVarName(r,fn.Name.String())))
		}
		if index!=len(rList)-1{
			varName+=fmt.Sprintf(",")
			isMutilRet=true
		}
	}//以上把返回值变量拼凑完毕了

	code+=varName
	if !isMutilRet {
		this.VarName=strings.Split(varName,",")[0]  //设置返回值变量 ,,对于多值返回 目前不支持。只支持单值返回
	}else{
		this.VarName=varName
	}



	if len(this.Anno.Params)>0 { //有指定参数
		code+=fmt.Sprintf("=%s.%s(%s)\n",pkgName,fn.Name,GetExistsParams(this.Anno.Params))
	}else {
		code+=fmt.Sprintf("=%s.%s(%s)\n",pkgName,fn.Name,GetParamsType(fn.Type.Params.List,pkgName.(string)))
	}

 	this.SetProp("Code",code)
}

