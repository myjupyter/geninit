VERSION=$(shell git describe --tag)

.PHONY: build
build:
	@go build -ldflags "-X main.version=$(VERSION)" -o bin/geninit main.go
