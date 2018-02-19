// `synse` is a CLI for interacting with Synse components including Synse Server,
// via its HTTP API, and Synse plugins.
//
// For usage information please see the help text or the README in this
// repository.
package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/commands"
	"github.com/vapor-ware/synse-cli/pkg/config"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
)

const (
	appName    = "synse"
	appUsage   = "Command line tool for interacting with Synse components"
	appVersion = "0.1.0"
)

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

// Create a new instance of the CLI application and run it.
func main() {
	app := cli.NewApp()

	app.Name = appName
	app.Usage = appUsage
	app.Version = appVersion

	app.Before = appBefore
	app.After = appAfter
	app.Action = appAction

	app.EnableBashCompletion = true
	app.Commands = commands.Commands

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

	// Run the CLI
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
