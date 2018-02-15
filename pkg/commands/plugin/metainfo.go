package plugin

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/client"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
	"github.com/vapor-ware/synse-cli/pkg/utils"
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

	resp, err := client.Grpc.Metainfo(c)
	if err != nil {
		return err
	}

	formatter := formatters.NewMetaFormatter(c.App.Writer)
	for _, meta := range resp {
		err = formatter.Add(meta)
		if err != nil {
			return err
		}
	}
	return formatter.Write()
}
