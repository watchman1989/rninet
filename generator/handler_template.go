package generator

var handlerTemplate = `
package handler

import (
	"context"
	"{{.Rpath}}/proto/{{.Package.Name}}"
)

type {{.Rpc.Name}}Handler struct {}


func (s *{{.Rpc.Name}}Handler) Check (ctx context.Context, req *{{.Package.Name}}.{{.Rpc.RequestType}}) (error) {
	
	//check code

	return nil
}


func (s *{{.Rpc.Name}}Handler) Run (ctx context.Context, req *{{.Package.Name}}.{{.Rpc.RequestType}}) (rsp *{{.Package.Name}}.{{.Rpc.ReturnsType}}, err error) {
	
	//work code

	return rsp, err
}

`
