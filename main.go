package main

import (
	"fmt"
	"gin-gin-gin/App"
	"gin-gin-gin/App/Services"
	. "gin-gin-gin/AppInit"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//logFile,err:= os.OpenFile("gin-log.log",os.O_CREATE|os.O_APPEND,0666)
	//gin.DefaultWriter=io.MultiWriter(logFile)
	//
	//if err!=nil{
	//	log.Fatal("日志文件创建失败",err)
	//}


	file, _ := os.Create("access.log")
	//gin.DefaultWriter = file
	gin.DefaultWriter = io.MultiWriter(file)// 效果是一样的
	router:=gin.Default()
	v1:=router.Group("v1")
	{
		bookService_List_Endpoint:=Services.BookListEndPoint(&Services.BookService{})//图书列表endpoint
		bookService_Detail_Endpoint:=Services.BookDetailEndPoint(&Services.BookService{})//图书详细endpoint
		bookService_Fav_Endpoint:=Services.BookFavEndPoint(&Services.BookService{}) //收藏图书
		bookResponseFunc:=Services.CreateBookResponse()
		bookListHandler:=App.RegisterHandler(bookService_List_Endpoint,//业务最终函数
			Services.CreateBookListRequest(),//怎么取参数
			bookResponseFunc, //怎么处理响应
			)
		bookDetailHandler:=App.RegisterHandler(bookService_Detail_Endpoint,//业务最终函数
			Services.CreateBookDetailRequest(),//怎么取参数
			bookResponseFunc, //怎么处理响应
		)
		bookFavHandler:=App.RegisterHandler(bookService_Fav_Endpoint,//业务最终函数
			Services.CreateBookFavRequest(),//怎么取参数
			bookResponseFunc, //怎么处理响应
		)


		v1.Handle(HTTP_METHOD_GET,"/prods",bookListHandler)
		v1.Handle(HTTP_METHOD_GET,"/prods/:id", bookDetailHandler)
		v1.Handle(HTTP_METHOD_POST,"/prods/fav", bookFavHandler)
	}

	errChan:=make(chan error)

	go func() { //启动http server
		err:=router.Run(SERVER_ADDRESS)
		if err!=nil{
			errChan<-err
		}
	}()
	//go func() {
	//	Jobs.MyCron.Start() //启动定时任务
	//}()
	go func() {
		sig_c:=make(chan os.Signal)
		signal.Notify(sig_c,syscall.SIGINT,syscall.SIGTERM)
		errChan<-fmt.Errorf("%s",<-sig_c)
	}()

	getErr:=<-errChan
	log.Fatal(getErr)





}
