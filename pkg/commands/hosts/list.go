package hosts

import (
	"sort"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/config"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

const (
	// listCmdName is the name for the 'list' command.
	listCmdName = "list"

	// listCmdUsage is the usage text for the 'list' command.
	listCmdUsage = "List all configured Synse Server hosts"

	// activeCmdDesc is the description for the 'list' command.
	listCmdDesc = `The list command shows all of the Synse Server instances
  that are currently configured with the CLI. The current
  active host will be denoted with a '*', if there is an
  active host.

Example:
  synse hosts list

Formatting:
  The 'hosts list' command supports the following formatting
  options (via the CLI global --format flag):
    - pretty (default)
		- yaml
    - json`
)

// hostsListCommand is the CLI sub-command for listing all configured hosts.
var hostsListCommand = cli.Command{
	Name:        listCmdName,
	Usage:       listCmdUsage,
	Description: listCmdDesc,
	ArgsUsage:   utils.NoArgs,

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
	configuredHosts := HostList()

	// Sort by host name
	sort.Sort(byHostName(configuredHosts))

	// Format output
	formatter := formatters.NewListFormatter(c, configuredHosts)
	err := formatter.Add(configuredHosts)
	if err != nil {
		return err
	}
	return formatter.Write()
}

// HostList creates an unsorted list of hosts present in the configuration and
// returns that list.
func HostList() []*config.HostConfig {
	var configuredHosts []*config.HostConfig
	for _, host := range config.Config.Hosts {
		configuredHosts = append(configuredHosts, host)
	}
	return configuredHosts
}
