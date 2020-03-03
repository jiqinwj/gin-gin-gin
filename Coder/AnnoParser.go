package Coder

import (
	"go/ast"
	"gin-gin-gin/Coder/annotations"
	"reflect"
)

type Annotation interface {
	Parse(comment string)   //解析方法
	IsHit() bool   //是否命中
	GetAnno() interface{}
	GetProp(name string) interface{}
	SetProp(name string,v interface{})
	Invoke(f interface{},index int)
}
type Annotations []Annotation

var (
	AnnoTargets=map[string]Annotation{  //初始值
		"gen":annotations.NewGen(),
		"router":annotations.NewRouter(), //一个新注解,处理路由相关
	}
)

//从初始 map里面 获取 注解
func(this Annotations) getFromMap(c string) Annotation  {
	if anno:=AnnoTargets[c];anno!=nil{
		return  anno
	}
	return nil
}

func(this Annotations) Get(c string) []Annotation{  //这里做了改动,返回的是切片。支持一个方法打多个注解
	anno:=this.getFromMap(c)
	if anno==nil{
		return nil
	}
	ret:=make(Annotations,0)
	for _,item:=range this {
		if reflect.TypeOf(item).Elem() == reflect.TypeOf(anno).Elem(){
			ret=append(ret,item)
		}
	}
	return ret
}
//是否存在注解
func(this Annotations) Contains(c string ) bool{
   anno:=this.Get(c)
   if anno==nil{
     return false
   }
  return true
}

//主函数 解析函数
func Parse(commentGroup *ast.CommentGroup) Annotations  {
	if commentGroup==nil{
		return nil
	}
	ret:=make(Annotations,0)
	tags:=[]string{}
	for _,c:=range commentGroup.List{ //对每一行注释 进行解析

		tags= parse_comment(c.Text,tags,&ret)
	}

	return ret
}
//判断是否存在 切片中
func  in_array(str string,arr []string) bool {
	for _, s := range arr {
		if str == s {
			return true
		}

	}
	return false
}

//复制 注解对象
func copyAnno(anno Annotation) Annotation {
	vt:=reflect.TypeOf(anno).Elem()
	newObj:=reflect.New(vt)
	newObj.Elem().Set(reflect.ValueOf(anno).Elem())
	return newObj.Interface().(Annotation)
}
func parse_comment(c string,tags []string,ret *Annotations) []string  {
	for key,anno:=range AnnoTargets{
		//if in_array(key,tags){ //相同的注解 只解析一次 /.目前注释掉 支持 一个方法生成多个
		//	continue
		//}

		cpAnno:=copyAnno(anno)
		cpAnno.Parse(c)
		if cpAnno.IsHit(){
			tags=append(tags,key)
			*ret=append(*ret,cpAnno)
		}
	}
	return tags

}
