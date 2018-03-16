package hosts

import (
	"testing"

	"github.com/gotestyourself/gotestyourself/assert"
	"github.com/gotestyourself/gotestyourself/golden"

	"github.com/vapor-ware/synse-cli/internal/test"
	"github.com/vapor-ware/synse-cli/pkg/config"
)

// TestChangeCommandError tests the 'change' command with no arguments
// given.
func TestChangeCommandError(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, HostsCommand)

	// expects exactly one arg, here passing none
	err := app.Run([]string{
		app.Name,
		HostsCommand.Name,
		hostsChangeCommand.Name,
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "change.error.no_args.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestChangeCommandError2 tests the 'change' command specifying a host
// that doesn't exist in the config.
func TestChangeCommandError2(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, HostsCommand)

	// specifying a host that does not exist - should cause failure
	err := app.Run([]string{
		app.Name,
		HostsCommand.Name,
		hostsChangeCommand.Name,
		"missing-host",
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "change.error.invalid_args.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestChangeCommandError3 tests the 'change' command specifying extra arguments.
func TestChangeCommandError3(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, HostsCommand)

	// expects exactly one arg, here passing in multiple
	err := app.Run([]string{
		app.Name,
		HostsCommand.Name,
		hostsChangeCommand.Name,
		"arg1", "arg2",
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "change.error.extra_args.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestChangeCommandSuccess tests the 'change' command, successfully changing
// the active host in the config.
func TestChangeCommandSuccess(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, HostsCommand)

	config.Config.Hosts["test-host"] = &config.HostConfig{
		Name:    "test-host",
		Address: "test-address",
	}

	// before running, check that the active host is nil
	if config.Config.ActiveHost != nil {
		t.Error("expected active host to be nil at test start")
	}

	err := app.Run([]string{
		app.Name,
		HostsCommand.Name,
		hostsChangeCommand.Name,
		"test-host",
	})

	assert.Assert(t, golden.String(app.OutBuffer.String(), "change.success.golden"))
	test.ExpectNoError(t, err)

	// after running, check that the active host is now set
	if config.Config.ActiveHost != config.Config.Hosts["test-host"] {
		t.Error("active host was not set to the 'test-host'")
	}
}
