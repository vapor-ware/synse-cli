// Vesh provides a cli to vapor.io infrastructure. Specifically it allows access
// to the openDCRE (http://www.vapor.io/opendcre/) REST API for running commands
// against connected infrastructure and devices.
//
// For usage information please see the help text or the README in this
// repository.
package main

import (
	"fmt"
	"os"

	"github.com/vapor-ware/vesh/client"
	"github.com/vapor-ware/vesh/commands"
	"github.com/vapor-ware/vesh/utils"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

// Main creates a new instance of cli.app (using https://github.com/urfave/cli)
// and sets the default configuration.
func main() {
	app := cli.NewApp()
	app.Name = "vesh"
	app.Usage = "Vapor Edge Shell"
	app.Version = "0.0.1"
	app.Authors = []cli.Author{{Name: "Tim Fall", Email: "tim@vapor.io"},
		{Name: "Thomas Rampelberg", Email: "thomasr@vapor.io"}}

	app.Commands = commands.Commands
	//app.CommandNotFound = commands.CommandNotFound
	app.EnableBashCompletion = true

	app.Before = func(c *cli.Context) error {
		// Allow debugging of the config loading process
		if c.Bool("debug") {
			log.SetLevel(log.DebugLevel)
		}

		err := utils.ConstructConfig(c)
		if err != nil {
			fmt.Println(err)
		}

		if utils.Config.Debug {
			log.SetLevel(log.DebugLevel)
		}

		client.Config(utils.Config.VaporHost)

		return nil
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, d",
			Usage: "Enable debug mode",
		},
		cli.StringFlag{
			EnvVar: "VESH_CONFIG_FILE",
			Name:   "config, c",
			Usage:  "Path to config `file`",
		},
		cli.StringFlag{
			EnvVar: "VAPOR_HOST",
			Name:   "vapor-host",
			Usage:  "Address of `Vapor Host`",
		},
	}

	app.Run(os.Args)
}
