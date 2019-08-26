
package handler

import (
	"context"
	"github.com/watchman1989/rninet/example/new/server/proto/greet"
)

type WelcomeHandler struct {}


func (s *WelcomeHandler) Check (ctx context.Context, req *greet.GreetRequest) (error) {
	
	//check code

	return nil
}


func (s *WelcomeHandler) Run (ctx context.Context, req *greet.GreetRequest) (rsp *greet.GreetResponse, err error) {
	
	//work code

	return rsp, err
}

