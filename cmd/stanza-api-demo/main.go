package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// TODO (Laura): add a page that shows currently running requests / recent history
// TODO (Laura): graph burst and best effort statuses

var (
	hub       string
	verbose   bool
	port      int
	rest_port int
)

const (
	decoratorName = "Expensive Limited Resource"
	env           = "tiered_quota"
	apikey        = "2dacc6dd-e1ec-4b09-ac02-ff3bfa2213df"
)

func main() {
	flag.StringVar(&hub, "hub", "localhost:9020", "The host:port hub to issue queries against.")
	flag.BoolVar(&verbose, "verbose", false, "Print out details on every success/failure.")
	flag.IntVar(&port, "metrics_port", 9277, "Prom metrics server port")
	flag.IntVar(&rest_port, "rest_port", 9278, "REST API and status page")

	flag.Parse()

	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	runner := MakeRequestRunner()

	router := gin.Default()
	router.POST("/run", runner.postRequest)
	router.LoadHTMLFiles("./cmd/stanza-api-demo/assets/status.tmpl")
	router.GET("/status", runner.status)
	router.Run(fmt.Sprintf(":%d", rest_port))
}
