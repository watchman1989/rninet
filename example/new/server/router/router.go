
package router

import (
	"context"
	"github.com/watchman1989/rninet/server"
	"github.com/watchman1989/rninet/example/new/server/proto/first"
	"github.com/watchman1989/rninet/example/new/server/handler"
)

type RouterServer struct {}


func (s *RouterServer) SayHello (ctx context.Context, req *first.GreetRequest) (rsp *first.GreetResponse, err error) {
	
	middlewareFunction := server.InsertMiddleware(middlewareSayHello)
	middlewareRsp, err := middlewareFunction(context.Background(), req)
	rsp = middlewareRsp.(*first.GreetResponse)

	return rsp, err
}


func middlewareSayHello (ctx context.Context, req interface{}) (interface{}, error) {

	r := req.(*first.GreetRequest)
	handler := &handler.SayHelloHandler{}
	if err := handler.Check(ctx, r); err != nil {
		
		return nil, err
	}

	rsp, err := handler.Run(ctx, r)
	
	return rsp, err
}



