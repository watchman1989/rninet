package ratelimit

import (
	"context"
	"github.com/watchman1989/rninet/middleware"
)

type Allower interface {
	Allow() bool
}


func RatelimitMiddleware (next middleware.MiddlewareFunction) middleware.MiddlewareFunction {

	return func(ctx context.Context, req interface{}) (rsp interface{}, err error) {

		

		rsp, err = next(ctx, req)

		return
	}
}
