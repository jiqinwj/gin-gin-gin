package Services

import (
	"context"
	"fmt"
	"gin-gin-gin/App"
	"gin-gin-gin/App/RedisUtil"
	"gin-gin-gin/AppInit"
)

func BookDetailCache() App.MiddleWare  {  //缓存中间件
	return func (next App.Endpoint) App.Endpoint{
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			req:=request.(*BookDetailRequest)
			//假设缓存key = books123445
			cacheKey:=fmt.Sprintf("%s%d","books",req.BookID)
			redConn:=AppInit.RedisDefaultPool.Get()
			defer redConn.Close()
			 strData:=RedisUtil.NewStringData()
			//  strRet:=<-strData.Get(cacheKey) //上节课代码
			return strData.Get(cacheKey).Then(func(rsp []byte) (interface{},error) {
				bookrsp:=&BookResponse{}
			    if RedisUtil.JsonDecode(rsp,bookrsp){
					return bookrsp,nil
				}
				return nil,nil
			}, func() (interface{},error) {  //缓存中没有
				rsp,err:=next(ctx,request)//从数据库取

				if rsp!=nil && err==nil{
					//jsonRet:=RedisUtil.JsonEncode(rsp)
					//if _,ok:=jsonRet.(string);ok{
						//redConn.Do("setex",cacheKey,20,jsonStr) //插入缓存
						strData.Set(cacheKey,rsp,200,false)
					//}
				}
				return rsp,err
			})


		}
	}
}
