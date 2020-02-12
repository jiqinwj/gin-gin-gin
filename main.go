package main

import (
	"github.com/gin-gonic/gin"
	"gin-gin-gin/App"
	"gin-gin-gin/App/Services"

	. "gin-gin-gin/AppInit"
)

func main() {

	router:=gin.Default()
	v1:=router.Group("v1")
	{
		//v1.Handle(HTTP_METHOD_GET,"/prods", func(context *gin.Context) {
		//		//	prods:=Models.BookList{}
		//		//	GetDB().Limit(10).Order("book_id desc").Find(&prods)
		//		//	context.JSON(200,prods)
		//		//})



		bookService_List_Endpoint:=Services.BookListEndPoint(&Services.BookService{})//图书endpoint
		bookService_Detail_Endpoint:=Services.BookDetailEndPoint(&Services.BookService{})//图书endpoint
		bookResponseFunc:=Services.CreateBookResponse()
		bookListHandler:=App.RegisterHandler(bookService_List_Endpoint,//业务最终函数
			Services.CreateBookListRequest(),//怎么取参数
			bookResponseFunc, //怎么处理响应
		)
		bookDetailHandler:=App.RegisterHandler(bookService_Detail_Endpoint,//业务最终函数
			Services.CreateBookDetailRequest(),//怎么取参数
			bookResponseFunc, //怎么处理响应
		)

		v1.Handle(HTTP_METHOD_GET,"/prods",bookListHandler)
		v1.Handle(HTTP_METHOD_GET,"/prods/:id", bookDetailHandler)
	}
	router.Run(SERVER_ADDRESS)
}
