// `synse` is a CLI for interacting with Synse components including Synse Server,
// via its HTTP API, and Synse plugins.
//
// For usage information please see the help text or the README in this
// repository.
package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/vapor-ware/synse-cli/pkg"
)

const (
	appName    = "synse"
	appUsage   = "Command line tool for interacting with Synse components"
	appVersion = "0.1.0"
)

// Create a new instance of the CLI application and run it.
func main() {
	app := pkg.NewSynseApp()

	app.Name = appName
	app.Usage = appUsage
	app.Version = appVersion

	// Run the CLI
	err := app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}
