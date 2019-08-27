package elasticsearch

type Options struct {
	Urls []string
	Sniff bool
}

type Option func (opts *Options)


func WithUrls (urls []string) Option {
	return func (opts *Options) {
		opts.Urls = urls
	}
}

func WithSniff (sniff bool) Option {
	return func (opts *Options) {
		opts.Sniff = sniff
	}
}
