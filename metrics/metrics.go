package metrics

import (
    "fmt"

    "github.com/prometheus/client_golang/prometheus"
)

const (
    metricsNamespace = "httpserver"
)

var (
    functionLatency = CreateExecutionTimeMrtric(metricsNamespace, "Time spent.")
)

// Register 提供prometheus.Register
func Register() {
    err := prometheus.Register(functionLatency)
    if err != nil {
        fmt.Println(err)
    }
}

// CreateExecutionTimeMrtric 提供NewHistogramVec
func CreateExecutionTimeMrtric(namespace string, help string) *prometheus.HistogramVec {
    return prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Namespace: namespace,
            Name:      "execution_latency_seconds",
            Help:      help,
            Buckets:   prometheus.ExponentialBuckets(0.001, 2, 15),
        }, []string{"step"},
    )
}
