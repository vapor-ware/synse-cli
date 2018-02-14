package commands

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/flags"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

// completionCommand is the CLI command for generating shell completion scripts.
var completionCommand = cli.Command{
	Name:  "completion",
	Usage: "Generate shell completion scripts for bash or zsh",
	Flags: []cli.Flag{
		flags.BashFlag,
		flags.ZshFlag,
	},
	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdCompletion(c))
	},
}

// cmdCompletion is the action for the completionCommand.
func cmdCompletion(c *cli.Context) error {
	switch {
	case c.IsSet("bash") && c.IsSet("zsh"):
		return fmt.Errorf("cannot create completion scripts for both bash and zsh")
	case c.IsSet("bash"):
		return utils.GenerateShellCompletion(c, "bash")
	case c.IsSet("zsh"):
		return utils.GenerateShellCompletion(c, "zsh")
	default:
		return cli.ShowSubcommandHelp(c)
	}
}
