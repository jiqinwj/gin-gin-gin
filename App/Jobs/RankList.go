package Jobs

import (
	"encoding/json"
	"log"
	"gin-gin-gin/AppInit"
	"gin-gin-gin/Models"
	"time"
)

//排行榜 定时任务
func init()  {
	_,err:=MyCron.AddFunc("0/6 * * * * *", func() {
		prods:=&Models.BookList{}
		sql:=`
select book_id,book_name,book_kind from (
select  book_id,book_name, meta_value,book_kind,IF(@pre=book_kind,@rownum:=@rownum+1,@rownum:=1) as rownum,@pre:=book_kind 
from (select book_kind,b.book_id,book_name,a.meta_value from book_metas a, books b where a.item_id=b.book_id AND
 meta_key='click' ORDER  by 
b.book_kind ,meta_value desc 
    ) a,(select @pre:='',@rownum:=0 ) b ) c 
where c.rownum<=10`
		db:=AppInit.GetDB().Raw(sql).Scan(prods)
		if db.Error!=nil{
			log.Println(db.Error)
		}else{
			b,_:=json.Marshal(prods)
			cache:=&Models.BooksCache{CacheType:Models.CacheType_RankList,CacheContent:string(b),UpdateTime:time.Now()}
			AppInit.GetDB().Set("gorm:insert_option",
				"ON DUPLICATE KEY UPDATE update_time=now()").Create(cache)
			log.Println("排行榜缓存生成成功")
		}
	})
	if err!=nil{
		log.Println(err)
	}
}
