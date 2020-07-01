# Inspired from article: https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
# and 

VERSION ?= $(shell cat VERSION)
.PHONY: help version
.DEFAULT_GOAL := help

help:
	@echo "go-rest-api: "
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: vendor ## Build docker image for dev purpose
	@echo "formatting sources"
	go fmt ./...

	@echo "building latest docker Image..."
	docker build . -t zebeurton/go-rest-api:latest

release: ## Build docker image to release a new version
	@echo "formatting sources"
	go fmt ./...

	@echo "building latest docker Image..."
	docker build . -t zebeurton/go-rest-api:$(VERSION)
	docker push zebeurton/go-rest-api:$(VERSION)
	
run: test vendor ## Runs locally docker container for development purpose
	docker-compose up -d

stop: ## Stops docker container in dev
	docker-compose down --rmi local

vendor: ## Install dependencies for server
	go mod vendor

test: ## Runs tests
	go fmt ./...
	go test -vet all ./... -count=1 -race

bump_patch: ## Release a patch version
	./scripts/bump.sh --patch

bump_minor: ## Release a minor version
	./scripts/bump.sh --minor

bump_major: ## Release a major version
	./scripts/bump.sh --major

version: ## Echo current version
	@cat VERSION