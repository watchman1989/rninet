
package router

import (
	"context"
	"github.com/watchman1989/rninet/middleware"
	"github.com/watchman1989/rninet/example/new/server/proto/greet"
	"github.com/watchman1989/rninet/example/new/server/handler"
)

type RouterServer struct {}


func (s *RouterServer) Hello (ctx context.Context, req *greet.GreetRequest) (rsp *greet.GreetResponse, err error) {
	
	middlewareFunction := middleware.InsertMiddleware(middlewareHello)
	middlewareRsp, err := middlewareFunction(context.Background(), req)
	rsp = middlewareRsp.(*greet.GreetResponse)

	return rsp, err
}


func middlewareHello (ctx context.Context, req interface{}) (interface{}, error) {

	r := req.(*greet.GreetRequest)
	handler := &handler.HelloHandler{}
	if err := handler.Check(ctx, r); err != nil {
		
		return nil, err
	}

	rsp, err := handler.Run(ctx, r)
	
	return rsp, err
}



func (s *RouterServer) Welcome (ctx context.Context, req *greet.GreetRequest) (rsp *greet.GreetResponse, err error) {
	
	middlewareFunction := middleware.InsertMiddleware(middlewareWelcome)
	middlewareRsp, err := middlewareFunction(context.Background(), req)
	rsp = middlewareRsp.(*greet.GreetResponse)

	return rsp, err
}


func middlewareWelcome (ctx context.Context, req interface{}) (interface{}, error) {

	r := req.(*greet.GreetRequest)
	handler := &handler.WelcomeHandler{}
	if err := handler.Check(ctx, r); err != nil {
		
		return nil, err
	}

	rsp, err := handler.Run(ctx, r)
	
	return rsp, err
}



