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
	"github.com/vapor-ware/synse-cli/pkg/version"
)

const (
	appName  = "synse"
	appUsage = "Command line tool for interacting with Synse components"
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
		formatter := formatters.NewConfigFormatter(c, config.Config)
		return formatter.Write()
	}
	return cli.ShowAppHelp(c)
}

// Create a new instance of the CLI application and run it.
func main() {
	versionInfo := version.Get()

	app := cli.NewApp()

	app.Name = appName
	app.Usage = appUsage
	app.Version = versionInfo.VersionString

	app.Before = appBefore
	app.After = appAfter
	app.Action = appAction

	app.EnableBashCompletion = true
	app.Commands = commands.Commands

	app.Flags = []cli.Flag{
		flags.DebugFlag,  // --debug, -d flag to enable debug logging output
		flags.ConfigFlag, // --config flag to display the YAML configuration for the CLI
		flags.FormatFlag, // --format flag to specify the output format for a command
	}

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Println("synse version:")
		fmt.Printf("\tVersion:    %s\n", versionInfo.VersionString)
		fmt.Printf("\tGo Version: %s\n", versionInfo.GoVersion)
		fmt.Printf("\tGit Commit: %s\n", versionInfo.GitCommit)
		fmt.Printf("\tGit Tag:    %s\n", versionInfo.GitTag)
		fmt.Printf("\tBuild Date: %s\n", versionInfo.BuildDate)
		fmt.Printf("\tOS/Arch:    %s/%s\n", versionInfo.OS, versionInfo.Arch)
	}

	cli.AppHelpTemplate = `
{{.Usage}}

Usage:
  {{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} COMMAND [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}
{{if .Commands}}
Commands:
{{range .Commands}}{{if not .HideHelp}}  {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
Global Options:
  {{range .VisibleFlags}}{{.}}
  {{end}}{{end}}
Use 'synse COMMAND --help' for more information on a command.
`

	cli.CommandHelpTemplate = `
{{.Usage}}

Usage:
  {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}}{{if .VisibleFlags}} [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}{{if .Category}}

Category:
  {{.Category}}{{end}}{{if .Description}}

Description:
  {{.Description}}{{end}}{{if .VisibleFlags}}

Options:
  {{range .VisibleFlags}}{{.}}
  {{end}}{{end}}
`

	cli.SubcommandHelpTemplate = `
{{.Usage}}{{if .Description}}

Description:
  {{.Description}}{{end}}

Usage:
  {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} command{{if .VisibleFlags}} [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}

Commands:{{range .VisibleCategories}}{{if .Name}}
  {{.Name}}:{{end}}{{range .VisibleCommands}}
  {{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}
{{end}}{{if .VisibleFlags}}
Options:
  {{range .VisibleFlags}}{{.}}
  {{end}}{{end}}
`

	// Run the CLI
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
