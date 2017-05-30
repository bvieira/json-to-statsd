#!/bin/bash
set -e

COMPILE_OS=$1
docker run --rm -i -v "$(pwd)":/gopath/src/github.com/bvieira/json-to-statsd -e "GOPATH=/gopath" -w /gopath/src/github.com/bvieira/json-to-statsd golang:latest sh -c "go test ./... && CGO_ENABLED=0 GOOS=$COMPILE_OS go build -v -a -installsuffix cgo --ldflags=\"-s\" -o json-to-statsd ."
