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
	"github.com/vapor-ware/synse-cli/pkg/flags"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
)

const (
	appName    = "synse"
	appUsage   = "Command line tool for interacting with Synse components"
	appVersion = "0.1.0"
)

// Create a new instance of the CLI application, configure it, and run it.
func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Usage = appUsage
	app.Version = appVersion

	app.Flags = flags.GlobalFlags
	app.Commands = commands.Commands
	app.EnableBashCompletion = true

	// Before running, load and construct the CLI configuration
	app.Before = func(c *cli.Context) error {
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

	// After running, persist the configuration
	app.After = func(c *cli.Context) error {
		return config.Persist()
	}

	app.Action = func(c *cli.Context) error {
		if c.IsSet("config") {
			return formatters.AsYAML(config.Config, c.App.Writer)
		}

		return cli.ShowAppHelp(c)
	}

	// Run the CLI
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
