CONTAINER_IMAGE = islavallee/librapi
CONTAINER_TAG ?= latest
GROUP_ID = $(shell id -g)
NAMESPACE = dev
RELEASE_NAME = librapi
USER_ID = $(shell id -u)
USERNAME = jenkins

DCAPI = docker-compose run --rm api
DCDEV = docker-compose -f docker-compose.dev.yml
DCHELM = docker-compose run --rm --user=$(USER_ID):$(GROUP_ID) -e HOME=/$(USERNAME) helm

.SILENT: ;               # no need for @
.ONESHELL: ;             # recipes execute in same shell
.NOTPARALLEL: ;          # wait for this target to finish
.EXPORT_ALL_VARIABLES: ; # send all vars to shell

.PHONY: help

help: ## Show Help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

###############GAZR###############

init: ## Bootstrap your application.
	docker-compose build
	$(DCAPI) go mod download

style: go-lint ## Check lint.
style: helm-lint

complexity: ## Cyclomatic complexity check.
	$(DCAPI) gocyclo .

format: ## Format code.
	$(DCAPI) gofmt -w .

test: test-unit ## Shortcut to launch all the test tasks (unit, functional and integration).
test: test-functional
test: test-integration

test-unit: ## Launch unit tests.
	$(DCAPI) go test -v ./...

test-functional: ## Launch functional tests. e.g behat, JBehave, Behave, CucumberJS, Cucumber etc...
	echo "no functional tests"

test-integration: ## Launch integration tests. e.g. pytest, jest (js), phpunit, JUnit (java) etc ...
	echo "no integration tests"

security-sast: ## launch static application security testing (SAST).
	$(DCAPI) gosec ./...

run: ## Locally run the application.
	docker run --rm -p 80:8080 ${CONTAINER_IMAGE}:${CONTAINER_TAG}

watch: dev-up ## Hot reloading for development.

build: ## Build the application.
	docker build -t ${CONTAINER_IMAGE}:${CONTAINER_TAG} .

###############DEV###############

dev-build: ## DEV ENV - docker-compose build
	$(DCDEV) build

dev-up: ## DEV ENV - docker-compose up
	$(DCDEV) up -d
	$(DCDEV) logs -f

dev-logs: ## DEV ENV - docker-compose logs
	$(DCDEV) logs -f

dev-down: ## DEV ENV - docker-compose down
	$(DCDEV) down

###############GO##############

go-lint: ## GO APP - Check lint on Go app.
	$(DCAPI) golint ./...

###############HELM###############

helm-lint: ## HELM - Check lint on helm chart.
	$(DCHELM) lint librapi/.

helm-template: ## HELM - Compile template
	$(DCHELM) template librapi/.

helm-package: ## HELM - Build helm package
	mkdir -p build/helm/
	$(DCHELM) package librapi/. -d build/helm/

helm-install: ## HELM - Install app
	helm install -n ${NAMESPACE} ${RELEASE_NAME} helm/librapi/.

helm-upgrade:
	helm upgrade -n ${NAMESPACE} ${RELEASE_NAME} helm/librapi/.

helm-ls:
	helm ls -n ${NAMESPACE}

helm-uninstall:
	helm uninstall -n ${NAMESPACE} ${RELEASE_NAME}

###############MICROK8S###############

image-push: ## push local image to microk8s image cache
	mkdir -p build/docker
	docker save -o build/docker/librapi.tar ${CONTAINER_IMAGE}:${CONTAINER_TAG}
	sudo microk8s ctr image import build/docker/librapi.tar
