# httpserver
## Feature
1.当访问`localhost/foo`时，接收客户端 request，并将 request 中带的 header 写入 response header
2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4.当访问 localhost/healthz 时，应返回200

## Start
设置环境变量`VERSION`
日志级别：
  1. V1 -- 正常信息
  2. V2 -- bebug信息

``` shell
Usage of ./main:
  -alsologtostderr
        log to standard error as well as files
  -host string
        specify the server binding host, default is localhost (default "localhost")
  -log_backtrace_at value
        when logging hits line file:N, emit a stack trace
  -log_dir string
        If non-empty, write log files in this directory
  -logtostderr
        log to standard error instead of files
  -port string
        specify the server binding port, default 8080 (default "8080")
  -stderrthreshold value
        logs at or above this threshold go to stderr
  -v value
        log level for V logs
  -vmodule value
        comma-separated list of pattern=N settings for file-filtered logging
```