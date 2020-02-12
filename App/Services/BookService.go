package Services

import (
	"gin-gin-gin/AppInit"
	"gin-gin-gin/Models"
)

//框架无关性代码
type BookService struct {
}
func(this *BookService) LoadBookList(req  *BookListRequest) *Models.BookList {
	prods:=&Models.BookList{}
	AppInit.GetDB().Limit(req.Size).Order("book_id desc").Find(prods)
	return prods
}
func(this *BookService) LoadBookDetail(req  *BookDetailRequest) *Models.Books {
	prods:=&Models.Books{}
	if AppInit.GetDB().Find(prods,req.BookID).RowsAffected!=1{
		return nil
	}
	return prods
}



