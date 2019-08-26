
package main
import (
	"fmt"
	"net"
	"google.golang.org/grpc"
	"github.com/watchman1989/rninet/example/new/server/router"
	"github.com/watchman1989/rninet/example/new/server/proto/greet"
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
	
	greet.RegisterGreetingServer(srv, routerServer)

	srv.Serve(lis)

}
