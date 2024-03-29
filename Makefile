# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

export PATH := $(abspath bin/protoc/bin/):$(abspath bin/):${PATH}

# Build variables
BUILD_DIR ?= build
VERSION ?= $(shell git describe --tags --exact-match 2>/dev/null || git symbolic-ref -q --short HEAD)
LDFLAGS += -X main.version=${VERSION}
export CGO_ENABLED ?= 0

.PHONY: build
build: ## Build all binaries
	@mkdir -p ${BUILD_DIR}
	go build -trimpath -ldflags "${LDFLAGS}" -o ${BUILD_DIR}/ .

.PHONY: check
check: test lint ## Run checks (tests and linters)

.PHONY: test
test: TEST_FORMAT ?= short
test: export CGO_ENABLED=1
test: ## Run tests
	@mkdir -p ${BUILD_DIR}
	gotestsum --no-summary=skipped --junitfile ${BUILD_DIR}/coverage.xml --jsonfile ${BUILD_DIR}/test.json --format ${TEST_FORMAT} -- -race -coverprofile=${BUILD_DIR}/coverage.txt -covermode=atomic ./...

.PHONY: lint
lint: ## Run linter
	golangci-lint run ${LINT_ARGS}

.PHONY: fmt
fmt: ## Format code
	golangci-lint run --fix

.PHONY: testproto
testproto: build
	protoc -I test --plugin=build/protoc-gen-go-kit --go_out=paths=source_relative:test/ --go-grpc_out=paths=source_relative:test/ --go-kit_out=paths=source_relative:test/ test/test.proto test/subtest/subtest.proto

.PHONY: check-release-config
check-release-config:
	goreleaser check

.PHONY: versions
versions:
	@go version
	@golangci-lint --version
	@gotestsum --version
	@protoc --version

# Dependency versions
GOTESTSUM_VERSION ?= 1.9.0
GOLANGCI_VERSION ?= 1.52.2
PROTOC_VERSION ?= 3.21.12
PROTOC_GEN_GO_VERSION ?= 1.29.1
PROTOC_GEN_GO_GRPC_VERSION ?= 1.3.0

deps: bin/gotestsum bin/golangci-lint bin/protoc bin/protoc-gen-go bin/protoc-gen-go-grpc

bin/gotestsum:
	@mkdir -p bin
	curl -L https://github.com/gotestyourself/gotestsum/releases/download/v${GOTESTSUM_VERSION}/gotestsum_${GOTESTSUM_VERSION}_$(shell uname | tr A-Z a-z)_amd64.tar.gz | tar -zOxf - gotestsum > ./bin/gotestsum
	@chmod +x ./bin/gotestsum

bin/golangci-lint:
	@mkdir -p bin
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | BINARY=golangci-lint bash -s -- v${GOLANGCI_VERSION}

bin/protoc:
	@mkdir -p bin/protoc
ifeq ($(shell uname | tr A-Z a-z), darwin)
	curl -L https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-osx-x86_64.zip > bin/protoc.zip
endif
ifeq ($(shell uname | tr A-Z a-z), linux)
	curl -L https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip > bin/protoc.zip
endif
	unzip bin/protoc.zip -d bin/protoc
	rm bin/protoc.zip

bin/protoc-gen-go:
	@mkdir -p bin
	curl -L https://github.com/protocolbuffers/protobuf-go/releases/download/v${PROTOC_GEN_GO_VERSION}/protoc-gen-go.v${PROTOC_GEN_GO_VERSION}.$(shell uname | tr A-Z a-z).amd64.tar.gz | tar -zOxf - protoc-gen-go > ./bin/protoc-gen-go
	@chmod +x ./bin/protoc-gen-go

bin/protoc-gen-go-grpc:
	@mkdir -p bin
	curl -L https://github.com/grpc/grpc-go/releases/download/cmd/protoc-gen-go-grpc/v${PROTOC_GEN_GO_GRPC_VERSION}/protoc-gen-go-grpc.v${PROTOC_GEN_GO_GRPC_VERSION}.$(shell uname | tr A-Z a-z).amd64.tar.gz | tar -zOxf - ./protoc-gen-go-grpc > ./bin/protoc-gen-go-grpc
	@chmod +x ./bin/protoc-gen-go-grpc

.PHONY: help
.DEFAULT_GOAL := help
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-10s\033[0m %s\n", $$1, $$2}'
