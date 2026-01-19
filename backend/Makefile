.PHONY: build run clean deps

build:
	go build -o server main.go

run:
	go run main.go

deps:
	go mod tidy

clean:
	rm -f server

install:
	go install ./...