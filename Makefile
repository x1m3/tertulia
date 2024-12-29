BIN := $(shell pwd)/bin
VERSION ?= $(shell git rev-parse --short HEAD)
GO?=$(shell which go)
export GOBIN := $(BIN)
export PATH := $(BIN):$(PATH)

BUILD_CMD := $(GO) install -ldflags "-X main.build=${VERSION}"

INFRASTRUCTURE_PATH = $(shell pwd)/infrastructure/local
DOCKER_COMPOSE_FILE := $(INFRASTRUCTURE_PATH)/docker-compose.yml
DOCKER_COMPOSE_CMD := docker compose -p storagenode -f $(DOCKER_COMPOSE_FILE)
DOCKER_COMPOSE_INFRA_CMD := docker compose -p tertulia -f $(INFRASTRUCTURE_PATH)/docker-compose-infra.yaml

# Set the default environment to "local" if not provided. "local|test|prod"
ENV ?= local

# Environment variables from config file
DB_URL ?= $(shell $(BIN)/godotenv -f $(shell pwd)/config.$(ENV).env printenv DB_URL)

.PHONY: init
init: $(BIN)/golangci-lint $(BIN)/oapi-codegen $(BIN)/goose $(BIN)/godotenv

.PHONY: lint
lint: $(BIN)/golangci-lint
	  $(BIN)/golangci-lint run

.PHONY: test
test:
	$(GO) test -v ./...

.PHONY: build/docker
build/docker: ## Build the docker image.
	DOCKER_BUILDKIT=1 \
	docker build \
		-f ./Dockerfile \
		-t ipfs_proxy_cache:$(VERSION) \
		--build-arg VERSION=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.


.PHONY: run
run:
	COMPOSE_DOCKER_CLI_BUILD=1 $(DOCKER_COMPOSE_CMD) up -d api

## Generate API files
.PHONY: api
api: $(BIN)/oapi-codegen
	$(BIN)/oapi-codegen -config ./api/config-oapi-codegen.yaml ./api/api.yaml > ./internal/interface/http/api.gen.go

migrate-create: $(BIN)/goose
	@if [ -z "$(name)" ]; then \
		echo "Error: _Migration name_ is required. Use it as make migrate-create name=<migration_name>"; \
		exit 1; \
	fi
	$(BIN)/goose -dir ./internal/db/migrations create $(name) sql

migrate-up: $(BIN)/godotenv $(BIN)/goose
	$(BIN)/goose -dir ./internal/db/migrations postgres "$(DB_URL)" up

migrate-down: $(BIN)/godotenv $(BIN)/goose
	$(BIN)/goose -dir ./internal/db/migrations postgres "$(DB_URL)" down

.PHONY: up
up:
	$(DOCKER_COMPOSE_INFRA_CMD) up -d postgres

## install code generator for API files.
$(BIN)/oapi-codegen: tools.go go.mod go.sum
	$(GO) install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

# install golang lint.
$(BIN)/golangci-lint: go.mod go.sum
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61

# install database migration tool. goose
$(BIN)/goose: go.mod go.sum
	go install github.com/pressly/goose/v3/cmd/goose@latest

# install godotenv
$(BIN)/godotenv: tools.go go.mod go.sum
	$(GO) install github.com/joho/godotenv/cmd/godotenv

