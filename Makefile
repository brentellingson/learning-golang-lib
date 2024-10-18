.PHONY: build run tvet tidy fmt clean

build: lint
	go build

test: lint
	go test ./...

lint: fmt
	go vet ./...
	staticcheck ./...
	revive -formatter friendly ./...

fmt: swag
	go mod tidy
	go fmt ./...
	goimports -l -w .

swag:

clean:
	go clean

