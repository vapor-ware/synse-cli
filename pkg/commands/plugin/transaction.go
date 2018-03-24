package plugin

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/client"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
	"github.com/vapor-ware/synse-cli/pkg/scheme"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

const (
	// transactionCmdName is the name for the 'transaction' command.
	transactionCmdName = "transaction"

	// transactionCmdUsage is the usage text for the 'transaction' command.
	transactionCmdUsage = "Get transaction info from a plugin"

	// transactionCmdArgsUsage is the argument usage for the 'transaction' command.
	transactionCmdArgsUsage = "TRANSACTION_ID"

	// transactionCmdDesc is the description for the 'transaction' command.
	transactionCmdDesc = `The transaction command gets the state and status of a
  write via the Synse gRPC API.

  Writes for Synse Plugins are asynchronous, so to verify
  that a write has completed successfully, you must check
  the state of the transaction afterwards.

  The possible transaction states and statuses are:

  STATUS        STATE
  ----------    ----------
  unknown       ok
  pending       error
  writing
  done

Example:
  synse plugin transaction bah8volrogrg01o4sjtg

Formatting:
  The 'plugin transaction' command supports the following formatting
  options (via the CLI global --format flag):
    - pretty (default)
		- yaml
    - json`
)

// pluginTransactionCommand is a CLI sub-command for getting transaction info from a plugin.
var pluginTransactionCommand = cli.Command{
	Name:        transactionCmdName,
	Usage:       transactionCmdUsage,
	Description: transactionCmdDesc,
	ArgsUsage:   transactionCmdArgsUsage,

	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdTransaction(c))
	},
}

// cmdTransaction is the action for pluginTransactionCommand. It prints out transaction info
// retrieved from the specified plugin.
func cmdTransaction(c *cli.Context) error {
	err := utils.RequiresArgsExact(1, c)
	if err != nil {
		return err
	}

	tid := c.Args().Get(0)

	status, err := client.Grpc.Transaction(c, tid)
	if err != nil {
		return err
	}
	s := &scheme.Transaction{
		ID:      tid,
		State:   status.State.String(),
		Status:  status.Status.String(),
		Created: status.Created,
		Updated: status.Updated,
	}

	formatter := formatters.NewTransactionFormatter(c, status)
	err = formatter.Add(s)
	if err != nil {
		return err
	}
	return formatter.Write()
}
