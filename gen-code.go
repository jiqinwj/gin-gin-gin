package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"gin-gin-gin/Coder"
	"strings"
)


func main()  {
	getPath:=""
	flag.StringVar(&getPath, "p", "", "代码文件所在路径")
	flag.Parse()
	if getPath==""{
		log.Fatal("文件路径不可为空")
	}
	getPath=strings.Replace(getPath,"\\","/",-1)
	files:=Coder.GetGoFiles(getPath)


	buildFile,err:=os.OpenFile("./autocode/service.auto.go",os.O_CREATE|os.O_RDWR|os.O_TRUNC,0666)
	if err!=nil {
		log.Fatal("生成代码文件失败:",err)
	}
	 fmt.Fprint(buildFile,"package autocode"+"\r\r") //写入包名

	annoList:=Coder.Annotations{}
	for _,file:=range files{
		ret:=Coder.GenCode(file)  //ret Annotations
		annoList=append(annoList,ret...)
	}

	Coder.WriteImport(buildFile,getPath)
	genAnnos:=annoList.Get("gen") //单独抽出 gen类型注解
	Coder.SortGen(&genAnnos) //排序
	Coder.WriteCode(buildFile,genAnnos) //代码写入文件

	routerAnnos:=annoList.Get("router") //单独抽出路由注解
	Coder.WriteCodeByRouter(buildFile,routerAnnos,genAnnos) //代码写入文件
}
