# Inspired from article: https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
# and 

VERSION ?= $(shell cat VERSION)
.PHONY: help version
.DEFAULT_GOAL := help

help:
	@echo "go-rest-api: "
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: vendor ## Build docker image
	@echo "formatting sources"
	go fmt ./...

	@echo "building Docker Image..."
	docker build . -t zebeurton/go-rest-api

run: test vendor ## Runs locally docker container for development purpose
	docker-compose up -d

stop: ## Stops docker container in dev
	docker-compose down --rmi local

vendor: ## Install dependencies for server
	go mod vendor

test: ## Runs tests
	go fmt ./...
	go test -vet all ./... -count=1 -race

version: ## Echo current version
	@cat VERSION