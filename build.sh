#!/usr/bin/env sh

export GO111MODULE=on
export GOPROXY=https://goproxy.io
export GOSUMDB=sum.golang.google.cn
GOOS=linux GOARCH=amd64 go build -o bin/linux_64/server cmd/server.go
GOOS=windows GOARCH=amd64 go build -o bin/windows_64/server.exe cmd/server.go
