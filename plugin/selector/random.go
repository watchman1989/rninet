package selector

import (
	"context"
	"errors"
	"math/rand"
)

type RandomSelect struct {}

func (r *RandomSelect) Name () string {
	return "random"
}

func (r *RandomSelect) Select (ctx context.Context, nodes []*Node) (*Node, error) {

	if len(nodes) == 0 {
		return nil, errors.New("NO_NODES")
	}

	var indexMap map[int]*Node = make(map[int]*Node)
	var index = 0
	for _, node := range nodes {
		if Weight == 0 {
			Weight = 1
		}
		for i := 0; i < Weight; i++ {
			indexMap[index] = node
			index++
		}
	}

	selectedNode := indexMap[rand.Intn(index)]

	return selectedNode, nil
}
