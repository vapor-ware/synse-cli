package hosts

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vapor-ware/synse-cli/pkg/cmd/hosts/active"
	"github.com/vapor-ware/synse-cli/pkg/cmd/hosts/add"
	"github.com/vapor-ware/synse-cli/pkg/cmd/hosts/change"
	del "github.com/vapor-ware/synse-cli/pkg/cmd/hosts/delete"
	"github.com/vapor-ware/synse-cli/pkg/cmd/hosts/list"
)

// New returns a new instance of the 'hosts' command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hosts",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// todo
			fmt.Println("< hosts")
		},
	}

	// Add sub-commands
	cmd.AddCommand(
		active.New(),
		add.New(),
		change.New(),
		del.New(),
		list.New(),
	)

	return cmd
}
