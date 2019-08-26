package handler

import (
	"context"
	"fmt"
	"github.com/watchman1989/rninet/middleware"
	"time"
)

func init () {
	fmt.Println("init cost middleware")
	middleware.Use(CostMiddleware)
}

func CostMiddleware (next middleware.MiddlewareFunction) middleware.MiddlewareFunction {

	return func (ctx context.Context, req interface{}) (rsp interface{}, err error) {

		beginTime := time.Now().UnixNano()
		rsp, err = next(ctx, req)
		endTime := time.Now().UnixNano()

		fmt.Printf("COST: %dus", (endTime - beginTime) / (1000))

		return
	}
}
