package main

import (
    "context"
    "flag"
    "fmt"
    "io"
    "net/http"
    "os"
    "os/signal"
    "strings"
    "syscall"
    "time"

    "github.com/golang/glog"
)

func main() {
    flag.Set("v", "2")
    glog.V(1).Info("Starting service1")

    mux := http.NewServeMux()
    mux.HandleFunc("/", rootHandler)

    srv := http.Server{
        Addr:    ":80",
        Handler: mux,
    }
    done := make(chan os.Signal, 1)
    signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

    go func() {
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            glog.Fatal("listen: %s\n", err)
        }
    }()
    glog.Info("Server Started")
    <-done
    glog.Info("Server Stopped")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := srv.Shutdown(ctx); err != nil {
        glog.Fatalf("Server Shutdown Failed:%+v", err)
    }
    glog.Info("Server Exited Properly")
}

func healthz(w http.ResponseWriter, r *http.Request) {
    glog.V(1).Info("healthz")
    for k, v := range r.Header {
        io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
    }

    io.WriteString(w, "ok\n")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
    glog.V(2).Info("service1 root")

    req, err := http.NewRequest("GET", "http://service2", nil)
    if err != nil {
        fmt.Printf("%s", err)
    }
    lowerCaseHeader := make(http.Header)
    for key, val := range r.Header {
        lowerCaseHeader[strings.ToLower(key)] = val
    }
    glog.Info("headers: ", lowerCaseHeader)
    req.Header = lowerCaseHeader
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        glog.Info("HTTP get failed with error: ", err)
    } else {
        glog.Info("HTTP get succeeded")
    }
    resp.Write(w)
}
