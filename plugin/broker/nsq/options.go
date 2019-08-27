package nsq


type Options struct {
	Addr string
	Topic string
	Channel string
	Interval int64
}

type Option func (opts *Options)


func WithAddrs (addrs string) Option {
	return func (opts *Options) {
		opts.Addr = addrs
	}
}

func WithTopic (topic string) Option {
	return func (opts *Options) {
		opts.Topic = topic
	}
}

func WithChannel (channel string) Option {
	return func (opts *Options) {
		opts.Channel = channel
	}
}

func WithInterval (interval int64) Option {
	return func (opts *Options) {
		opts.Interval = interval
	}
}
