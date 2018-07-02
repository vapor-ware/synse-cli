package server

import (
	"github.com/urfave/cli"
)

const (
	cmdName = "server"

	cmdUsage = "Interact with Synse Server"

	cmdDescription = `This sub-command allows you to interact with a Synse Server
  instance. The instance that is being interfaced with is set
  and managed by the 'synse hosts' commands.`
)

// ServerCommand is the CLI command for interacting with Synse Server.
var ServerCommand = cli.Command{
	Name:        cmdName,
	Usage:       cmdUsage,
	Description: cmdDescription,

	Subcommands: []cli.Command{
		configCommand,
		infoCommand,
		capabilitiesCommand,
		pluginsCommand,
		readCommand,
		scanCommand,
		statusCommand,
		transactionCommand,
		versionCommand,
		writeCommand,
	},
}
