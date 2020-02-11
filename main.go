package main

import (
	"github.com/gin-gonic/gin"
	. "gin-gin-gin/AppInit"
	"gin-gin-gin/Models"
)

func main() {

	router:=gin.Default()
	v1:=router.Group("v1")
	{
		v1.Handle(HTTP_METHOD_GET,"/prods", func(context *gin.Context) {
				prods:=Models.BookList{}
				GetDB().Limit(10).Order("book_id desc").Find(&prods)
				context.JSON(200,prods)
		})
	}
	router.Run(SERVER_ADDRESS)
}
