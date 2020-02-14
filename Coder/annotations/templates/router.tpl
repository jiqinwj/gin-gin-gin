func GetAutoRouter() *gin.Engine{
    logFile,err:= os.OpenFile("gin-log.log",os.O_RDWR|os.O_CREATE|os.O_TRUNC,0777)
	gin.DefaultWriter=io.MultiWriter(logFile)
	if err!=nil{
		log.Fatal("日志文件创建失败",err)
	}
   router:=gin.Default()
   {{range .Groups }}
     {{.}}:=router.Group("{{.}}")
   {{ end }}
    {{range .Cods }}
     {{.}}
   {{ end }}
   return router
}


