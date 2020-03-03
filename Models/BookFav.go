package Models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"gin-gin-gin/AppInit"
	"time"
)

type BookFav struct {
	ItemID int `gorm:"column:item_id;AUTO_INCREMENT;PRIMARY_KEY"`
	BookID int `gorm:"column:book_id;type:int"`
	UserID int `gorm:"column:user_id;type:int"`
	AddTime time.Time `gorm:"column:add_time;type:datetime"`
}

func NewBookFav(bookid int,userid int) *BookFav  {
	return &BookFav{BookID:bookid,UserID:userid,AddTime:time.Now()}
}
func(this *BookFav) Save(db ...*gorm.DB)  error {
	if len(db)>0{
		  ret:=db[0].Create(this)
		  if ret.Error!=nil || ret.RowsAffected!=1{
		  	return fmt.Errorf("error fav")
		  }else {
		  	return nil
		  }
	}
	return AppInit.GetDB().Create(this).Error
}