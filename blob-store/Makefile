export CGO_ENABLED=off

.PHONY: all build dep

all: build

dep:
	go mod tidy

generate:
	go generate ./api

lint:
	golangci-lint run

build:
	go build -o blob-store cmd/blob-store