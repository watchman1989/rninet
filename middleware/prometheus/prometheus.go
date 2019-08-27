package prometheus

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/watchman1989/rninet/generator"
	"github.com/watchman1989/rninet/middleware"
)


var (
	DefaultServerMetrics = NewServerMetrics()
)


type ServerMetrics struct {
	RequestCount *prometheus.CounterVec

}


func NewServerMetrics () *ServerMetrics {
	return &ServerMetrics{
		RequestCount: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "rninet_server_request_total",
				Help: "Total count of rpc request",
			},
			[]string{"service", "method"},
		),
	}
}


func (s *ServerMetrics) Inc (ctx context.Context, srvName string, method string) {
	s.RequestCount.WithLabelValues(srvName, method).Inc()
}


func PrometheusMiddleware (next middleware.MiddlewareFunction) middleware.MiddlewareFunction {

	return func(ctx context.Context, req interface{}) (rsp interface{}, err error) {
		middlewareMeta := generator.GetMiddlewareMeta(ctx)
		DefaultServerMetrics.Inc(ctx, middlewareMeta.ServiceName, middlewareMeta.Method)

		rsp, err = next(ctx, req)

		return
	}
}