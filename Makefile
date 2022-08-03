.PHONY: build
build:
	@go build cmd/geninit/main.go

run:
	@./main -filename ./example.go

.PHONY: show
show:
	@./show -filename ./example.go
