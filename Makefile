#!/bin/bash

export PKGS=$(shell go list ./... | grep -v vendor/)

all:
	        build test run-server run-client
build:
	        build-server build-client

build-server:
	        @echo "Building server..."
		@cd proto/ && bash ./compile.sh & cd ..
	        @go build -o server ./cmd/server/
	        @echo "Done."

build-client:
	        @echo "Building client..."
		@cd proto/ && bash ./compile.sh & cd ..
	        @go build -o client ./cmd/client/
	        @echo "Done."

run-server:
	        @echo "Running server binary..."
	        @./server -c config.toml

run-client:
	        @echo "Running client binary..."
	        @./client -c config.toml
