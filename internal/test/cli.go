package test

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/config"
	"github.com/vapor-ware/synse-cli/pkg/flags"
)

// Setup can be run at the beginning of tests to perform setup actions.
// The current setup actions are:
//   - Clearing the CLI config
func Setup() {
	config.Config = config.CliConfig{
		Debug:      false,
		ActiveHost: nil,
		Hosts:      make(map[string]*config.HostConfig),
	}
}

// NewFakeApp creates a new "fake" application used for testing.
func NewFakeApp() *cli.App {
	app := &cli.App{
		Name: "test app",
	}

	app.Flags = []cli.Flag{
		flags.DebugFlag,
		flags.ConfigFlag,
		flags.FormatFlag,
	}

	// prevent the fake app from calling os.Exit() on failure
	cli.OsExiter = func(code int) {}
	return app
}
