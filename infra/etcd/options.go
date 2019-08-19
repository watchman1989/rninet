package etcd


type Options struct {
	Addrs []string
	Timeout int64
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

