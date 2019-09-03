package server

type Options struct {
	Name string
	Addr string
}

type Option func (opts *Options)


func Name (name string) Option {
	return func (opts *Options) {
		opts.Name = name
	}
}

func Addr (addr string) Option {
	return func (opts *Options) {
		opts.Addr = addr
	}
}
