
package handler

import (
	"context"
	"github.com/watchman1989/rninet/example/new/server/proto/greet"
)

type HelloHandler struct {}


func (s *HelloHandler) Check (ctx context.Context, req *greet.GreetRequest) (error) {
	
	//check code

	return nil
}


func (s *HelloHandler) Run (ctx context.Context, req *greet.GreetRequest) (rsp *greet.GreetResponse, err error) {
	
	//work code

	return rsp, err
}

