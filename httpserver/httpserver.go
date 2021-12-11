package httpserver

import (
    "fmt"
    "math/rand"
    "net"
    "net/http"
    "os"
    "strings"
    "time"

    "github.com/golang/glog"
    "github.com/prometheus/client_golang/prometheus/promhttp"

    "github.com/zrz616/httpserver/metrics"
)

const welcomeMsg = "Check the Version in Responses Headers"

// 获取客户端ip
func getClientIP(r *http.Request) string {
    if ip, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
        return ip
    }
    return ""
}

func randInt(min int, max int) int {
    rand.Seed(time.Now().UTC().UnixNano())
    return min + rand.Intn(max-min)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
    glog.V(2).Info("fooHandler")
    glog.V(2).Info(r.Header)
    timer := metrics.NewTimer()
    defer timer.ObserveTotal()
    delay := randInt(10, 2000)
    time.Sleep(time.Millisecond * time.Duration(delay))
    for k, v := range r.Header {
        w.Header().Set(k, strings.Join(v, ", "))
    }
    version := os.Getenv("VERSION")
    w.Header().Add("VERSION", version)
    w.WriteHeader(http.StatusOK)
    ip := getClientIP(r)
    glog.V(1).Infof("IP: %s; Status code: %d", ip, http.StatusOK)
    glog.V(2).Info(w.Header())
    w.Write([]byte(welcomeMsg))
}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
    glog.V(2).Info("fooHandler")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, http.StatusOK)
    ip := getClientIP(r)
    glog.V(1).Infof("IP: %s; Check Ok", ip)
}

// NewServer 根据addr提供http.Server
func NewServer(addr string) *http.Server {
    metrics.Register()
    mux := http.NewServeMux()
    mux.HandleFunc("/", rootHandler)
    mux.HandleFunc("/healthz", healthcheckHandler)
    mux.Handle("/metrics", promhttp.Handler())
    return &http.Server{
        Addr:    addr,
        Handler: mux,
    }
}
