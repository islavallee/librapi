CONTAINER_IMAGE = slavallee/librapi
CONTAINER_TAG ?= latest
DC = docker-compose run --rm api 
DCDEV = docker-compose -f docker-compose.dev.yml

.SILENT: ;               # no need for @
.ONESHELL: ;             # recipes execute in same shell
.NOTPARALLEL: ;          # wait for this target to finish
.EXPORT_ALL_VARIABLES: ; # send all vars to shell

.PHONY: help

help: ## Show Help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

###############GAZR###############

init: ##  Bootstrap your application.
	docker-compose build
	$(DC) go mod download	

style: ## Check lint.
	$(DC) golint ./...

complexity: ## Cyclomatic complexity check.
	$(DC) gocyclo .

format: ## Format code.
	$(DC) gofmt -w .

test: test-unit ## Shortcut to launch all the test tasks (unit, functional and integration).
test: test-functional
test: test-integration

test-unit: ## Launch unit tests.
	$(DC) go test -v ./...

test-functional: ## Launch functional tests. e.g behat, JBehave, Behave, CucumberJS, Cucumber etc...
	echo "no functional tests"

test-integration: ## Launch integration tests. e.g. pytest, jest (js), phpunit, JUnit (java) etc ...
	echo "no integration tests"

security-sast: ## launch static application security testing (SAST).
	$(DC) gosec ./...

run: ## Locally run the application.
	docker run --rm ${CONTAINER_IMAGE}:${CONTAINER_TAG}

watch: dev-up ## Hot reloading for development.

build: ## Build the application.
	docker build -t ${CONTAINER_IMAGE}:${CONTAINER_TAG} .

###############DEV###############

dev-build:
	$(DCDEV) build

dev-up:
	$(DCDEV) up -d
	$(DCDEV) logs -f

dev-logs:
	$(DCDEV) logs -f

dev-down:
	$(DCDEV) down
