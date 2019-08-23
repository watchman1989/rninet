package generator



var serverTemplate = `
package main
import (
	"fmt"
	"net"
	"google.golang.org/grpc"
	"{{.Rpath}}/router"
)

var (
	port = ":10001"
	routerServer = &router.RouterServer{}
)

func main() {

	fmt.Printf("START_SERVER\n")
	
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("NET_LISTEN_ERROR: %v\n", err)
		return
	}
	
	srv := grpc.NewServer()
	
	{{.Package.Name}}.Register{{.Service.Name}}Server(srv, routerServer)

	srv.Serve(lis)

}
`