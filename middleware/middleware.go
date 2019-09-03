package middleware

import (
	"context"
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


