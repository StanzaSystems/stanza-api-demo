package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	hub          string
	hub_insecure bool
	verbose      bool
	port         int
	rest_port    int
)

const (
	guardName      = "Expensive Limited Resource"
	env            = "tiered_quota"
	apikey_default = "2dacc6dd-e1ec-4b09-ac02-ff3bfa2213df"
)

func main() {
	flag.StringVar(&hub, "hub", "hub.demo.getstanza.io:9020", "The hub address host:port to issue queries against.")
	flag.BoolVar(&hub_insecure, "hub_insecure", false, "Skip Hub TLS validation (for local development only).")
	flag.BoolVar(&verbose, "verbose", false, "Print out details on every success/failure.")
	flag.IntVar(&port, "metrics_port", 9277, "Prom metrics server port")
	flag.IntVar(&rest_port, "rest_port", 9278, "REST API")

	flag.Parse()

	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	runner := MakeRequestRunner()

	router := gin.Default()
	router.POST("/run", runner.postRequest)
	router.Run(fmt.Sprintf(":%d", rest_port))
}
