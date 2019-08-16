package infra

import (
	"errors"
	"sync"
)

var (
	starterManage *StarterManage = &StarterManage{
		Starters: make(map[string]Starter),
	}
)


type Starter interface {
	Init() error
	Stop() error
}


type StarterManage struct {
	Starters map[string]Starter
	mu sync.Mutex
}

func (s *StarterManage) add (name string, starter Starter) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Starters[name] = starter
}

func (s *StarterManage) all () map[string]Starter {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.Starters
}

func AddStarter (name string, starter Starter) {
	starterManage.add(name, starter)
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
