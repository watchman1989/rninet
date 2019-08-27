package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)


var (
	cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_temperature_celsius",
		Help: "Current temperature of the CPU",
	})
	hdFailures = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "hd_errors_total",
		Help: "Number of hard-disk errors .",
	},
	[]string{"device", "service"},
	)
)


func init() {
	prometheus.MustRegister(cpuTemp)
	prometheus.MustRegister(hdFailures)
}


func main()  {


	http.Handle("/matrics", promhttp.Handler())
	http.ListenAndServe(":10002", nil)
}
