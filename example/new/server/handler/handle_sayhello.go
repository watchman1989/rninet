
package handler

import (
	"context"
	"github.com/watchman1989/rninet/example/new/server/proto/first"
)

type SayHelloHandler struct {}


func (s *SayHelloHandler) Check (ctx context.Context, req *first.GreetRequest) (error) {
	
	//check code

	return nil
}


func (s *SayHelloHandler) Run (ctx context.Context, req *first.GreetRequest) (rsp *first.GreetResponse, err error) {
	
	//work code

	return rsp, err
}

