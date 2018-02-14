package pkg

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/config"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
)

// NewSynseApp creates a new instance of a cli.App that is configured with the
// commands, flags, and actions of the Synse CLI. The application name, version,
// and usage information is not set here. Instead it is expected to be set by the
// consumer of this function (e.g. cmd/synse/synse.go)
func NewSynseApp() *cli.App {
	app := cli.NewApp()

	app.Commands = Commands
	app.EnableBashCompletion = true

	app.Before = appBefore
	app.After = appAfter
	app.Action = appAction

	app.Flags = []cli.Flag{
		// --debug, -d flag to enable debug logging output
		cli.BoolFlag{
			Name:  "debug, d",
			Usage: "enable debug mode",
		},
		// --config flag to display the YAML configuration for the CLI
		cli.BoolFlag{
			Name:  "config",
			Usage: "display the current CLI configuration",
		},
	}

	return app
}

// appBefore defines the action to take before the command is processed. Currently,
// this reads in any existing CLI configuration and sets the logging level based on
// the configuration it finds, if any.
func appBefore(c *cli.Context) error {
	// Allow debugging of the config loading process
	if c.Bool("debug") {
		log.SetLevel(log.DebugLevel)
	}

	// Construct the config for this session.
	err := config.ConstructConfig(c)
	if err != nil {
		fmt.Println(err)
	}

	if config.Config.Debug {
		log.SetLevel(log.DebugLevel)
	}
	return nil
}

// appAfter defines the action to take after a command has completed. Currently,
// this persists the configuration state for future runs of the CLI application.
func appAfter(c *cli.Context) error {
	return config.Persist()
}

// appAction defines the action to take if the application is called
// with no commands.
//
// If the --config flag is set, it will display the application configuration,
// otherwise it will display the usage information.
func appAction(c *cli.Context) error {
	if c.IsSet("config") {
		return formatters.AsYAML(config.Config, c.App.Writer)
	}
	return cli.ShowAppHelp(c)
}
