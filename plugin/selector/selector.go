package selector

import "context"

const (
	DEFAULT_WIEGHT = 1
)


type Node struct {
	Addr string
	Weight int
}

type Selector interface {
	Name() string
	Select(ctx context.Context, nodes []*Node) (*Node, error)
}
