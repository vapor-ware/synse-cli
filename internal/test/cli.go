package test

import (
	"bytes"
	"fmt"

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

type FakeApp struct {
	OutBuffer *bytes.Buffer
	ErrBuffer *bytes.Buffer
	*cli.App
}

// NewFakeApp creates a new "fake" application used for testing.
func NewFakeApp() *FakeApp {
	outBuffer := new(bytes.Buffer)
	errBuffer := new(bytes.Buffer)

	cliApp := &cli.App{
		// Name of the test application
		Name: "test app",

		// Write out to the `outBuffer` - this way we can later
		// read out from it to validate the output
		Writer: outBuffer,

		// Write errors out to the `errBuffer` - this way we can
		// later read out from it to validate the output
		ErrWriter: errBuffer,
		ExitErrHandler: func(context *cli.Context, err error) {
			fmt.Fprintln(errBuffer, err)
		},
	}

	app := &FakeApp{
		OutBuffer: outBuffer,
		ErrBuffer: errBuffer,
		App:       cliApp,
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
