<p align="center"><img src="docs/img/logo.png" width="360"></p>
<p align="center">
    <a href="https://app.buddy.works/timfall/synse-cli/pipelines/pipeline/50439"><img src="https://app.buddy.works/timfall/synse-cli/pipelines/pipeline/50439/badge.svg?token=3ae6c804af4fdb5947b58ba1c544c232bf8d28f6e6d2b07321added2d1cc0bad"></a>
    <a href="https://circleci.com/gh/vapor-ware/synse-cli"><img src="https://circleci.com/gh/vapor-ware/synse-cli.svg?style=shield&circle-token=7e11598b349e1d280c7cd78517ababef0f837bc3"></a>
    <a href="http://godoc.org/github.com/vapor-ware/synse-cli"><img src="https://godoc.org/github.com/vapor-ware/synse-cli?status.svg"></a>
    <a href="https://goreportcard.com/report/github.com/vapor-ware/synse-cli"><img src="https://goreportcard.com/badge/github.com/vapor-ware/synse-cli"></a>
        
<h1 align="center">Synse Command Line Interface</h1>
</p>

A CLI for for Vapor IO's [Synse Server][synse-server] and Synse Plugins.


## Overview

[Synse Server](https://github.com/vapor-ware/synse-server) provides a programatic API for bi-directional access to hardware
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

**NOTE**: Windows is not supported at this time, due to incompatibilities in some utilized libraries. If there is
enough interest this may be added in the future.

To get the code and compile it, simply clone this repository into your `GOPATH` and run the standard `go build`
command. For example:

```shell
git clone https://github.com/vapor-ware/synse-cli
cd synse-cli
go get
go build
```

This should produce a binary in the same directory that matches the system architecture that it was built on.
You can also use `go install` to build the binary and install it in your `PATH`.

Alternatively, you can use `go`'s built in package management system to fetch the source code and place it
in your `GOPATH`.

```shell
go get github.com/vapor-ware/synse-cli
cd $GOPATH/src/github.com/vapor-ware/synse-cli
go get
go build
```

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
- Configuration file settings (e.g. `SynseHost: awesome.sauce`)

#### Configuration Options

There are currently a number of configuration options available.

- Synse Host
   - This sets the API endpoint for where synse is serving data. It is given as a resolvable address without
   any leading or trailing information (e.g. awesome.sauce.io).
   - It can be set with the following
      - `--synse_host` or `--host` flag
      - `SYNSE_HOST` environment variable
      - `SynseHost` (case sensitive) in the configuration file
- Debug
   - This enables debug logging to `STDOUT`, printing debug information to the screen.
   - In can be enabled by setting the following
      - `--debug` flag
      - `SYNSE_DEBUG` environment variable
      - `debug: true` in the configuration file
- Configuration File
   - This gives the path for where to locate the configuration file (more details in [Configuration File](#configuration-file)).
   - It can be set with the following
      - `--config` or `-c` flags
      - `SYNSE_CONFIG_FILE` environment variable

#### Configuration File

Configuration options can be read in from a file at runtime. Currently this file is named `.synse.yaml`
(notice the leading dot). By default synse cli will look for this file in two places at runtime, the root
level of the current user's home directory (`~/.synse.yaml`) or, if it is not found there, in the current
directory the command is being run from (`./.synse.yaml`). As mentioned above, specific settings in this
file can be overriden on a per command basis using a higher precedence method.

The configuration file follows standard YAML syntax and accepts the following settings:

- `SynseHost: some.host.com`
- `debug: <true/false>`

Configuration values _are_ case sensitive, but the cli will attempt to decode any values that match the
above keys.

### Contributing

Synse CLI is (un)-lovingly maintained by [timfallmk](https://github.com/timfallmk), who is far over worked and
underpaid. We happily accept issues and pull requests logged in this repository. Please just be nice and follow
appropriate rules when submitting anything.

Any code in this repository is governed under the license given therein.
