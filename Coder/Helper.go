package Coder

import (
	"fmt"
	"go/ast"
	"strings"
)

var (
	zeroType=map[string]string{
		"int":"0",
		"string":"\"\"",
	}
)


////是否需要生成
//func ShouldGen(comment string) bool  {
//	reg,_:=regexp.Compile(`^@gen\(\)`)
//	return reg.MatchString(comment)
//}
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

//处理函数参数的生成
func GetParamsType(p []*ast.Field) string  {
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
		  }else{
			  //ret=append(ret,"nil")
			  //做了改动，如果前面没有包名，也是Ident
			  ret=append(ret,fmt.Sprintf("&%s{}",p.X))
		  }

	  }
  }
  return strings.Join(ret,",")
}
