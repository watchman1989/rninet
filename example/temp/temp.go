package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main()  {


	http.Handle("/matrics", promhttp.Handler())
	http.ListenAndServe(":10002", nil)
}
