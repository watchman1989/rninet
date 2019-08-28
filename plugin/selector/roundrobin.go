package selector

import (
	"context"
	"errors"
)

type RoundrobinSelect struct {
	index int
}

func (r *RoundrobinSelect) Name () string {
	return "roundrobin"
}


func (r *RoundrobinSelect) Select (ctx context.Context, nodes []*Node) (*Node, error) {

	if len(nodes) == 0 {
		return nil, errors.New("NO_NODES")
	}

	r.index = (r.index + 1) % (len(nodes))

	selectNode := nodes[r.index]

	return selectNode, nil
}
