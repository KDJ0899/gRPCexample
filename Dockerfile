FROM golang:1.13.8 as builder
WORKDIR $GOPATH/src/KDJ/gRPCexample
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o ex_server ex_sever/main.go

FROM alpine:latest
COPY --from=builder . .
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
CMD ["ex_server"]
