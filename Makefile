# Makefile
.PHONY: build run test

build:
	go build -o bin/punyGo main.go
run: build
	./bin/punyGo
