package server

import (
	"github.com/urfave/cli"
)

// ServerCommand is the CLI command for interacting with Synse Server.
var ServerCommand = cli.Command{
	Name:  "server",
	Usage: "Interact with Synse Server",

	Subcommands: []cli.Command{
		configCommand,
		infoCommand,
		pluginsCommand,
		readCommand,
		scanCommand,
		statusCommand,
		transactionCommand,
		versionCommand,
		writeCommand,
	},
}
