package generator

type Options struct {
	Output string
	ProtoFile string
	CliFlag bool
	SrvFlag bool
	PrometheusPort int
}


type Option func (opts *Options)


func WithOutput (output string) Option {
	return func(opts *Options) {
		opts.Output = output
	}
}

func WithCliFlag () Option {
	return func(opts *Options) {
		opts.CliFlag = true
	}
}

func WithSrvFlag () Option {
	return func(opts *Options) {
		opts.SrvFlag = true
	}
}

func WithProtoFile (protoFile string) Option {
	return func(opts *Options) {
		opts.ProtoFile = protoFile
	}
}

func WithPrometheusPort (prometheusPort int) Option {
	return func (opts *Options) {
		opts.PrometheusPort = prometheusPort
	}
}
