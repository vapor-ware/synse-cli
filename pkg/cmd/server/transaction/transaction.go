package transaction

import (
	"fmt"

	"github.com/spf13/cobra"
)

// New returns a new instance of the 'server transaction' command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transaction",
		Short: "Check transaction state/status",
		Long:  `Check the state and status of a write transaction.

Write in Synse Server are asynchronous. To verify the outcome
of a write, this command can be used to check the transaction
afterwards. Below are the possible states and statuses.

States      Statuses
 - ok        - unknown
 - error     - pending
             - writing
             - done

A transaction is considered complete either when it has reached
the 'done' status or is in the error state.

For more information, see:
https://vapor-ware.github.io/synse-server/#transaction`,

		Run: func(cmd *cobra.Command, args []string) {
			// todo
			fmt.Println("< server transaction")
		},
	}

	return cmd
}
