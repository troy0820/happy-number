SHELL:=/bin/bash


all: happy


happy: clean
	go build -o happy

clean:
	rm -rf happy



.PHONY: happy clean
