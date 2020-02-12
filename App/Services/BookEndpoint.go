package Services

import (
	"context"
	"gin-gin-gin/App"
)
///prods?size=
type BookListRequest struct {
	Size int `form:"size"`
}
type BookResponse struct {
	Result interface{} `json:"result"`
}

//  /prods/300
type BookDetailRequest struct {
	BookID int `uri:"id" binding:"required,gt=0,max=70000"`
}


//图书列表相关的业务最终函数
func BookListEndPoint(book *BookService)  App.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req:=request.(*BookListRequest)
		return &BookResponse{Result:book.LoadBookList(req)},nil
	}
}
//图书详细
func BookDetailEndPoint(book *BookService)  App.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req:=request.(*BookDetailRequest)
		return &BookResponse{Result:book.LoadBookDetail(req)},nil
	}
}
