package metrics

import (
    "fmt"
    "time"

    "github.com/prometheus/client_golang/prometheus"
)

const (
    metricsNamespace = "httpserver"
)

var (
    functionLatency = CreateExecutionTimeMrtric(metricsNamespace, "Time spent.")
)

type ExecutionTimer struct {
    histo *prometheus.HistogramVec
    start time.Time
    last  time.Time
}

func NewTimer() *ExecutionTimer {
    return NewExecutionTimer(functionLatency)
}

func NewExecutionTimer(histo *prometheus.HistogramVec) *ExecutionTimer {
    now := time.Now()
    return &ExecutionTimer{
        histo: histo,
        start: now,
        last:  now,
    }
}

func (t *ExecutionTimer) ObserveTotal() {
    (*t.histo).WithLabelValues("total").Observe(time.Now().Sub(t.start).Seconds())
}

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
