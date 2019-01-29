package update

import (
	"github.com/jpillora/overseer"
	"github.com/jpillora/overseer/fetcher"
	"github.com/spf13/cobra"
)

const (
	ghOrgName  = "vapor-ware"
	ghRepoName = "synse-cli"
)

var (
	debug = false
)

// New returns a new instance of the 'update' command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Install the newest release of synse",
		Long: `The update command checks for any new releases of the Synse CLI. If this
version is out-of-date, the new version will be installed and replace
the current binary.

Updating may require administrative privileges if the binary is in a
restricted location (such as /bin).

synse may be updated manually by download a release from its GitHub page:
https://github.com/vapor-ware/synse-cli/releases`,

		Run: func(cmd *cobra.Command, args []string) {
			overseer.Run(overseer.Config{
				Required:  true,
				Program:   func(_ overseer.State) {},
				NoRestart: true,
				Debug:     debug,
				NoWarn:    false,
				Fetcher: &fetcher.Github{
					User: ghOrgName,
					Repo: ghRepoName,
				},
			})
		},
	}

	// Add flag options to the command.
	cmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable debug output for the updater")

	return cmd
}
