package main

import (
    "context"
    "gin-gin-gin/App"
    "gin-gin-gin/App/Services"
)

//@gen()
func GetMe(a *Services.BookDetailRequest,s string) (App.Endpoint,error) {
    return func(ctx context.Context, request interface{}) (response interface{}, err error) {
        return nil,nil
    },nil
}

func GetYou() string   {
    return "zhangsan"
}

func main()  {


}

