package infra

import (
	"context"
	"errors"
	"sync"
)

var (
	starterManage *StarterManage = &StarterManage{
		Starters: make(map[string]Starter),
	}
)

type Starter interface {
	Init(ctx context.Context, opts ...interface{}) error
	Stop() error
}


type StarterManage struct {
	Starters map[string]Starter
	mu sync.Mutex
}

func (s *StarterManage) add (ctx context.Context, name string, starter Starter, opts ...interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Starters[name] = starter
	starter.Init(ctx, opts...)
}

func (s *StarterManage) all () map[string]Starter {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.Starters
}

func AddStarter (ctx context.Context, name string, starter Starter, opts ...interface{}) {
	starterManage.add(ctx, name, starter, opts...)
}

func AllStarters () map[string]Starter {
	return starterManage.all()
}

func GetStarter (name string) (Starter, error) {

	starter, ok := starterManage.Starters[name]
	if !ok {
		return nil, errors.New("STARTER_NOT_EXISTS")
	}

	return starter, nil
}
