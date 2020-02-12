package Services

import (
	"github.com/gin-gonic/gin"
	"gin-gin-gin/App"
)
//图书列表 请求参数获取
func CreateBookListRequest() App.EncodeRequestFunc{
	return func(context *gin.Context) (i interface{}, e error) {
		bReq:=&BookListRequest{}
		err:=context.ShouldBindQuery(bReq) //和框架有关   /v1/books?size=100
		if err!=nil{
			return nil,err
		}
		return bReq,nil
	}
}
//加载图书详细请求函数
func CreateBookDetailRequest() App.EncodeRequestFunc{
	return func(context *gin.Context) (i interface{}, e error) {
		bReq:=&BookDetailRequest{}
		err:=context.ShouldBindUri(bReq)
		if err!=nil{
			return nil,err
		}
		return bReq,nil
	}
}

func CreateBookFavRequest() App.EncodeRequestFunc{
	return func(context *gin.Context) (i interface{}, e error) {
		bReq:=&BookMetaRequest{}
		err:=context.ShouldBindJSON(bReq)
		if err!=nil{
			return nil,err
		}
		bReq.Type="fav"
		return bReq,nil
	}
}


func CreateBookResponse()  App.DecodeResponseFunc  {
	return func(context *gin.Context, res interface{}) error {
		context.JSON(200,res)
		return nil
	}
}

