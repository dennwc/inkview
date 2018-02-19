#!/usr/bin/env bash
export GOROOT=/go
export GOPATH=/gopath
export PATH="$GOROOT/bin:$PATH"
cd /app
CC=arm-none-linux-gnueabi-gcc GOOS=linux GOARCH=arm GOARM=5 CGO_ENABLED=1 go build "$@"
