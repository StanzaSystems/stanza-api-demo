package main

import (
	"fmt"
	"sort"

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
	quota_granted     map[string]*prometheus.CounterVec
	quota_not_granted map[string]*prometheus.CounterVec
	quota_error       map[string]*prometheus.CounterVec
}

func MakeMeters() *meters {
	res := meters{}
	res.quota_granted = make(map[string]*prometheus.CounterVec)
	res.quota_not_granted = make(map[string]*prometheus.CounterVec)
	res.quota_error = make(map[string]*prometheus.CounterVec)

	return &res
}

func (m *meters) GetQuotaGrantedCounter(labels map[string]string) *prometheus.CounterVec {
	labelStr := labelsToString(labels)

	fmt.Printf("getting quota granted counter, labelstr is %s\n", labelStr)

	res, ok := m.quota_granted[labelStr]
	if !ok {
		keys := getKeys(labels)

		res = promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: quota_granted_name,
				Help: quota_granted_help,
			},
			keys)

		m.quota_granted[labelStr] = res
	}
	return res
}

func (m *meters) GetQuotaNotGrantedCounter(labels map[string]string) *prometheus.CounterVec {
	labelStr := labelsToString(labels)

	res, ok := m.quota_not_granted[labelStr]
	if !ok {
		keys := getKeys(labels)

		res = promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: quota_not_granted_name,
				Help: quota_not_granted_help,
			},
			keys)

		m.quota_not_granted[labelStr] = res
	}
	return res
}

func (m *meters) GetQuotaErrorCounter(labels map[string]string) *prometheus.CounterVec {
	labelStr := labelsToString(labels)

	res, ok := m.quota_error[labelStr]
	if !ok {
		keys := getKeys(labels)

		res = promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: quota_error_name,
				Help: quota_error_help,
			},
			keys)

		m.quota_error[labelStr] = res
	}
	return res
}

// turn a set of tags into a string to use as mapkey for results
func labelsToString(labels map[string]string) string {
	result := ""
	keys := make([]string, 0)
	for k := range labels {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		result = result + k + ","
	}

	if len(result) > 0 {
		return result[0 : len(result)-1]
	} else {
		return result
	}
}

func getKeys(labels map[string]string) []string {
	keys := make([]string, 0)
	for k := range labels {
		keys = append(keys, k)
	}
	return keys
}
