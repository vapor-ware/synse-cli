<p align="center"><img src="assets/logo.png" width="360"></p>
<p align="center">
    <a href="https://app.buddy.works/timfall/synse-cli/pipelines/pipeline/50439"><img src="https://app.buddy.works/timfall/synse-cli/pipelines/pipeline/50439/badge.svg?token=3ae6c804af4fdb5947b58ba1c544c232bf8d28f6e6d2b07321added2d1cc0bad"></a>
    <a href="https://circleci.com/gh/vapor-ware/synse-cli"><img src="https://circleci.com/gh/vapor-ware/synse-cli.svg?style=shield&circle-token=7e11598b349e1d280c7cd78517ababef0f837bc3"></a>
    <a href="http://godoc.org/github.com/vapor-ware/synse-cli"><img src="https://godoc.org/github.com/vapor-ware/synse-cli?status.svg"></a>
    <a href="https://goreportcard.com/report/github.com/vapor-ware/synse-cli"><img src="https://goreportcard.com/badge/github.com/vapor-ware/synse-cli"></a>
        
<h1 align="center">Synse Command Line Interface</h1>
</p>

<p align="center">A CLI for for Vapor IO's Synse Server and Synse Plugins.</p>

## Overview

[Synse Server](https://github.com/vapor-ware/synse-server) provides a programmatic API for bi-directional access to hardware
and other components via its configured plugins. For more information, see the [README](https://github.com/vapor-ware/synse-server/blob/master/README.md)
for Synse Server.

The Synse CLI provides a command line interface to Synse Server instances as well as direct access to plugins.
It allows for real-time queries and interaction with the hardware that is managed by Synse Server, and is meant
to have feature parity with the Synse Server HTTP API. Although it is not designed to be programmed against
directly (for that, the Synse Server HTTP API should be used), it can be used as a tool for local access and
debugging of Synse components and may serve as an example on how to interface other applications with Synse Server.


## Quick Start

The Synse CLI is provided as a single pre-compiled binary, available for use on most platforms. To get it, simply
download the binary that matches your architecture (e.g. x86 or amd64).

For example, if running on macOS, you could do the following:

```shell
wget https://github.com/vapor-ware/synse-cli/releases/download/<version number>/synse_darwin_amd64 -O /usr/local/bin/synse
chmod +x /usr/local/bin/synse
```

Or:

```shell
curl https://github.com/vapor-ware/synse-cli/releases/download/<version number>/synse_darwin_amd64 -o /usr/local/bin/synse
chmod +x /usr/local/bin/synse
```

To save the binary to `/usr/local/bin/`. The binary can be saved to any location you wish.

> **NOTE**: We recommend that you save the binary with the name `synse` as it is easier to remember. All commands
in this guide will use this name.

You are now ready to run `synse` cli. To see a list of commands, type `synse -h`.

## Slightly Longer Way

### Building from Source

Synse CLI is written in [Go](https://golang.org) and is provided as source. It can be compiled directly from source
and built into a binary for any supported system.

> **NOTE**: Windows is not supported at this time, due to incompatibilities in some utilized libraries. If there is
enough interest this may be added in the future.

To get the code and compile it, clone this repository into your `GOPATH`. Then, get the vendored dependencies
with [dep](https://github.com/golang/dep) and build.

```shell
git clone https://github.com/vapor-ware/synse-cli
cd synse-cli
dep ensure
go build
```

This should produce a binary in the same directory that matches the system architecture that it was built on.
You can also use `go install` to build the binary and install it in your `PATH`.

Alternatively, you can use `go`'s built in package management system to fetch the source code and place it
in your `GOPATH`.

```shell
go get github.com/vapor-ware/synse-cli
cd $GOPATH/src/github.com/vapor-ware/synse-cli
dep ensure
go build
```

Makefile targets exist to make it easy to get dependencies, build, lint, format and test. Once you have the
source, either by `git clone` or `go get`, you can simply

```shell
make setup build
```

This will install `dep` and `gometalinter`, if not already installed, ensure the vendored dependencies
exist, then build the CLI and output a `synse` binary to the `build/` subdirectory. From there, it can
be moved onto your `PATH`.

### Running Commands

The Synse CLI is built to run commands with multiple verbs (`git`-like). Each command has it's own help
documentation which shows what it does and how it should be used. A more detailed description of command structure
and how they are used is available in the [documentation](http://godoc.org/github.com/vapor-ware/synse-cli).

For example the help output for the top-level command is printed by running `synse -h` or simply `synse`.

There are different groups of commands. In general, those are:
- configuration management, e.g. managing Synse Server instances to interface with
- Synse Server commands, e.g. reading, writing, etc. via the Synse Server HTTP API
- Synse Plugin commands, e.g. reading, writing, etc. via the internal gRPC API

### Configuration

Configuration options for customizing synse CLI can be input in a number of different ways. Synse CLI uses a
standard cascading order of precedence when evaluating configuration options. Options with the highest
priority are first.

- Command line flags (e.g. `--debug`)
- Environment variables (e.g. `SYNSE_DEBUG`)
- Configuration file settings (e.g. `debug: true`)

#### Configuration Options

The default configuration for the Synse CLI looks like:

```yaml
debug: false
activehost:
  name: local
  address: localhost:5000
hosts:
  local:
    name: local
    address: localhost:5000
```

When running the CLI, it will parse configuration options in the precedence order listed above. When
looking for configuration files, it will first look for the `.synse.yml` (note the leading dot) file
in the current directory the command is being run from (e.g. `.`). If not found there, it will look
in the current user's home directory (e.g. `~`). If the file does not exist in either location, it will
be created in the home directory upon command termination.

- `debug` enables debug logging to `STDOUT`, printing debug information to the screen.
- `activehost` specifies which of the `hosts` is currently set as active. This host will be the one
   used when issuing any subsequent command against Synse Server. This value can either be updated
   manually or via the `hosts change` command.
- `hosts` specifies a list of all Synse Server hosts which the user can choose to interface with.
  by default a host named `local` at address `localhost:5000` is added and set as the active host.
  Hosts can be added manually or by the `hosts add` command.

### Bash Completion

Bash/Zsh completion can be setup for the Synse CLI using the `completion` command, e.g.
```console
$ # completion for bash -- updates .bashrc
$ synse completion bash

$ # completion for zsh -- updates .zshrc
$ synse completion zsh
```

### Contributing

We happily accept issues and pull requests logged in this repository. Please just be nice and follow
appropriate rules when submitting anything.

Any code in this repository is governed under the license given therein.
