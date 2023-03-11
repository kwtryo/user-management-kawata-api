.PHONY: help build build-local up down logs ps test
.DEFAULT_GOAL := help

DOCKER_TAG := latest
build: ## Build docker image to deploy
	docker build -t kwtryo/portfolio-api:${DOCKER_TAG} \
		--target deploy ./

build-local: ## Build docker image to local development
	docker compose build --no-cache

up: ## Do docker compose up
	docker compose up

down: ## Do docker compose down
	docker compose down

logs: ## Tail docker compose logs
	docker compose logs -f

ps: ## Check container status
	docker compose ps

test: ## Execute tests
	go test -race -shuffle=on ./...

dry-migrate: ## Try migration
	psqldef --dry-run --host=docker.for.mac.localhost --port=33306 --user=user_management --password=user_management user_management < ./_tools/postgres/init.sql

migrate:  ## Execute migration
	psqldef --host=docker.for.mac.localhost --port=33306 --user=user_management --password=user_management user_management < ./_tools/postgres/init.sql

get-schema:
	psqldef --host=docker.for.mac.localhost --port=33306 --user=user_management --password=user_management user_management --export --export > ./_tools/postgres/schema.sql

generate: ## Generate codes
	go generate ./...

help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'