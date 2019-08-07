package registry


type Options struct {
	Addrs []string
	Path string
	TTL int64
	Interval int64
}

type Option func (opts *Options)


func WithAddrs (addrs []string) Option {
	return func (opts *Options) {
		opts.Addrs = addrs
	}
}

func WithPath (path string) Option {
	return func (opts *Options) {
		opts.Path = path
	}
}

func WithTTL (ttl int64) Option {
	return  func (opts *Options) {
		opts.TTL = ttl
	}
}


func WithInterval (interval int64) Option {
	return func (opts *Options) {
		opts.Interval = interval
	}
}
