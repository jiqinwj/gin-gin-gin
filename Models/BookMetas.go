package Models

import (
	"github.com/jinzhu/gorm"
	"gin-gin-gin/AppInit"
)

type BookMetas struct{
	MetaId int `gorm:"column:meta_id;AUTO_INCREMENT;PRIMARY_KEY"`
	MetaKey string `gorm:"column:meta_key;type:varchar(200)"`
	MetaValue string `gorm:"column:meta_value;type:text"`
	ItemId int `gorm:"column:item_id;type:int"`
}
func NewBookMeta(key string,value string,itemid int) *BookMetas  {
	return &BookMetas{MetaKey:key,MetaValue:value,ItemId:itemid}
}
func(this *BookMetas) Save(db ...*gorm.DB) error {
	if len(db)>0{
		return db[0].Set("gorm:insert_option",
			"ON DUPLICATE KEY UPDATE meta_value=meta_value+1").Create(this).Error
	}
	return AppInit.GetDB().Set("gorm:insert_option",
		"ON DUPLICATE KEY UPDATE meta_value=meta_value+1").Create(this).Error
}
