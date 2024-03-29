# Common
VCS = github.com
ORG = oxyd-io
NAME = faker
COMMIT := $(shell git rev-parse --short HEAD)
VERSION := $(shell git tag --points-at HEAD)
ACTION ?= build

# Build
GO_PACKAGE = ${VCS}/${ORG}/${NAME}
GC_FLAGS = -gcflags 'all=-N -l'
LD_FLAGS = -ldflags '-s -v -w -X main.commit=${COMMIT} -X main.version=${VERSION}'
BUILD_CMD = CGO_ENABLED=0 go build -o bin/${NAME} ${LD_FLAGS} ${GO_PACKAGE}/cmd/${NAME}
DEBUG_CMD = CGO_ENABLED=0 go build -o bin/${NAME} ${GC_FLAGS} ${GO_PACKAGE}/cmd/${NAME}

# Docker
REGISTRY_URL = docker.pkg.github.com
DOCKER_IMAGE_NAME = ${REGISTRY_URL}/${ORG}/${NAME}/${NAME}
DOCKER_APP_FILENAME = deployments/docker/Dockerfile
DOCKER_GOLANG_IMAGE = docker.pkg.github.com/oxyd-io/docker-images/go-image:1.15

# Other
.DEFAULT_GOAL = build
THIS_FILE = $(lastword $(MAKEFILE_LIST))

.PHONY: api
api:
	protoc -I. \
		-I/usr/local/include \
		--go_out=plugins=grpc:. \
		--go_opt=paths=source_relative \
		api/*.proto

.PHONY: lint
lint:
	@which golangci-lint &>/dev/null || GO111MODULE=off go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	golangci-lint run --enable-all --disable gomnd --disable dupl --disable gochecknoglobals --disable gofumpt

.PHONY: tests
tests:
	@echo "Run tests"
	@go test -v -race ./...

.PHONY: tidy
tidy: GOPRIVATE=${VCS}/${ORG}/*
tidy:
	go mod tidy
	go mod vendor

.PHONY: clean
clean:
	@echo "> Cleaning binaries for ${NAME}"
	@-rm -rf bin/${NAME}

.PHONY: build
build: clean
	@echo "Build: ${NAME}"
	${BUILD_CMD}

.PHONY: build_debug
build_debug: clean
	@echo "Build debug: ${NAME}"
	${DEBUG_CMD}

.PHONY: docker_build
docker_build:
	docker run \
		-v `pwd`:/go/src/${GO_PACKAGE} \
		-w /go/src/${GO_PACKAGE} \
		-e 'ACTION=${ACTION}' \
		-i ${DOCKER_GOLANG_IMAGE} \
		/bin/sh -c "${BUILD_CMD}"

.PHONY: docker_build_local
docker_local_push:
	docker build -t ${NAME}:local -f ${DOCKER_APP_FILENAME} ACTION=${ACTION} .

.PHONY: docker_build_push_versioned
docker_build_push_image:
	docker build -t ${DOCKER_IMAGE_NAME}:${COMMIT} -f ${DOCKER_APP_FILENAME} --build-arg ACTION=${ACTION} .
	docker push ${DOCKER_IMAGE_NAME}:${COMMIT}
