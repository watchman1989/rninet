package generator



var serverTemplate = `
package main
import (
	"fmt"
	"net"
	"google.golang.org/grpc"
)

var (
	port = ":10001"
)

func main() {

	fmt.Printf("")
	
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("NET_LISTEN_ERROR: %v\n", err)
		return
	}
	
	srv := grpc.NewServer()
	
	{{}}.Register{{.Service.Name}}Server(srv, new({{}}))

	srv.Serve(lis)

}


`