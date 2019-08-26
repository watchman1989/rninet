package generator

var clientTemplate = `
package main

import (
	"context"
	"fmt"
	"{{.Rpath}}/proto/{{.Package.Name}}"
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

	//Call rpc function

}

`
