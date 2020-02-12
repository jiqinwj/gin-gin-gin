package Models

import "time"

const (
	CacheType_RankList=1 //排行榜
)
type BooksCache struct{
	ItemID int `gorm:"column:item_id;AUTO_INCREMENT;PRIMARY_KEY"`
	CacheContent string `gorm:"column:cache_content;type:json"`
	CacheType int `gorm:"column:cache_type;type:int"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime"`
}