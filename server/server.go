package server

import (
	"fmt"
	"github.com/watchman1989/rninet/middleware"
	"google.golang.org/grpc"
	"net"
)

var (
	options *Options
	rninetServer = &RninetServer{
		Server: grpc.NewServer(),
	}
)

type RninetServer struct {
	*grpc.Server
	middleware []middleware.Middleware
}

func Use (m ...middleware.Middleware) {
	rninetServer.middleware = append(rninetServer.middleware, m)
}

func Init(opts ...Option) error {

	for _, opt := range opts {
		opt(options)
	}

	if options.Name == "" {
		options.Name = "rninet.server"
	}

	if options.Addr == "" {
		options.Addr = ":9090"
	}

	return nil
}


func Serve() error {

	lis, err := net.Listen("tcp", options.Addr)
	if err != nil {
		return err
	}

	if err := rninetServer.Serve(lis); err != nil {
		return err
	}

	return nil
}


func GRPCServer() *grpc.Server {
	return rninetServer.Server
}


func InsertMiddleware(middlewareFunction middleware.MiddlewareFunction) (middleware.MiddlewareFunction) {

	var middlewareList []middleware.Middleware

	if len(rninetServer.middleware) > 0 {
		middlewareList = append(middlewareList, rninetServer.middleware...)
	}

	if len(middlewareList) > 0 {
		m := middleware.Chain(middlewareList[0], middlewareList[1:]...)
		fmt.Println("+++++++++++")
		return m(middlewareFunction)
	}

	fmt.Println("-----------")
	return middlewareFunction
}
