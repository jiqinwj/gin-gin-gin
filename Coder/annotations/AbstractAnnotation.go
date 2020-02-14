package annotations

import (
	"reflect"
)

type AbstractAnnotation struct {
   Code string
   VarName string
   PkgName string //包名
}
func(this *AbstractAnnotation) SetProp(name string,v interface{})  {
	vv:=reflect.ValueOf(this).Elem()
	f:=vv.FieldByName(name)
	f.Set(reflect.ValueOf(v))
}
func(this *AbstractAnnotation) GetProp(name string) interface{}  {
	vv:=reflect.ValueOf(this).Elem()
	f:=vv.FieldByName(name)
	return f.Interface()
}