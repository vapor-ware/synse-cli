package hosts

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/config"
	"github.com/vapor-ware/synse-cli/utils"
)

// hostsListCommand is the CLI sub-command for listing all configured hosts.
var hostsListCommand = cli.Command{
	Name:   "list",
	Usage:  "List the configured Synse Server hosts",
	Action: cmdList,
}

// cmdList is the action for hostsListCommand. It prints out all of the configured
// hosts' names and addresses.
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
