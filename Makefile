.PHONY: build run tvet tidy fmt clean

build: vet
	go build

run: vet
	go run main.go

test: vet
	go test ./...

vet: tidy
	go vet ./...

tidy: fmt
	go mod tidy

fmt: swag
	go fmt ./...

swag:

clean:
	go clean

