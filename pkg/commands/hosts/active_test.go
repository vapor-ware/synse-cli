package hosts

import (
	"testing"

	"github.com/vapor-ware/synse-cli/internal/test"
	"github.com/vapor-ware/synse-cli/pkg/config"
)

func TestActiveCommandError(t *testing.T) {

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, hostsActiveCommand)

	// Set the active host to nil
	config.Config.ActiveHost = nil

	err := app.Run([]string{app.Name, hostsActiveCommand.Name})

	test.ExpectExitCoderError(t, err)
}

func TestActiveCommandSuccess(t *testing.T) {

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, hostsActiveCommand)

	// Set the active host to a HostConfig
	config.Config.ActiveHost = &config.HostConfig{
		Name:    "test-host",
		Address: "test-address",
	}

	err := app.Run([]string{app.Name, hostsActiveCommand.Name})

	test.ExpectNoError(t, err)
}
