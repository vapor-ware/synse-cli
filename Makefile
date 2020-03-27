#
# Synse CLI
#

BIN_NAME    := synse
BIN_VERSION := 3.0.0

GIT_COMMIT  ?= $(shell git rev-parse --short HEAD 2> /dev/null || true)
GIT_TAG     ?= $(shell git describe --tags 2> /dev/null || true)
BUILD_DATE  := $(shell date -u +%Y-%m-%dT%T 2> /dev/null)
GO_VERSION := $(shell go version | awk '{ print $$3 }')

LDFLAGS := -w \
	-X github.com/vapor-ware/synse-cli/pkg.BuildDate=${BUILD_DATE} \
	-X github.com/vapor-ware/synse-cli/pkg.Commit=${GIT_COMMIT} \
	-X github.com/vapor-ware/synse-cli/pkg.Tag=${GIT_TAG} \
	-X github.com/vapor-ware/synse-cli/pkg.GoVersion=${GO_VERSION} \
	-X github.com/vapor-ware/synse-cli/pkg.Version=${BIN_VERSION}


.PHONY: build
build:  ## Build the executable binary
	CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags "${LDFLAGS}" -o ${BIN_NAME} cmd/synse.go

.PHONY: clean
clean:  ## Remove temporary files and build artifacts
	go clean -v
	rm -rf dist
	rm -f ${BIN_NAME} coverage.out

.PHONY: cover
cover: test  ## Run unit tests and open the coverage report
	go tool cover -html=coverage.out

.PHONY: fmt
fmt:  ## Run goimports on all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do goimports -w "$$file"; done

.PHONY: github-tag
github-tag:  ## Create and push a tag with the current version
	git tag -a ${BIN_VERSION} -m "synse-cli v${BIN_VERSION}"
	git push -u origin ${BIN_VERSION}

.PHONY: lint
lint:  ## Lint project source files
	golint -set_exit_status ./cmd/... ./pkg/...

.PHONY: test
test:  ## Run unit tests
	@ # Note: this requires go1.10+ in order to do multi-package coverage reports
	go test -race -coverprofile=coverage.out -covermode=atomic ./...

.PHONY: version
version: ## Print the version
	@echo "${BIN_VERSION}"

.PHONY: help
help:  ## Print usage information
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

.DEFAULT_GOAL := help
