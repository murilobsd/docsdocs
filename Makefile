.PHONY: default client deps clean fmt

export GOPATH:=$(shell pwd)
export GOBIN=$(shell pwd)/bin

SHELL := /bin/bash

default: all

deps:
	@echo "--> intalling deps"
	@glide install
	$(shell ln -sr vendor/* src/)
	@echo ""

client:
	@echo "--> intall client"
	@go install docsdocs/main/docs
	@echo ""

fmt:
	@echo "--> formating"
	@go fmt docsdocs/...
	@echo ""

clean:
	@echo "--> formating"
	go clean -i -r docsdocs/...
	@echo ""

all: fmt client
