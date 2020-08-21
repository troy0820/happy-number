#!/bin/bash

go mod download

go build -o happy ./cmd/happy
