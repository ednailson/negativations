APPLICATION_NAME := $(shell grep "const ApplicationName " version.go | sed -E 's/.*"(.+)"$$/\1/')
BIN_NAME=${APPLICATION_NAME}

BASE_VERSION := $(shell grep "const Version " version.go | sed -E 's/.*"(.+)"$$/\1/')
VERSION="${BASE_VERSION}.$(shell date +%s | head -c 9)"

.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

default: help

up-deps: ## Install projects dependecies with GOMOD
	go mod tidy

docker-build: ## Build docker image
	docker build -t ${APPLICATION_NAME} ./

tests: ## Run project tests
	mkdir -p ./test/cover
	go test -race -coverpkg= ./... -coverprofile=./test/cover/cover.out
	go tool cover -html=./test/cover/cover.out -o ./test/cover/cover.html