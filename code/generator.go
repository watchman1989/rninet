package code

import (
	"errors"
	"fmt"
)

var (
	generatorManager *GeneratorManager = &GeneratorManager{
		generators: make(map[string]Generator),
	}
)


type Generator interface {
	Gen(opts ...Option) error
}


type GeneratorManager struct {
	generators map[string]Generator
}

func (g *GeneratorManager) Add (name string, generator Generator, opts ...Option) error {

	if _, ok := g.generators[name]; ok {
		fmt.Printf("GENERATOR %s IS EXISTS\n", name)
		return errors.New("GENERATOR_IS_EXISTS")
	}

	g.generators[name] = generator

	if err := generator.Gen(opts...); err != nil {
		return err
	}

	return  nil
}


func AddGenerator (name string, gen Generator, opts ...Option) error {

	err := generatorManager.Add(name, gen, opts...)

	return err
}


