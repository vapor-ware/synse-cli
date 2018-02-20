package commands

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/commands/hosts"
	"github.com/vapor-ware/synse-cli/pkg/commands/plugin"
	"github.com/vapor-ware/synse-cli/pkg/commands/server"
)

// Commands provides the global list of commands used by the CLI.
var Commands = []cli.Command{
	hosts.HostsCommand,
	plugin.PluginCommand,
	server.StatusCommand,
	server.VersionCommand,
	server.ScanCommand,
	server.ConfigCommand,
	server.PluginsCommand,
	server.ReadCommand,
	server.WriteCommand,
	server.InfoCommand,
	server.TransactionCommand,
	completionCommand,
}
