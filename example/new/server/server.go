
package main
import (
	"fmt"
	"github.com/watchman1989/rninet/server"
	"github.com/watchman1989/rninet/example/new/server/router"
	"github.com/watchman1989/rninet/example/new/server/proto/first"
)

var (
	routerServer = &router.RouterServer{}
)

func main() {

	fmt.Printf("START_SERVER\n")
	
	if err := server.Init(); err != nil {
		fmt.Printf("SERVER_INIT_ERROR: %v\n", err)
		return
	}
	
	first.RegisterGreetServer(server.GRPCServer(), routerServer)

	if err := server.Serve(); err != nil {
		fmt.Printf("SERVER_ERROR: %v\n", err)
		return
	}
}
