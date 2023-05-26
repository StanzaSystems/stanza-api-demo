package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	hub     string
	verbose bool
	port    int
)

const (
	decoratorName = "Expensive Limited Resource"
	env           = "tiered_quota"
	apikey        = "2dacc6dd-e1ec-4b09-ac02-ff3bfa2213df"
)

func main() {
	flag.StringVar(&hub, "hub", "localhost:9020", "The host:port hub to issue queries against.")
	flag.BoolVar(&verbose, "verbose", false, "Print out details on every success/failure.")
	flag.IntVar(&port, "port", 9277, "REST API and prom metrics server")

	flag.Parse()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
