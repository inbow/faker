# Common
ORG := oxyd-io
NAME := $(shell basename $(CURDIR))
VERSION := $(shell git tag --points-at HEAD --sort -version:refname | head -1)
COMMIT := $(shell git rev-parse --short HEAD)

# Build
GO_PACKAGE := github.com/${ORG}/${NAME}
GC_FLAGS := -gcflags 'all=-N -l'
LD_FLAGS := -ldflags '-s -v -w -X main.service=${NAME} -X main.version=${COMMIT}'
BUILD_CMD := CGO_ENABLED=0 go build -o bin/${NAME} ${LD_FLAGS} ${GO_PACKAGE}/cmd/${NAME}
DEBUG_CMD := CGO_ENABLED=0 go build -o bin/${NAME} ${GC_FLAGS} ${LD_FLAGS} ${GO_PACKAGE}/cmd/${NAME}

# Docker
REGISTRY_URL := docker.pkg.github.com
DOCKER_IMAGE_NAME := ${REGISTRY_URL}/${ORG}/${NAME}/${NAME}
DOCKER_APP_FILENAME := deployments/docker/Dockerfile
DOCKER_COMPOSE_FILE := deployments/docker-compose/docker-compose.yml

# Other
.DEFAULT_GOAL := build
THIS_FILE := $(lastword $(MAKEFILE_LIST))

.PHONY: linters
linters:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint
	golangci-lint run --enable-all --disable gomnd --disable dupl --disable gochecknoglobals --disable gofumpt

.PHONY: tests
tests:
	@echo "Without race"
	@go test -v ./...
	@echo "With race"
	@go test -v -race ./...

.PHONY: clean
clean:
	@echo "> Cleaning binaries for ${NAME}"
	@-rm -rf bin/${NAME}

.PHONY: build
build: clean
	${BUILD_CMD}

.PHONY: build_debug
build_debug: clean
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
