package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	quota_granted = promauto.NewCounter(prometheus.CounterOpts{
		Name: "stanza-demo_quota_granted_total",
		Help: "The total number of processed events",
	})

	quota_not_granted = promauto.NewCounter(prometheus.CounterOpts{
		Name: "stanza-demo_quota_not_granted_total",
		Help: "The total number of processed events",
	})
)
