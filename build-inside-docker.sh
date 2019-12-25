#!/usr/bin/env sh

docker run --rm -i \
  -w /data/src \
  -v $(pwd):/data/src \
  -v /data/lib/go:/data/lib/go \
  -e GO111MODULE=on \
  -e GOPROXY=https://goproxy.io \
  -e GOSUMDB=sum.golang.google.cn \
  -e GOPATH=/data/lib/go \
  golang:1.13 sh build.sh
