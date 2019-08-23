package generator


var routerTemplate = `
package router

import (
	"context"
	"{{.Rpath}}/router"
	"{{.Rpath}}/proto/{{.Package.Name}}"
)

type RouterServer struct {}

{{range .Rpcs}}
func (s *RouterServer) {{.Name}} (ctx context.Context, req *{{$.Package.Name}}.{{.RequestType}}) (rsp *{{$.Package.Name}}.{{.ReturnsType}}, err error) {
	

	return
}

{{end}}
`
