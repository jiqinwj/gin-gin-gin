package main

import (
	"gin-gin-gin/App"
	"gin-gin-gin/App/Services"
	. "gin-gin-gin/AppInit"
	"github.com/gin-gonic/gin"
)

func main() {

	router:=gin.Default()
	v1:=router.Group("v1")
	{
		//v1.Handle(HTTP_METHOD_GET,"/prods", func(context *gin.Context) {
		//		prods:=Models.BookList{}
		//		GetDB().Limit(10).Order("book_id desc").Find(&prods)
		//		context.JSON(200,prods)
		//})

		bs:=&Services.BookService{}
		bookListHandler:=App.RegisterHandler(Services.BookListEndPoint(bs),Services.CreateBookListRequest(),
			Services.CreateBookListResponse())
		v1.Handle(HTTP_METHOD_GET,"/prods",bookListHandler)


	}
	router.Run(SERVER_ADDRESS)
}
