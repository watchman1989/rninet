package redis


type Options struct {
	Host      string
	Port      int
	User      string
	Pass      string
}

type Option func (opts *Options)


func WithHost (host string) Option {
	return func (opts *Options) {
		opts.Host = host
	}
}

func WithPort (port int) Option {
	return func (opts *Options) {
		opts.Port = port
	}
}

func WithUser (user string) Option {
	return func (opts *Options) {
		opts.User = user
	}
}


func WithPass (pass string) Option {
	return func (opts *Options) {
		opts.Pass = pass
	}
}