FROM golang:1.16-alpine3.10 AS build
WORKDIR /go/src/httpserver/
COPY . /go/src/httpserver/
RUN go build -o /bin/httpserver

FROM alpine3.10
COPY --from=build /bin/httpserver /bin/httpserver
ENTRYPOINT ["/bin/httpserver"]
CMD ["-v=1 -logtostderr=true"]