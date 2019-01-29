package completion

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// New returns a new instance of the 'completion' command. It requires the
// root command to be passed through in order to properly set up completion.
func New(rootCmd *cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "completion",
		Short: "Generate bash completion scripts",
		Long: `Generate bash completion scripts.

To load bash completion for the current session, run:

  . <(synse completion)

To configure your bash shell to load synse completion for all
new sessions, add the above to your bashrc, e.g.

  echo ". <(synse completion)" >> ~/.bashrc
`,

		Run: func(cmd *cobra.Command, args []string) {
			if err := rootCmd.GenBashCompletion(os.Stdout); err != nil {
				// TODO: error out in a consistent way.
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	return cmd
}
