/*
Package commands defines the commands, sub-commands, and flags used in app.Cli to form
the structure of the CLI. Definitions, usage strings, help text, and flags are
delegated to app.Cli. The `Action:` field gives the function called when each
command is run.
*/
package commands

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/commands/hosts"
	"github.com/vapor-ware/synse-cli/commands/plugin"
	"github.com/vapor-ware/synse-cli/commands/server"
	"github.com/vapor-ware/synse-cli/utils"
)

// Commands provides the global list of commands used by the CLI.
var Commands = []cli.Command{
	configCommand,
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
	{
		Name:   "shell-completion",
		Usage:  "Generate shell completion scripts for bash or zsh",
		Hidden: true,
		Action: func(c *cli.Context) error {
			switch {
			case c.IsSet("bash") && c.IsSet("zsh"):
				// return utils.CommandHandler(c, utils.GenerateShellCompletion)
				fmt.Println("Can't do both") // FIXME: Once we figure out how to handle this
				return nil
			case c.IsSet("bash"):
				return utils.CmdHandler(utils.GenerateShellCompletion("bash"))
			case c.IsSet("zsh"):
				return utils.CmdHandler(utils.GenerateShellCompletion("zsh"))
			}
			return cli.ShowSubcommandHelp(c)
		},
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "bash",
				Usage: "bash completion",
			},
			cli.BoolFlag{
				Name:  "zsh",
				Usage: "zsh completion",
			},
		},
	},
}
