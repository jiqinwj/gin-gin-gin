package Services

import (
	"context"
	"gin-gin-gin/App"
)

type BookMetaRequest struct {
	Type string `json:"type"`
	UserID int `json:"uid"`
	BookID int `json:"bookid"`
}

///prods?size=
type BookListRequest struct {
	Size int `form:"size"`
	Type string `form:"t"`
}
type BookResponse struct {
	Result interface{} `json:"result"`
	Metas interface{} `json:"metas"`
}

//  /prods/300
type BookDetailRequest struct {
	BookID int `uri:"id" binding:"required,gt=0,max=70000"`
}


//图书列表相关的业务最终函数
func BookListEndPoint(book *BookService)  App.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req:=request.(*BookListRequest)
		result,err:=book.LoadBookList(req)
		return &BookResponse{Result:result},err
	}
}
//图书详细
func BookDetailEndPoint(book *BookService)  App.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req:=request.(*BookDetailRequest)
		result,metas,err:=book.LoadBookDetail(req)
		return &BookResponse{Result:result,Metas:metas},err
	}
}
//收藏图书最终函数
func BookFavEndPoint(book *BookService)  App.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req:=request.(*BookMetaRequest)
		getError:=book.BookFav(req)
		if err!=nil{
			return &BookResponse{Result:"error"},getError
		}
		return &BookResponse{Result:"success"},nil
	}
}

