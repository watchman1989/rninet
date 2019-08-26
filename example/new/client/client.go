package main

import (
	"context"
	"fmt"
	"github.com/watchman1989/rninet/example/new/server/proto/greet"
	"google.golang.org/grpc"
	"os"
)

var (
	addr = "127.0.0.1:10001"
)



func main() {

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("GRPC_DIAL_ERROR: %v\n", err)
		return
	}
	defer conn.Close()

	client := greet.NewGreetingClient(conn)

	name := os.Args[1]

	rsp, err := client.Hello(context.Background(), &greet.GreetRequest{Name: name})
	if err != nil {
		fmt.Printf("CALL_HELLO_ERROR: %v\n", err)
		return
	}

	fmt.Printf("RESPONSE: %d %s\n", rsp.Status, rsp.Reply)


}
