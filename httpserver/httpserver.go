package httpserver

import (
    "fmt"
    "net"
    "net/http"
    "os"
    "strings"

    "github.com/golang/glog"
)

const welcomeMsg = "Check the Version in Responses Headers"

// 获取客户端ip
func getClientIP(r *http.Request) string {
    if ip, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
        return ip
    }
    return ""
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
    glog.V(2).Info("fooHandler")
    glog.V(2).Info(r.Header)
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
    mux := http.NewServeMux()
    mux.HandleFunc("/", rootHandler)
    mux.HandleFunc("/healthz", healthcheckHandler)
    return &http.Server{
        Addr:    addr,
        Handler: mux,
    }
}
