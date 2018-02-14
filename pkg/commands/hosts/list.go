package hosts

import (
	"sort"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/config"
	"github.com/vapor-ware/synse-cli/pkg/formatters/hosts"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

// hostsListCommand is the CLI sub-command for listing all configured hosts.
var hostsListCommand = cli.Command{
	Name:  "list",
	Usage: "List the configured Synse Server hosts",
	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdList(c))
	},
}

// byHostName is the container type for sorting
type byHostName []*config.HostConfig

func (h byHostName) Len() int {
	return len(h)
}

func (h byHostName) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h byHostName) Less(i, j int) bool {
	return h[i].Name < h[j].Name
}

// cmdList is the action for hostsListCommand. It prints out all of the configured
// hosts' names and addresses.
func cmdList(c *cli.Context) error {
	var configuredHosts []*config.HostConfig
	for _, c := range config.Config.Hosts {
		configuredHosts = append(configuredHosts, c)
	}

	// Sort by host name
	sort.Sort(byHostName(configuredHosts))

	// Format output
	formatter := hosts.NewListFormatter(c.App.Writer)
	err := formatter.Add(configuredHosts)
	if err != nil {
		return err
	}
	return formatter.Write()
}
