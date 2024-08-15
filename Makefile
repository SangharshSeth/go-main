.PHONY: build run migrate

build:
	go build -o bin/main cmd/main/main.go

run: build
	./bin/main
