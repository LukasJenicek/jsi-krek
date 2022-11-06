.PHONY: build-api

export GOFLAGS=-mod=vendor
export CGO_ENABLED=0
export GOGC=off

run:
	GOGC=off go build -o ./bin/api ./cmd/api/main.go

build-api:
	go build -gcflags='-e' -o bin/api cmd/api/main.go

vendor:
	go mod vendor && go mod tidy
