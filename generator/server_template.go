package generator

var serverTemplate = `
package main
import (
	"fmt"
	"net"
	"github.com/watchman1989/rninet/server"
	"{{.Rpath}}/router"
	"{{.Rpath}}/proto/{{.Package.Name}}"
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
	
	{{.Package.Name}}.Register{{.Service.Name}}Server(server.GRPCServer(), routerServer)

	if err := server.serve(); err != nil {
		fmt.Printf("SERVER_ERROR: %v\n", err)
		return
	}
}
`