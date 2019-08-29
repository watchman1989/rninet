package kafka

type Options struct {
	Addrs []string
}

type Option func (opts *Options)

func WithAddrs(addrs []string) Option {
	return func(opts *Options) {
		opts.Addrs = addrs
	}
}