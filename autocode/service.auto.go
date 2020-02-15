package autocode

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"

	"gin-gin-gin/App"

	"gin-gin-gin/App/Services"
)
var (
 Registerhandler_fun=App.RegisterHandler( Booklistendpoint_endpoint, Createbooklistrequest_encoderequestfunc, Createbookresponse_decoderesponsefunc)
 Registerhandler_fun_1=App.RegisterHandler( Bookdetailendpoint_endpoint, Createbookdetailrequest_encoderequestfunc, Createbookresponse_decoderesponsefunc)
 Registerhandler_fun_2=App.RegisterHandler( Bookfavendpoint_endpoint, Createbookfavrequest_encoderequestfunc, Createbookresponse_decoderesponsefunc)
 Createbooklistrequest_encoderequestfunc=Services.CreateBookListRequest()
 Createbookdetailrequest_encoderequestfunc=Services.CreateBookDetailRequest()
 Createbookfavrequest_encoderequestfunc=Services.CreateBookFavRequest()
 Booklistendpoint_endpoint=Services.BookListEndPoint(&Services.BookService{})

 Bookdetailendpoint_endpoint=Services.BookDetailCache()(Services.BookDetailEndPoint(&Services.BookService{}))

 Bookfavendpoint_endpoint=Services.BookFavEndPoint(&Services.BookService{})
 Createbookresponse_decoderesponsefunc=Services.CreateBookResponse()
 Test_string=Services.Test("",0,"")

)
func GetAutoRouter() *gin.Engine{
    //logFile,err:= os.OpenFile("gin-log.log",os.O_RDWR|os.O_CREATE|os.O_TRUNC,0777)
	//gin.DefaultWriter=io.MultiWriter(logFile)

	file, _ := os.Create("access.log")
	//gin.DefaultWriter = file
	gin.DefaultWriter = io.MultiWriter(file)// 效果是一样的

	//if err!=nil{
	//	log.Fatal("日志文件创建失败",err)
	//}
   router:=gin.Default()
   
     v1:=router.Group("v1")
   
    
     v1.Handle("GET","/prods", Registerhandler_fun)
   
     v1.Handle("GET","/prods/:id", Registerhandler_fun_1)
   
     v1.Handle("POST","/prods/fav", Registerhandler_fun_2)
   
   return router
}


