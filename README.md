[![buddy pipeline](https://app.buddy.works/timfall/synse-cli/pipelines/pipeline/50439/badge.svg?token=3ae6c804af4fdb5947b58ba1c544c232bf8d28f6e6d2b07321added2d1cc0bad "buddy pipeline")](https://app.buddy.works/timfall/synse-cli/pipelines/pipeline/50439)
[![CircleCI](https://circleci.com/gh/vapor-ware/synse-cli.svg?style=shield&circle-token=7e11598b349e1d280c7cd78517ababef0f837bc3)](https://circleci.com/gh/vapor-ware/vesh)
[![GoDoc](https://godoc.org/github.com/vapor-ware/synse-cli?status.svg)](http://godoc.org/github.com/vapor-ware/synse-cli)

# Synse Command Line Interface

## Overview

[Synse](https://github.com/vapor-ware/synse-server) provides a programatic API for bi-directional access to hardware and other components. For more information, see the [README](https://github.com/vapor-ware/synse-server/blob/master/README.md) in that repository.

Synse CLI provides an command line interface to the underlying synse components. It allows for real-times queries and interaction with hardware endpoints monitored by synse, and is meant to have feature parity with the synse API. Although it is not designed to be programmed against directly, it can be used as an example of how to interface other apps with synse.

To get started, follow the instructions below.

## Quick Start

Synse CLI is provided as a single pre-compiled binary, available for use on most platforms. To get it, simply download the binary that matches your architecture (e.g. x86 or amd64).

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

**NOTE**: We recommend that you save the binary with the name `synse` as it is easier to remember. All commands in this guide will use this name.

You are now ready to run `synse` cli. To see a list of commands, type `synse -h`.

## Slightly Longer Way

### Building from Source

Synse CLI is written in [Go](https://golang.org) and is provided as source. It can be compiled directly from source and built into a binary for any supported system.

**NOTE**: Windows is not supported at this time, due to incompatibilities in some utilized libraries. If there is enough interest this may be added in the future.

To get the code and compile it, simply clone this repository into your `GOPATH` and run the standard `go build` command. For example:

```shell
git clone https://github.com/vapor-ware/synse-cli
cd synse-cli
go get
go build
```

This should produce a binary in the same directory that matches the system architecture that it was built on. You can also use `go install` to build the binary and install it in your `PATH`.

Alternatively, you can use `go`'s built in package management system to fetch the source code and place it in your `GOPATH`.

```shell
go get github.com/vapor-ware/synse-cli
cd $GOPATH/src/github.com/vapor-ware/synse-cli
go get
go build
```

### Running Commands

Synse cli is built to run commands with multiple verbs (`git`-like). Each command has it's own help documentation which shows what it does and how it should be used. A more detailed description of command structure and how they are used is available in the [documentation](http://godoc.org/github.com/vapor-ware/synse-cli).

For example the help output for the top-level command is printed by running `synse -h` or simply `synse`:

```shell
NAME:
   synse - Synse Shell

USAGE:
   synse [global options] command [command options] [arguments...]

VERSION:
   0.0.1

AUTHORS:
   Tim Fall <tim@vapor.io>
   Thomas Rampelberg <thomasr@vapor.io>

COMMANDS:
     status, stat  Get the status of the current deployment
     scan          Scan the infrastructure and display device summary
     assets        Manage and get information about physical devices
     help, h       Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --debug, -d                                 Enable debug mode
   --config file, -c file                      Path to config file [$SYNSE_CONFIG_FILE]
   --synse-host Synse Host, --host Synse Host  Address of Synse Host [$SYNSE_HOST]
   --help, -h                                  show help
   --version, -v                               print the version
```

Possible following verbs are shown under `COMMANDS` and will generate their own help pages down the line.

An example of a complete command would be:

```shell
synse assets temperature list
```

This would list the output of all temperature sensors for all devices known to synse.

### Configuration

Configuration options for customizing synse CLI can be input in a number of different ways. Synse CLI uses a standard cascading order of precedence when evaluating configuration options. Options with the highest priority are first.

- Command line flags (e.g. `--debug`)
- Environment variables (e.g. `SYNSE_DEBUG`)
- Configuration file settings (e.g. `SynseHost: awesome.sauce`)

#### Configuration Options

There are currently a number of configuration options available.

- Synse Host
   - This sets the API endpoint for where synse is serving data. It is given as a resolvable address without any leading or trailing information (e.g. awesome.sauce.io).
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
   - This gives the path for where to locate the configuration file (more details in [Configuration File](#configuration)).
   - It can be set with the following
      - `--config` or `-c` flags
      - `SYNSE_CONFIG_FILE` environment variable

#### Configuration File

Configuration options can be read in from a file at runtime. Currently this file is named `.synse.yaml` (notice the leading dot). By default synse cli will look for this file in two places at runtime, the root level of the current user's home directory (`~/.synse.yaml`) or, if it is not found there, in the current directory the command is being run from (`./.synse.yaml`). As mentioned above, specific settings in this file can be overriden on a per command basis using a higher precedence method.

The configuration file follows standard YAML syntax and accepts the following settings:

- `SynseHost: some.host.com`
- `debug: <true/false>`

Configuration values _are_ case sensitive, but the cli will attempt to decode any values that match the above keys.

### Contributing

Synse CLI is (un)-lovingly maintained by [timfallmk](https://github.com/timfallmk), who is far over worked and underpaid. We happily accept issues and pull requests logged in this repository. Please just be nice and follow appropriate rules when submitting anything.

Any code in this repository is governed under the license given therein.
