package Services

import (
	"encoding/json"
	"fmt"
	"gin-gin-gin/AppInit"
	"gin-gin-gin/Models"
	"strconv"
)

//框架无关性代码
type BookService struct {}
//加载图书列表
func(this *BookService) LoadBookList(req  *BookListRequest)  (*Models.BookList,error) {
	prods:=&Models.BookList{}
	if req.Type=="top"{
		return  this.LoadBookTopList(req)
	}

    db:=AppInit.GetDB().Limit(req.Size).Order("book_id desc").Find(prods)
 	if db.Error!=nil{
		return nil,db.Error
	}
	return prods,nil
}
//加载商品详细
func(this *BookService) LoadBookDetail(req  *BookDetailRequest) (*Models.Books,[]*Models.BookMetas,error) {
	prods:=&Models.Books{}
	if AppInit.GetDB().Find(prods,req.BookID).RowsAffected!=1{
		return nil,nil,fmt.Errorf("no book")
	}
    Models.NewBookMeta("click","1",prods.BookID).Save()
	metas:=[]*Models.BookMetas{}
	AppInit.GetDB().Where("item_id=?",prods.BookID).Find(&metas)
	return prods,metas,nil
}
//获取排行榜
func(this *BookService) LoadBookTopList(req  *BookListRequest)  (*Models.BookList,error) {
	prods:=&Models.BookList{}
	cache:=&Models.BooksCache{}
	db:=AppInit.GetDB().Last(cache,"cache_type="+strconv.Itoa(Models.CacheType_RankList))
	if db.Error==nil && db.RowsAffected==1{
		if err:=json.Unmarshal([]byte(cache.CacheContent),prods);err==nil{
			return prods,nil
		}else{
			return nil,err
		}
	}
//缓存没有生成
	return prods,nil
}

//收藏商品
func(this *BookService) BookFav(req *BookMetaRequest) error  {
	tx:=AppInit.GetDB().Begin()
	err2:=Models.NewBookFav(req.BookID,req.UserID).Save(tx) //商品收藏表
	err1:= Models.NewBookMeta("fav","1",req.BookID).Save(tx) //元数据表
	if err1!=nil || err2!=nil{
		tx.Rollback()
		return fmt.Errorf("error fav")
	}else{
		tx.Commit()
		return nil
	}


}

