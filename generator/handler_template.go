package generator

var handlerTemplate = `
package handler

import (
	"context"
)

type {{.Rpc.Name}} struct {}


func (s *{{.Rpc.Name}}) {{}}


`
