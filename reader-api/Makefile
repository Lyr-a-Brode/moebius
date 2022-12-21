export CGO_ENABLED=off

.PHONY: all build dep

all: build

dep:
	go mod tidy

generate:
	go generate ./rest_api

lint:
	golangci-lint run

build:
	go build -o reader-api cmd/reader-api