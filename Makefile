APPLICATION_NAME := $(shell grep "const ApplicationName " version.go | sed -E 's/.*"(.+)"$$/\1/')
BIN_NAME=${APPLICATION_NAME}

VERSION := $(shell grep "const Version " version.go | sed -E 's/.*"(.+)"$$/\1/')
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
default: help

help: ## Presents the available commands
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

up-deps: ## Install project's dependecies
	go mod tidy

run-test: ## Run project's tests
	mkdir -p ./test/cover
	go test -race -coverpkg= ./... -coverprofile=./test/cover/cover.out
	go tool cover -html=./test/cover/cover.out -o ./test/cover/cover.html
