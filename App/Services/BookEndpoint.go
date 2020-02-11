package Services

import (
	"context"
	"gin-gin-gin/App"

	"gin-gin-gin/Models"
)

type BookListRequest struct {
	Size int `form:"size"`
}
type BookListResponse struct {
	Result *Models.BookList
}

func BookListEndPoint(book *BookService)  App.Endpoint {
   return func(ctx context.Context, request interface{}) (response interface{}, err error) {
	   req:=request.(*BookListRequest)
	   return &BookListResponse{Result:book.LoadBookList(req)},nil
   }
}