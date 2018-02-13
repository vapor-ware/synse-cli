package plugin

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/formatters"
	"github.com/vapor-ware/synse-cli/scheme"
	"github.com/vapor-ware/synse-cli/utils"
	"github.com/vapor-ware/synse-server-grpc/go"
	"golang.org/x/net/context"
)

// pluginWriteCommand is a CLI sub-command for writing to a plugin
var pluginWriteCommand = cli.Command{
	Name:  "write",
	Usage: "Write data directly to a plugin",
	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdWrite(c))
	},
}

// cmdWrite is the action for pluginWriteCommand. It writes directly to
// the specified plugin.
func cmdWrite(c *cli.Context) error { // nolint: gocyclo
	err := utils.RequiresArgsInRange(4, 5, c)
	if err != nil {
		return err
	}

	rack := c.Args().Get(0)
	board := c.Args().Get(1)
	device := c.Args().Get(2)
	action := c.Args().Get(3)
	raw := c.Args().Get(4)

	wd := &synse.WriteData{
		Action: action,
	}
	if raw != "" {
		wd.Raw = [][]byte{[]byte(raw)}
	}

	pluginClient, err := makeGrpcClient(c)
	if err != nil {
		return err
	}

	transactions, err := pluginClient.Write(context.Background(), &synse.WriteRequest{
		Device: device,
		Board:  board,
		Rack:   rack,
		Data:   []*synse.WriteData{wd},
	})
	if err != nil {
		return err
	}

	t := make([]scheme.WriteTransaction, len(transactions.Transactions))
	for id, ctx := range transactions.Transactions {
		var raw []string
		for _, r := range ctx.Raw {
			raw = append(raw, string(r))
		}

		t = append(t, scheme.WriteTransaction{
			Transaction: id,
			Context: scheme.WriteContext{
				Action: ctx.Action,
				Raw:    raw,
			},
		})
	}

	formatter := formatters.NewWriteFormatter(c.App.Writer)
	err = formatter.Add(t)
	if err != nil {
		return err
	}
	return formatter.Write()
}
