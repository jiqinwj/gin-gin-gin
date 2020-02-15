package RedisUtil

import (
	"encoding/json"
)

func IsBool(input interface{}) bool  {
	if _,ok:=input.(bool);ok{
		return true
	}
	return false
}
// json_encode
func JsonEncode(v interface{}) interface{}  {
	  b,err:=json.Marshal(v)
	  if err!=nil{
	  	return false
	  }
	  return string(b)//json字符串
}
//json_decode
func JsonDecode(data []byte,object interface{}) bool  {
	 err:=json.Unmarshal(data,object)
	 if err!=nil{
	 	return false
	 }
	 return true
}
