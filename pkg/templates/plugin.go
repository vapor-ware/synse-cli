// Synse CLI
// Copyright (c) 2019 Vapor IO
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package templates

import "github.com/MakeNowJust/heredoc"

// Templates for generating new plugin projects.
var (
	PluginMainGo = heredoc.Doc(`
package main

import (
	"github.com/{{ }}/{{  }}/pkg"
	"github.com/vapor-ware/synse-sdk/sdk"
)

const (
	pluginName       = "{{ }}"
	pluginMaintainer = "{{ }}"
	pluginDesc       = "{{ }}"
	pluginVcs        = "github.com/{{ }}/{{ }}"
)

func main() {
	// Set the plugin metadata
	sdk.SetPluginInfo(
		pluginName,
		pluginMaintainer,
		pluginDesc,
		pluginVcs,
	)

	// Create the plugin
	plugin := pkg.MakePlugin()

	// Run the plugin
	if err := plugin.Run(); err != nil {
		log.Fatal(err)
	}
}
`)

	PluginDockerfile = heredoc.Doc(`
#
# Builder Image
#
FROM vaporio/golang:1.11 as builder

#
# Final Image
#
FROM scratch

LABEL org.label-schema.schema-version="1.0" \
      org.label-schema.name="{{ }}/{{ }}" \
      org.label-schema.vcs-url="https://github.com/{{ }}/{{ }}" \
      org.label-schema.vendor="{{ }}"

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# Copy the executable.
COPY {{ }} ./plugin

ENTRYPOINT ["./plugin"]
`)

	PluginMakefile = heredoc.Doc(`
#
# {{ }}
#

PLUGIN_NAME    := {{  }}
PLUGIN_VERSION := 0.0.0
IMAGE_NAME     := {{  }}
BIN_NAME       := {{  }}

GIT_COMMIT     ?= $(shell git rev-parse --short HEAD 2> /dev/null || true)
GIT_TAG        ?= $(shell git describe --tags 2> /dev/null || true)
BUILD_DATE     := $(shell date -u +%Y-%m-%dT%T 2> /dev/null)
GO_VERSION     := $(shell go version | awk '{ print $$3 }')

PKG_CTX := github.com/{{ }}/{{ }}/vendor/github.com/vapor-ware/synse-sdk/sdk
LDFLAGS := -w \
	-X ${PKG_CTX}.BuildDate=${BUILD_DATE} \
	-X ${PKG_CTX}.GitCommit=${GIT_COMMIT} \
	-X ${PKG_CTX}.GitTag=${GIT_TAG} \
	-X ${PKG_CTX}.GoVersion=${GO_VERSION} \
	-X ${PKG_CTX}.PluginVersion=${PLUGIN_VERSION}


.PHONY: build
build:  ## Build the plugin binary
	go build -ldflags "${LDFLAGS}" -o ${BIN_NAME}

.PHONY: build-linux
build-linux:  ## Build the plugin binarry for linux amd64
	GOOS=linux GOARCH=amd64 go build -ldflags "${LDFLAGS}" -o ${BIN_NAME} .

.PHONY: clean
clean:  ## Remove temporary files
	go clean -v
	rm -rf dist

.PHONY: dep
dep:  ## Ensure and prune dependencies
	dep ensure -v

.PHONY: deploy
deploy:  ## Run a local deployment of the plugin with Synse Server
	docker-compose -f compose.yml up -d

.PHONY: docker
docker:  ## Build the docker image
	docker build -f Dockerfile \
		--label "org.label-schema.build-date=${BUILD_DATE}" \
		--label "org.label-schema.vcs-ref=${GIT_COMMIT}" \
		--label "org.label-schema.version=${PLUGIN_VERSION}" \
		-t ${IMAGE_NAME}:latest .

.PHONY: fmt
fmt:  ## Run goimports on all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do goimports -w "$$file"; done

.PHONY: github-tag
github-tag:  ## Create and push a tag with the current plugin version
	git tag -a ${PLUGIN_VERSION} -m "${PLUGIN_NAME} plugin version ${PLUGIN_VERSION}"
	git push -u origin ${PLUGIN_VERSION}

.PHONY: lint
lint:  ## Lint project source files
	golint -set_exit_status ./pkg/...

.PHONY: version
version:  ## Print the version of the plugin
	@echo "${PLUGIN_VERSION}"

.PHONY: help
help:  ## Print usage information
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

.DEFAULT_GOAL := help

`)

	PluginPkgPluginGo = heredoc.Doc(`
package pkg

import (
	"log"

	"github.com/{{ }}/{{ }}/pkg/devices"
	"github.com/{{ }}/{{ }}/pkg/outputs"
	"github.com/vapor-ware/synse-sdk/sdk"
)

// MakePlugin creates a new instance of the {{ }}.
func MakePlugin() *sdk.Plugin {
	plugin, err := sdk.NewPlugin()
	if err != nil {
		log.Fatal(err)
	}

	// Register custom output types.
	err = plugin.RegisterOutputs(
		&outputs.Airflow,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Register device handlers
	err = plugin.RegisterDeviceHandlers(
		&devices.Airflow,
		&devices.Fan,
		&devices.Humidity,
		&devices.LED,
		&devices.Pressure,
		&devices.Temperature,
		&devices.Lock,
	)
	if err != nil {
		log.Fatal(err)
	}

	return plugin
}
`)

)


