package annotations

import (
	"fmt"
)

type Router struct {//注解 作用是：判断是否要生成代码
	Str string //注释中 字符串值，用来做正则判断
	Comment string  //注释内容
	Hit bool //是否命中
	Anno *struct{
		Method string
		Uri string
		Handler string
		Group string
	}
	AbstractAnnotation
}
func NewRouter() *Router {
	return &Router{Str:"@router",Hit:false}
}
//解析方法
func(this *Router) Parse(comment string)  {
	this.Comment=comment
	this.Anno=&struct {
		Method string
		Uri string
		Handler string
		Group string
	}{ }
	err:=CommentToAnno(this.Comment,this.Str,this.Anno)
	if err!=nil {
		this.Hit=false
	}else {
		this.Hit=true
	}
}
//是否命中
func(this *Router) IsHit() bool  {
	return this.Hit
}
func(this *Router) GetAnno() interface{} {
	return this.Anno
}
func(this *Router)Invoke (f interface{},findex int)  {
	code:=fmt.Sprintf(`%s.Handle("%s","%s",#{%s})`,
		this.Anno.Group,this.Anno.Method,
		this.Anno.Uri,this.Anno.Handler)
	this.SetProp("Code",code)

}