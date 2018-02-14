package plugin

import (
	"io"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
	"github.com/vapor-ware/synse-cli/pkg/utils"
	"github.com/vapor-ware/synse-server-grpc/go"
	"golang.org/x/net/context"
)

// pluginMetainfoCommand is a CLI sub-command for getting metainfo from a plugin.
var pluginMetainfoCommand = cli.Command{
	Name:  "meta",
	Usage: "Get the metainformation from a plugin",

	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdMeta(c))
	},
}

// cmdMeta is the action for pluginMetainfoCommand. It prints out the meta-information
// provided by the specified plugin.
func cmdMeta(c *cli.Context) error {
	pluginClient, err := makeGrpcClient(c)
	if err != nil {
		return err
	}

	stream, err := pluginClient.Metainfo(
		context.Background(),
		&synse.MetainfoRequest{},
	)
	if err != nil {
		return err
	}

	formatter := formatters.NewMetaFormatter(c.App.Writer)
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		err = formatter.Add(resp)
		if err != nil {
			return err
		}
	}
	return formatter.Write()
}
