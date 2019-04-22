default_target: all

all: lint build

.PHONY: all tidy build lint

LINTER_INSTALLED := $(shell sh -c 'which golangci-lint')

tidy:
	go mod tidy

build: tidy
	go build

lint:
ifdef LINTER_INSTALLED
	golangci-lint run
else
	$(error golangci-lint not found, skipping linting. Installation instructions: https://github.com/golangci/golangci-lint#ci-installation)
endif