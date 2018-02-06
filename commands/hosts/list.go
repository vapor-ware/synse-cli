package hosts

import (
	"sort"

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

// hostInfo represents a single host to be listed.
type hostInfo struct {
	Active  string
	Name    string
	Address string
}

// ToRow converts a hostInfo to a table row.
func (host *hostInfo) ToRow() []string {
	return []string{
		host.Active,
		host.Name,
		host.Address,
	}
}

// byHostName is the container type for sorting
type byHostName []*hostInfo

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
	var hosts []*hostInfo
	for _, host := range config.Config.Hosts {
		isActive := ""
		if config.Config.ActiveHost != nil && *host == *config.Config.ActiveHost {
			isActive = "*"
		}
		hosts = append(hosts, &hostInfo{
			Active:  isActive,
			Name:    host.Name,
			Address: host.Address,
		})
	}

	// Sort by host name
	sort.Sort(byHostName(hosts))

	var data [][]string
	for _, host := range hosts {
		data = append(data, host.ToRow())
	}

	utils.MinimalTableOutput(data)
	return nil
}
