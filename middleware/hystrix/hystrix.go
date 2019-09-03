package hystrix

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/watchman1989/rninet/middleware"
)

func HystrixMiddleware (next middleware.MiddlewareFunction) middleware.MiddlewareFunction {
	return func(ctx context.Context, req interface{}) (rsp interface{}, err error) {

		if err := hystrix.Do("", nil, nil); err != nil {
			return nil, err
		}

		rsp, err = next(ctx, req)
		return
	}
}
