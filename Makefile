default: help

help: ## Presents the available commands
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

up-deps: ## Install project's dependecies
	go mod tidy
	go mod download

run-test: ## Run project's tests
	go test -race ./...

run-test-coverage: ## Run project's tests and prints coverage
	mkdir -p ./test/cover
	go test -race -coverpkg= ./... -coverprofile=./test/cover/cover.out
	go tool cover -html=./test/cover/cover.out -o ./test/cover/cover.html