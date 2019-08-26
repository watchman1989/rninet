
package handler

import (
	"context"
	"fmt"
	"github.com/watchman1989/rninet/example/new/server/proto/greet"
)

type WelcomeHandler struct {}


func (s *WelcomeHandler) Check (ctx context.Context, req *greet.GreetRequest) (error) {
	
	//check code

	return nil
}


func (s *WelcomeHandler) Run (ctx context.Context, req *greet.GreetRequest) (rsp *greet.GreetResponse, err error) {
	
	//work code

	fmt.Println("Hello: ", req.Name)
	rsp = &greet.GreetResponse{Status: 0, Reply: "Welcome: " + req.Name}

	return rsp, err
}

