SHELL:=/bin/bash


all: happy

happy: clean test
	go build -o happy ./cmd/happy
test:
	go test ./... -v
clean:
	rm -rf happy



.PHONY: happy test clean
