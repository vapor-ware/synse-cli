#
# Synse CLI
#

BIN_NAME    := synse
BIN_VERSION ?= local

GIT_COMMIT ?= $(shell git rev-parse --short HEAD 2> /dev/null || true)
GIT_TAG    ?= $(shell git describe --tags 2> /dev/null || true)
BUILD_DATE := $(shell date -u +%Y-%m-%dT%T 2> /dev/null)
GO_VERSION := $(shell go version | awk '{ print $$3 }')

PKG_CTX := github.com/vapor-ware/synse-cli/pkg/version
LDFLAGS := -w \
	-X ${PKG_CTX}.BuildDate=${BUILD_DATE} \
	-X ${PKG_CTX}.GitCommit=${GIT_COMMIT} \
	-X ${PKG_CTX}.GitTag=${GIT_TAG} \
	-X ${PKG_CTX}.GoVersion=${GO_VERSION} \
	-X ${PKG_CTX}.VersionString=${BIN_VERSION}

HAS_LINT := $(shell which gometalinter)
HAS_DEP  := $(shell which dep)
HAS_GOX  := $(shell which gox)


.PHONY: build
build:  ## Build the CLI locally
	go build \
		-o build/${BIN_NAME} \
		-ldflags "$(LDFLAGS)" \
		github.com/vapor-ware/synse-cli/cmd

.PHONY: build-ci
build-ci: ## Build binaries in CI
ifndef HAS_GOX
	go get -v github.com/mitchellh/gox
endif
	gox --ldflags "${LDFLAGS}" --parallel=10 -osarch='!darwin/386' --output="build/${BIN_NAME}_{{ .OS }}_{{ .Arch }}" ${OPTS} ./cmd/...


.PHONY: ci
ci:  ## Run CI checks locally (build, test, lint)
	@$(MAKE) build test lint

.PHONY: clean
clean:  ## Remove temporary files
	go clean -v

.PHONY: cover
cover:  ## Run tests and open the coverage report
	./bin/coverage.sh
	go tool cover -html=coverage.txt

.PHONY: dep
dep:  ## Ensure and prune dependencies
ifndef HAS_DEP
	go get -u github.com/golang/dep/cmd/dep
endif
	dep ensure -v

.PHONY: fmt
fmt:  ## Run goimports on all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do goimports -w "$$file"; done

.PHONY: lint
lint:  ## Lint project source files
ifndef HAS_LINT
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install
endif
	@ # disable gotype: https://github.com/alecthomas/gometalinter/issues/40
	gometalinter ./... \
		--disable=gotype \
		--disable=errcheck \
		--tests \
		--vendor \
		--sort=severity \
		--aggregate \
		--deadline=5m

.PHONY: setup
setup:  ## Install the build and development dependencies and set up vendoring
	go get -u github.com/alecthomas/gometalinter
	go get -u github.com/golang/dep/cmd/dep
	gometalinter --install
ifeq (,$(wildcard ./Gopkg.toml))
	dep init
endif
	@$(MAKE) dep

.PHONY: test
test:  ## Run all tests
	go test -race -cover ./...

.PHONY: version
version: ## Print the version of the CLI
	@echo "$(BIN_VERSION)"

.PHONY: help
help:  ## Print usage information
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

.DEFAULT_GOAL := help
