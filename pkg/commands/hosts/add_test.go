package hosts

import (
	"testing"

	"github.com/gotestyourself/gotestyourself/assert"
	"github.com/gotestyourself/gotestyourself/golden"

	"github.com/vapor-ware/synse-cli/internal/test"
	"github.com/vapor-ware/synse-cli/pkg/config"
)

// TestAddCommandError tests the 'add' command with no arguments given.
func TestAddCommandError(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, hostsAddCommand)

	err := app.Run([]string{app.Name, hostsAddCommand.Name})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "add.error.no_args.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestAddCommandError2 tests the 'add' command when the given arguments
// match a host that already exists
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

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "add.error.duplicate_args.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestAddCommandError3 tests the 'add' command, supplying more than two
// arguments.
func TestAddCommandError3(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, hostsAddCommand)

	// Before adding the "name / address" host, we want it to already exist.
	config.Config.Hosts["name"] = &config.HostConfig{
		Name:    "name",
		Address: "address",
	}

	err := app.Run([]string{app.Name, hostsAddCommand.Name, "name", "address", "extra"})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "add.error.extra_args.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestAddCommandSuccess tests the 'add' command, successfully adding a host
// to the config.
func TestAddCommandSuccess(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, hostsAddCommand)

	err := app.Run([]string{app.Name, hostsAddCommand.Name, "name", "address"})

	assert.Assert(t, golden.String(app.OutBuffer.String(), "add.success.golden"))
	test.ExpectNoError(t, err)
}
