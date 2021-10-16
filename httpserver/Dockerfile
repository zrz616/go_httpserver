FROM golang:1.16-alpine AS build
WORKDIR /go/src/httpserver/
COPY . /go/src/httpserver/
RUN go build -o /bin/httpserver

FROM busybox
COPY --from=build /bin/httpserver /bin/httpserver
ENTRYPOINT ["/bin/httpserver"]
CMD ["--help"]