package middleware

import (
	"context"
	"fmt"
)

type MiddlewareFunction func (ctx context.Context, req interface{}) (rsp interface{}, err error)

type Middleware func (MiddlewareFunction) MiddlewareFunction

var userMiddleware []Middleware

func Chain (outer Middleware, others ...Middleware) Middleware {

	return func (next MiddlewareFunction) MiddlewareFunction {
		for i := len(others) - 1; i >= 0; i-- {
			next = others[i](next)
		}
		return outer(next)
	}
}

func Use (m ...Middleware) {
	userMiddleware = append(userMiddleware, m...)
}


func InsertMiddleware(middlewareFunction MiddlewareFunction) (MiddlewareFunction) {

	var middlewareList []Middleware

	if len(userMiddleware) > 0 {
		middlewareList = append(middlewareList, userMiddleware...)
	}

	if len(middlewareList) > 0 {
		m := Chain(middlewareList[0], middlewareList[1:]...)
		fmt.Println("+++++++++++")
		return m(middlewareFunction)
	}

	fmt.Println("-----------")
	return middlewareFunction
}

