FROM golang:latest

MAINTAINER Razil "503630985@qq.com"

WORKDIR $GOPATH/src/tool-backend

ADD . $GOPATH/src/tool-backend

RUN go get -d -v ./...

RUN go build -o goweb

EXPOSE 9091

ENTRYPOINT  ["./goweb"]
