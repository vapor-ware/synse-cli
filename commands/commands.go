/*
Package commands defines the commands, sub-commands, and flags used in app.Cli to form
the structure of the CLI. Definitions, usage strings, help text, and flags are
delegated to app.Cli. The `Action:` field gives the function called when each
command is run.
*/
package commands

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/commands/hosts"
	"github.com/vapor-ware/synse-cli/commands/plugin"
	"github.com/vapor-ware/synse-cli/commands/server"
)

// Commands provides the global list of commands used by the CLI.
var Commands = []cli.Command{
	hosts.HostsCommand,
	plugin.PluginCommand,
	server.StatusCommand,
	server.VersionCommand,
	server.ScanCommand,
	server.ConfigCommand,
	server.ReadCommand,
	server.WriteCommand,
	server.InfoCommand,
	server.TransactionCommand,
	completionCommand,
}
