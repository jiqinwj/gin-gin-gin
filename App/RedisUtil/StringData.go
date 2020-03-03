package RedisUtil

import (
	"fmt"
	"gin-gin-gin/AppInit"
	"github.com/gomodule/redigo/redis"
	"log"
)
type StringResult chan []byte
//两个参数，第一个是有值的情况  第二个是没有值
func(this StringResult) Then(resolve func(rsp []byte) (interface{},error),reject func() (interface{},error)) (interface{},error){
    ret:=<-this
    if ret!=nil{
		return resolve(ret)
	}else {
		return reject()
	}
}
type StringData struct {}
func NewStringData() *StringData  {
	return &StringData{}
}
func(this *StringData) Get(key string) StringResult{
	c:=make(chan []byte)
	redConn:=AppInit.RedisDefaultPool.Get()
	go func() {
		defer redConn.Close()
		b,err:=redis.Bytes(redConn.Do("get",key))
		//到这一步咋办？
		if err!=nil{
			log.Println(err)
			c<-nil
		}else{
			c<-b
		}
	}()
	return c
}

//set 封装
func(this *StringData) Set(key string,v interface{},ex int,json bool) bool{
	conn:=AppInit.RedisDefaultPool.Get()
	defer conn.Close()
	var reply interface{}
	replyStr:=JsonEncode(v)
	var err error
	if ex>0{

		reply,err=conn.Do("setex",key,ex,replyStr)


	}else{
		reply,err=conn.Do("set",key,replyStr)

	}

	fmt.Printf("%T",reply)

	ret,err:=redis.String(reply,err)
	//fmt.Print("dddsss")
	//fmt.Println(boolRet)


	if ret=="OK" && err==nil {
		return true
	}

	return false

}
