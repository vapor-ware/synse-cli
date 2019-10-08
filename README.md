# Synse CLI

[![Build Status](https://build.vio.sh/buildStatus/icon?job=vapor-ware/synse-cli/master)](https://build.vio.sh/blue/organizations/jenkins/vapor-ware%2Fsynse-cli/activity)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fvapor-ware%2Fsynse-cli.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fvapor-ware%2Fsynse-cli?ref=badge_shield)
[![Go ReportCard](https://goreportcard.com/badge/github.com/vapor-ware/synse-cli)](https://goreportcard.com/report/github.com/vapor-ware/synse-cli)

A command-line interface for Vapor IO's [Synse platform](https://github.com/vapor-ware/synse).

## Overview

The `synse` CLI provides a simple but comprehensive tool to interact with [Synse Server](https://github.com/vapor-ware/synse-server)
(via its HTTP API) and Synse plugins (via the internal [gRPC API](https://github.com/vapor-ware/synse-server-grpc)).
It allows for real-time queries and interaction with devices exposed by Synse. This makes getting started
with Synse easy, and enabled rapid debugging and development against various Synse components.

## Installing

### Homebrew

The Synse CLI may be install via [Homebrew](https://brew.sh/). First, add the vapor-ware tap

```
brew tap vapor-ware/formula
```

Then, you can install the CLI

```
brew install vapor-ware/formula/synse
```

### Precompiled Binaries

Precompiled binaries are available as artifacts on GitHub [releases](https://github.com/vapor-ware/synse-cli/releases).
To download the binary and place it on your $PATH:

```shell
# Set variables for download
export CLI_VERSION="3.0.0"
export CLI_OS="darwin"
export CLI_ARCH="amd64"

# Download and install the CLI
wget \
  https://github.com/vapor-ware/synse-cli/releases/download/${CLI_VERSION}/synse-cli_${CLI_VERSION}_${CLI_OS}_${CLI_ARCH}.tar.gz \
  -O /usr/local/bin/synse

# Make the binary executable
chmod +x /usr/local/bin/synse
``` 

### From Source

If you wish to build from source, you will first need to fork and clone the repo. From within the
project directory, you can build using the Makefile target:

```
make build
```

Which will create the `synse` binary in the project directory. If you wish, you can add it to
your PATH.


## Getting Started

With the CLI installed, you can run `synse --help` to get usage info. You can get additional
info on all commands and sub-commands by running the command with the `--help` flag.

There are three primary commands to be aware of:

- `context`: Configuration management for server/plugin instances.
- `server`: Interact with a Synse Server instance via HTTP.
- `plugin`: Interact with a plugin instance via gRPC.


### Contexts

Prior to interacting with a server or plugin instance, a new context for it needs to be created.
If running the [example deployment](synse.yaml) found in this repo (which runs Synse Server at
localhost:5000 and the emulator plugin at localhost:5001), this can be done with:

```bash
# Add a server context and set it as the current server.
synse context add server local localhost:5000 --set

# Add a plugin context and set it as the current plugin.
synse context add plugin emulator localhost:5001 --set
``` 

You can then list the contexts and see that those are both present and marked as active.
Now when you run a `synse server ...` or `synse plugin ...` command, it knows which instance
to communicate with.

```console
$ synse context list
CURRENT   NAME       TYPE     ADDRESS
*         emulator   plugin   localhost:5001
*         local      server   localhost:5000
```


## License

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fvapor-ware%2Fsynse-cli.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fvapor-ware%2Fsynse-cli?ref=badge_large)
