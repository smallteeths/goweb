FROM golang:latest

MAINTAINER Razil "503630985@qq.com"

WORKDIR $GOPATH/src

ADD . $GOPATH/src

RUN go build -o goweb

EXPOSE 9091

ENTRYPOINT  ["./goweb"]
