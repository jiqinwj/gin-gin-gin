package Services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gin-gin-gin/App"
)

func CreateBookListRequest() App.EncodeRequestFunc{
	return func(context *gin.Context) (i interface{}, e error) {
		bReq:=&BookListRequest{}
		err:=context.BindQuery(bReq)
		fmt.Println(bReq)
		if err!=nil{
			return nil,err
		}
		return bReq,nil
	}
}
func CreateBookListResponse()  App.DecodeResponseFunc  {
	return func(context *gin.Context, i interface{}) error {
		res:=i.(*BookListResponse)
		context.JSON(200,res)
		return nil
	}
}