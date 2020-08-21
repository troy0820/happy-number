#!/bin/sh
# Test the Golang application

go test ./... -v

go build -o happy ./cmd/happy
