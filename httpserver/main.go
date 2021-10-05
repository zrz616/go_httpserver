package main

import (
    "flag"
    "fmt"
    "net"
    "net/http"
    "os"
    "strings"

    "github.com/golang/glog"
)

// 获取客户端ip
func getClientIP(r *http.Request) string {
    if ip, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
        return ip
    }
    return ""
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
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
}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
    glog.V(2).Info("fooHandler")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, http.StatusOK)
    ip := getClientIP(r)
    glog.V(1).Infof("IP: %s; Check Ok", ip)
}

func main() {
    host := flag.String("host", "localhost", "specify the server binding host, default is localhost")
    port := flag.String("port", "8080", "specify the server binding port, default 8080")
    flag.Parse()
    defer glog.Flush()
    // log.SetFlags(glog.Ldate | glog.Ltime)
    address := *host + ":" + *port
    glog.V(1).Infof("listening %s\n", address)
    http.HandleFunc("/foo", fooHandler)
    http.HandleFunc("/healthz", healthcheckHandler)
    glog.Fatal(http.ListenAndServe(address, nil))
}
