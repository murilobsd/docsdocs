.PHONY: default client deps clean fmt lint glide

export GOPATH:=$(shell pwd)
export GOBIN=$(shell pwd)/bin

SHELL := /bin/bash

default: all

deps:
	@echo "--> intalling deps"
	@glide install
	@cp -R vendor/* src/
	@echo ""

client:
	@echo "--> intall client"
	@go install docsdocs/main/docs
	@echo ""

fmt:
	@echo "--> formating"
	@go fmt docsdocs/...
	@echo ""

lint:
	@echo "--> lint "
	@golint src/docsdocs/...
	@echo ""

clean:
	@echo "--> formating"
	go clean -i -r docsdocs/...
	@echo ""

install_glide:
	@command -v glide >/dev/null ; if [ $$? -ne 0 ]; then \
		echo "--> installing glide"; \
		curl https://glide.sh/get | sh; \
	fi

test:
	@go test -v $(shell go list docsdocs/... | grep -v docsdocs/client/views/gui)

coverage:
	@goverage -v -coverprofile=coverage.txt $(shell go list docsdocs/... | grep -v docsdocs/client/views/gui)

all: fmt client
