package plugin

import (
	"io"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/formatters"
	"github.com/vapor-ware/synse-cli/scheme"
	"github.com/vapor-ware/synse-cli/utils"
	"github.com/vapor-ware/synse-server-grpc/go"
	"golang.org/x/net/context"
)

// pluginReadCommand is a CLI sub-command for getting a reading from a plugin.
var pluginReadCommand = cli.Command{
	Name:  "read",
	Usage: "Get a reading from a plugin",
	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdRead(c))
	},
}

// cmdRead is the action for pluginReadCommand. It prints out a reading that was
// retrieved from the specified plugin.
func cmdRead(c *cli.Context) error { // nolint: gocyclo
	err := utils.RequiresArgsExact(3, c)
	if err != nil {
		return err
	}

	rack := c.Args().Get(0)
	board := c.Args().Get(1)
	device := c.Args().Get(2)

	pluginClient, err := makeGrpcClient(c)
	if err != nil {
		return err
	}

	stream, err := pluginClient.Read(context.Background(), &synse.ReadRequest{
		Device: device,
		Board:  board,
		Rack:   rack,
	})
	if err != nil {
		return err
	}

	formatter := formatters.NewReadFormatter(c.App.Writer)
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		r := &scheme.Read{
			Type: resp.Type,
			Data: map[string]scheme.ReadData{
				resp.Type: {
					Value:     resp.Value,
					Timestamp: resp.Timestamp,
				},
			},
		}
		err = formatter.Add(r)
		if err != nil {
			return err
		}
	}
	return formatter.Write()
}
