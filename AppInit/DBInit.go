package AppInit

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql",
		MYSQL_DSN)
	if err != nil {
		log.Fatal(err)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(MYSQL_MAXIDLE)
	db.DB().SetMaxOpenConns(MYSQL_MAXCONN)
	fmt.Println("当前连接池",MYSQL_MAXCONN,MYSQL_MAXIDLE)
	//
	//db.LogMode(true)
}
func  GetDB() *gorm.DB {
	return db
}

