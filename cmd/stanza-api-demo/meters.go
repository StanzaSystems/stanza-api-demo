package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	quota_granted_name     = "stanza_demo_quota_granted_total"
	quota_granted_help     = "The total number of quota tokens granted"
	quota_not_granted_name = "stanza_demo_quota_not_granted_total"
	quota_not_granted_help = "The total number of quota tokens not granted"
	quota_error_name       = "stanza_demo_quota_error_total"
	quota_error_help       = "The total number of errors observed when requesting quota"
)

type meters struct {
	quota_granted     *prometheus.CounterVec
	quota_not_granted *prometheus.CounterVec
	quota_error       *prometheus.CounterVec
	latency           prometheus.Histogram
}

func MakeMeters() *meters {
	res := meters{}

	labels := []string{"priorityBoost", "environment", "tags"}

	res.quota_granted = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: quota_granted_name,
			Help: quota_granted_help,
		},
		labels)

	res.quota_not_granted = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: quota_not_granted_name,
			Help: quota_not_granted_help,
		},
		labels)

	res.quota_error = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: quota_error_name,
			Help: quota_error_help,
		},
		labels)

	res.latency = promauto.NewHistogram(
		prometheus.HistogramOpts{
			Name: "stanza_demo_quota_latency",
			Help: "Latency distribution for requests to Stanza Hub",
		},
	)
	return &res
}

func (m *meters) GetQuotaGrantedCounter() *prometheus.CounterVec {
	return m.quota_granted
}

func (m *meters) GetQuotaNotGrantedCounter() *prometheus.CounterVec {
	return m.quota_not_granted
}

func (m *meters) GetQuotaErrorCounter() *prometheus.CounterVec {
	return m.quota_error
}
