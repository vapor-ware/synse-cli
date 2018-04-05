package commands

import (
	"github.com/urfave/cli"

	"github.com/jpillora/overseer/fetcher"
	"github.com/timfallmk/overseer"

	"github.com/vapor-ware/synse-cli/pkg/utils"
)

const (
	// updateCmdName is the name for the 'update' command.
	updateCmdName = "update"

	//updateoCmdUsage is the usage text for the 'update' command.
	updateCmdUsage = "Check for and download new releases"

	// updateCmdArgsUsage is the argument usage for the 'update' command.
	updateCmdArgsUsage = ""

	// updateCmdDesc is the description for the 'update' command.
	updateCmdDesc = `The update command checks the GitHub repository associated with this
	application for releases newer than the current releases. If found, it will stream the
	newer release in place, replacing the current binary. This may require administrative privileges if the binary is in a restricted location (such and /bin).

	You may also download any releases manually at:
	https://github.com/vapor-ware/synse-cli/releases`
)

const (
	repoUsername = "vapor-ware"
	repoName     = "synse-cli"
)

// updateCommand is the command specification for running updates
var updateCommand = cli.Command{
	Name:        updateCmdName,
	Usage:       updateCmdUsage,
	Description: updateCmdDesc,
	ArgsUsage:   utils.NoArgs,

	// Since we don't produce an error, skip using utils.CommandHandler
	Action: func(c *cli.Context) error {
		cmdUpdate()
		return nil
	},
}

func cmdUpdate() {
	overseer.Run(overseer.Config{
		Required:  true,
		Program:   func(_ overseer.State) {},
		NoRestart: true,
		Debug:     true,
		NoWarn:    false,
		Fetcher: &fetcher.Github{
			User: repoUsername,
			Repo: repoName,
		},
	})
}
