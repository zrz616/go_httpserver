package main

import (
    "context"
    "flag"
    "fmt"
    "net"
    "net/http"
    "os"
    "os/signal"
    "strings"
    "syscall"
    "time"

    "github.com/golang/glog"
)

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
}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
    glog.V(2).Info("fooHandler")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, http.StatusOK)
    ip := getClientIP(r)
    glog.V(1).Infof("IP: %s; Check Ok", ip)
}

func main() {
    addr := flag.String("addr", ":8080", "specify the server binding address")
    glogLevel := os.Getenv("GLOG_LEVEL")
    flag.Set("v", glogLevel)
    flag.Set("logtostderr", "true")
    flag.Parse()
    defer glog.Flush()
    // log.SetFlags(glog.Ldate | glog.Ltime)
    glog.V(1).Infof("listening %s\n", *addr)
    mux := http.NewServeMux()
    mux.HandleFunc("/", rootHandler)
    mux.HandleFunc("/healthz", healthcheckHandler)
    // glog.Fatal(http.ListenAndServe(*addr, nil))
    srv := &http.Server{
        Addr:    *addr,
        Handler: mux,
    }

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            glog.Fatal("Listen: %s\n", err)
        }
    }()
    glog.V(1).Info("httpserver started")
    <-quit
    glog.V(1).Info("httpserver shutting down...")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := srv.Shutdown(ctx); err != nil {
        glog.Fatal("httpserver forced to shutdown: ", err)
    }

    glog.V(1).Info("httpserver exiting")
}
