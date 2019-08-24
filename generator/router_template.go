package generator


var routerTemplate = `
package router

import (
	"context"
	"github.com/watchman1989/rninet/middleware"
	"{{.Rpath}}/router"
	"{{.Rpath}}/proto/{{.Package.Name}}"
)

type RouterServer struct {}

{{range .Rpcs}}
func (s *RouterServer) {{.Name}} (ctx context.Context, req *{{$.Package.Name}}.{{.RequestType}}) (rsp *{{$.Package.Name}}.{{.ReturnsType}}, err error) {
	
	middlewareFunction := middleware.InsertMiddleware(middleware{{.Name}})
	middlewareRsp, err := middlewareFunction(context.Background(), req)
	rsp := middlewareRsp.(*{{$.Package.Name}}.{{.ReturnsType}})

	return rsp, err
}


func middleware{{.Name}} (ctx context.Context, req interface{}) (interface{}, error) {

	r := req.(*{{$.Package.Name}}.{{.RequestType}})
	handler := &handler.{{.Name}}Handler{}
	if err := handler.Check(ctx, r); err != nil {
		
		return nil, err
	}

	rsp, err := handler.Run(ctx, r)
	
	return rsp, err
}


{{end}}
`
