.PHONY: help build build-local up down logs ps test migrate dry-migrate generate
.DEFAULT_GOAL := help

DOCKER_TAG := latest
build: ## Build docker image to deploy
	docker build -t fujiiiiih/gotodo:$(DOCKER_TAG) \
		--target deploy ./

build-local: ## Build docker image to local development
	docker compose build --no-cache

up: ## Do docker compose up with hot reload
	docker compose up -d

down: ## Do docker compose down
	docker compose down

logs: ## Tail docker compose logs
	docker compose logs -f

ps: ## Check conteiner status
	docker compose ps

test: ## Execute tests
	go test -race -shuffle=on ./...
migrate: ## Migrate database
	mysqldef -u lab -p lab -h 127.0.0.1 -P 33306 lab < ./_tools/mysql/schema.sql 

dry-migrate: ## Migrate database dry
	mysqldef -u lab -p lab -h 127.0.0.1 -P 33306 lab --dry-run < ./_tools/mysql/schema.sql

generate: ## Generate codes
	go generate ./...

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
