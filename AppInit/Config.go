package AppInit

import (
	"github.com/go-ini/ini"
	"log"
)

const (
	HTTP_METHOD_GET="GET"
	HTTP_METHOD_POST="POST"

)
var(
	SERVER_ADDRESS=":8080"
	MYSQL_DSN=""
	MYSQL_MAXIDLE=10
	MYSQL_MAXCONN=50
)


func init()  {
	cfg, err := ini.Load("gin.ini")
	if err != nil {
		log.Fatal("config err")
	}
	//muststring 带默认值
	SERVER_ADDRESS=cfg.Section("server").Key("port").MustString(":8080")
	//数据库连接字符串
	MYSQL_DSN=cfg.Section("mysql").Key("dsn").String()
	//最大空闲连接
	MYSQL_MAXIDLE=cfg.Section("mysql").Key("maxidle").MustInt(10)
	//最大连接数
	MYSQL_MAXCONN=cfg.Section("mysql").Key("maxconn").MustInt(50)
}