# Common
ORG := oxyd-io
NAME := $(shell basename $(CURDIR))
VERSION := $(shell git tag --points-at HEAD --sort -version:refname | head -1)
COMMIT := $(shell git rev-parse --short HEAD)

# Build
GO_PACKAGE := github.com/${ORG}/${NAME}
BUILD_CMD := CGO_ENABLED=0 go build -o bin/${NAME} -ldflags '-v -w -s -X main.version=${VERSION}' ./cmd/${NAME}
DEBUG_CMD := CGO_ENABLED=0 go build -o bin/${NAME} -gcflags "all=-N -l" -ldflags '-X main.version=${COMMIT}' ./cmd/${NAME}

# Docker
REGISTRY_URL := docker.pkg.github.com
DOCKER_IMAGE_NAME := ${REGISTRY_URL}/${ORG}/${NAME}/${NAME}
DOCKER_APP_FILENAME := deployments/docker/Dockerfile
DOCKER_COMPOSE_FILE := deployments/docker-compose/docker-compose.yml

# Other
.DEFAULT_GOAL := build
THIS_FILE := $(lastword $(MAKEFILE_LIST))

.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor

.PHONY: linters
linters:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint
	golangci-lint run --enable-all --disable gomnd --disable dupl --disable gochecknoglobals

.PHONY: tests
tests: linters
	@go test -v ./...

.PHONY: clean
clean:
	@echo "> Cleaning binaries for ${NAME}"
	@-rm -rf bin/${NAME}

.PHONY: build
build: vendor clean
	${BUILD_CMD}

.PHONY: build_debug
build_debug: vendor clean
	${DEBUG_CMD}

.PHONY: docker_local_push
docker_local_push:
	docker build -f ${DOCKER_APP_FILENAME} -t ${NAME} .

.PHONY: docker_build_push_image
docker_build_push_image:
	docker build -t ${DOCKER_IMAGE_NAME}:${VERSION} -f ${DOCKER_APP_FILENAME} --build-arg NAME=${NAME} --build-arg VERSION=${VERSION} .
	docker push ${DOCKER_IMAGE_NAME}:${VERSION}
	docker image rm ${DOCKER_IMAGE_NAME}:${VERSION}
	docker image rm ${DOCKER_IMAGE_NAME}:latest
