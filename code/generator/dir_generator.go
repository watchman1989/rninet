package generator

import (
	"fmt"
	"os"
	"path"
	"projects/rninet/code"
)


type DirGenerator struct {
	options *code.Options
	dirs []string
}


func (d *DirGenerator) Gen (opts ...code.Option) error {

	d.options = &code.Options{}
	for _, opt := range opts {
		opt(d.options)
	}

	for _, dir := range d.dirs {
		fullPath := path.Join(d.options.Output, dir)
		err := os.MkdirAll(fullPath, 0755)
		if err != nil {
			fmt.Printf("MAKEDIR %s ERROR\n", fullPath)
			continue
		}
	}

	return nil
}