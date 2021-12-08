package main

import (
    "context"
    "flag"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/golang/glog"
    "github.com/zrz616/httpserver/httpserver"
)

func main() {
    addr := flag.String("addr", ":8080", "specify the server binding address")
    glogLevel := os.Getenv("GLOG_LEVEL")
    flag.Set("v", glogLevel)
    flag.Set("logtostderr", "true")
    flag.Parse()
    defer glog.Flush()
    // log.SetFlags(glog.Ldate | glog.Ltime)
    glog.V(1).Infof("listening %s\n", *addr)
    // glog.Fatal(http.ListenAndServe(*addr, nil))
    srv := httpserver.NewServer(*addr)
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            glog.Fatalf("Listen: %s\n", err)
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
