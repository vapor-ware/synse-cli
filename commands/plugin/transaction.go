package plugin

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/formatters/plugin"
	"github.com/vapor-ware/synse-cli/utils"
	"github.com/vapor-ware/synse-server-grpc/go"
	"golang.org/x/net/context"
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

	pluginClient, err := makeGrpcClient(c)
	if err != nil {
		return err
	}

	formatter := plugin.NewTransactionFormatter(c.App.Writer)

	status, err := pluginClient.TransactionCheck(context.Background(), &synse.TransactionId{
		Id: tid,
	})
	if err != nil {
		return err
	}
	err = formatter.Add(status)
	if err != nil {
		return err
	}
	return formatter.Write()
}
