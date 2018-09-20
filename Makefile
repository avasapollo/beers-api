APP_NAME = beers-api
GOENV = CGO_ENABLED=0 GOOS=linux GOARCH=amd64
PACKAGES := $(shell go list ./... | grep -v /vendor/)

default: help

.PHONY: deps
deps: ## Install all the deps
	dep ensure -v

.PHONY: build
build: ## Build the binary
	$(GOENV) go build -a -installsuffix cgo \
	-o $(APP_NAME) ./cmd/$(APP_NAME)

.PHONY: test
test: ## Run Tests into the packages
	go test -v $(PACKAGES)

.PHONY: tests
tests: deps test

.PHONY: all
all: deps test build

.PHONY: help
help: ## Help
	@echo "Please use 'make <target>' where <target> is ..."
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
