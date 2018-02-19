package plugin

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/client"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
	"github.com/vapor-ware/synse-cli/pkg/scheme"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

// pluginTransactionCommand is a CLI sub-command for getting transaction info from a plugin.
var pluginTransactionCommand = cli.Command{
	Name:  "transaction",
	Usage: "Get transaction info from a plugin",

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

	formatter := formatters.NewTransactionFormatter(c)
	err = formatter.Add(s)
	if err != nil {
		return err
	}
	return formatter.Write()
}
