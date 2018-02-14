package hosts

import (
	"testing"

	"github.com/vapor-ware/synse-cli/internal/test"
	"github.com/vapor-ware/synse-cli/pkg/config"
)

func TestAddCommandError(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, hostsAddCommand)

	err := app.Run([]string{app.Name, hostsAddCommand.Name})

	test.ExpectExitCoderError(t, err)
}

func TestAddCommandError2(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, hostsAddCommand)

	// Before adding the "name / address" host, we want it to already exist.
	config.Config.Hosts["name"] = &config.HostConfig{
		Name:    "name",
		Address: "address",
	}

	err := app.Run([]string{app.Name, hostsAddCommand.Name, "name", "address"})

	test.ExpectExitCoderError(t, err)
}

func TestAddCommandSuccess(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, hostsAddCommand)

	err := app.Run([]string{app.Name, hostsAddCommand.Name, "name", "address"})

	test.ExpectNoError(t, err)
}
