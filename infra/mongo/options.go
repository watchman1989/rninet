package mongo


type Options struct {
	Url string
}

type Option func (opts *Options)


func WithUrl (url string) Option {
	return func (opts *Options) {
		opts.Url = url
	}
}
