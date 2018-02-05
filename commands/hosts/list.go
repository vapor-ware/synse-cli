package hosts

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/utils"
	"github.com/vapor-ware/synse-cli/config"
)

var hostsListCommand = cli.Command{
	Name: "list",
	Usage: "list the configured Synse Server hosts",
	Action: cmdList,

}

func cmdList(c *cli.Context) error {
	var data [][]string
	for _, host := range config.Config.Hosts {
		isActive := ""
		if config.Config.ActiveHost != nil && *host == *config.Config.ActiveHost {
			isActive = "*"
		}
		data = append(data, []string{
			isActive,
			host.Name,
			host.Address,
		})
	}

	// FIXME (etd) - data should be sorted
	utils.MinimalTableOutput(data)
	return nil
}
