.PHONY: build-api

export GOFLAGS=-mod=vendor
export CGO_ENABLED=0
export GOGC=off

build-api:
	go build -gcflags='-e' -o bin/api cmd/api/main.go
