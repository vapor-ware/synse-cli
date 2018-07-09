package plugin

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/client"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
	"github.com/vapor-ware/synse-cli/pkg/scheme"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

const (
	// healthCmdName is the name for the 'health' command.
	healthCmdName = "health"

	// healthCmdUsage is the usage text for the 'health' command.
	healthCmdUsage = "Get the device health for a plugin"

	// healthCmdDesc is the description for the 'health' command.
	healthCmdDesc = `The health command gets information on the health status
  summarizing the plugin's health, as well as a list of the HealthChecks
  which make up that status. The health information returned here
  is similar to that of a 'synse server plugins health' command with
  a plugin tag specified.

  The 'plugin health' command takes no arguments.

Example:
  synse plugin health

Formatting:
  The 'plugin health' command supports the following formatting
  options (via the CLI global --format flag):
    - pretty (default)`
)

// pluginHealthCommand is a CLI sub-command for getting health info from a plugin.
var pluginHealthCommand = cli.Command{
	Name:        healthCmdName,
	Usage:       healthCmdUsage,
	Description: healthCmdDesc,
	ArgsUsage:   utils.NoArgs,

	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdHealth(c))
	},
}

// cmdHealth is the action for pluginHealthCommand. It prints out
// the health information provided by the specified plugin.
func cmdHealth(c *cli.Context) error {
	health, err := client.Grpc.Health(c)
	if err != nil {
		return err
	}

	formatter := formatters.NewPluginHealthFormatter(c)

	var checks []scheme.CheckData
	for _, check := range health.Checks {
		checks = append(checks, scheme.CheckData{
			Name:      check.Name,
			Status:    check.Status.String(),
			Message:   check.Message,
			Timestamp: check.Timestamp,
			Type:      check.Type,
		})
	}

	err = formatter.Add(&scheme.HealthData{
		Timestamp: health.Timestamp,
		Status:    health.Status.String(),
		Checks:    checks,
	})
	if err != nil {
		return err
	}

	return formatter.Write()
}
