package consul

type Options struct {
	Addrs []string
	Timeout int64
	TTL int64
	Interval int64
}

type Option func (opts *Options)


func WithAddrs (addrs []string) Option {
	return func (opts *Options) {
		opts.Addrs = addrs
	}
}

func WithTimeout (timeout int64) Option {
	return func (opts *Options) {
		opts.Timeout = timeout
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

