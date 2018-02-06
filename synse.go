// Synse provides a cli to vapor.io infrastructure. Specifically it allows access
// to the Synse (http://www.vapor.io/synse/) REST API for running commands
// against connected infrastructure and devices.
//
// For usage information please see the help text or the README in this
// repository.
package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/commands"
	"github.com/vapor-ware/synse-cli/config"
	"github.com/vapor-ware/synse-cli/flags"
)

// Main creates a new instance of cli.app (using https://github.com/urfave/cli)
// and sets the default configuration.
func main() {
	app := cli.NewApp()
	app.Name = "synse"
	app.Usage = "Synse CLI"
	app.Version = "0.1.0"
	app.Authors = []cli.Author{
		{Name: "Tim Fall", Email: "tim@vapor.io"},
		{Name: "Thomas Rampelberg", Email: "thomasr@vapor.io"},
	}

	app.Flags = flags.GlobalFlags
	app.Commands = commands.Commands
	app.EnableBashCompletion = true

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

	app.After = func(c *cli.Context) error {
		return config.Persist()
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
