
VERSION := $(shell cat cmd/synse.go | grep 'appVersion =' | awk '{print $$3}')

HAS_LINT := $(shell which gometalinter)
HAS_DEP  := $(shell which dep)

.PHONY: build
build:  ## Build the CLI locally
	go build -o build/synse github.com/vapor-ware/synse-cli/cmd

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
	rm coverage.txt

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
	gometalinter ./... --vendor --tests --deadline=5m \
		--disable=gotype

.PHONY: setup
setup:  ## Install the build and development dependencies
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/alecthomas/gometalinter
	go get -u golang.org/x/tools/cmd/cover
	gometalinter --install
	@$(MAKE) dep

.PHONY: test
test:  ## Run all tests
	go test -cover ./...

.PHONY: ci-test
ci-test:
	go test -v ./... 2>&1 | tee /tmp/${TEST_DIRECTORY}/test.out
	cat /tmp/${TEST_DIRECTORY}/test.out \
		| go-junit-report \
		> /tmp/${TEST_DIRECTORY}/report.xml

.PHONY: ci-create-release
ci-create-release:
	$(eval GIT_TAG := $(shell git describe --exact-match --tags HEAD))
	ghr \
		-u ${GITHUB_USER} \
		-t ${GITHUB_TOKEN} \
		-b "$(cat ./CHANGELOG.md)" \
		-p 1 \
		-draft \
		${GIT_TAG} build/

.PHONY: version
version: ## Print the version of the CLI
	@echo "$(VERSION)"

.PHONY: help
help:  ## Print usage information
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

.DEFAULT_GOAL := help
