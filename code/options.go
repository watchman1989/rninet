package code


type Options struct {
	Output string
	Client bool
	Server bool
}


type Option func (opts *Options)


func WithOutput (output string) Option {
	return func(opts *Options) {
		opts.Output = output
	}
}

func WithClient (client bool) Option {
	return func(opts *Options) {
		opts.Client = client
	}
}

func WithServer (server bool) Option {
	return func(opts *Options) {
		opts.Server = server
	}
}