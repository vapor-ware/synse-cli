package hosts

import (
	"testing"

	"github.com/gotestyourself/gotestyourself/assert"
	"github.com/gotestyourself/gotestyourself/golden"

	"github.com/vapor-ware/synse-cli/internal/test"
	"github.com/vapor-ware/synse-cli/pkg/config"
)

// TestDeleteCommandError tests the 'delete' command when no arguments are given.
func TestDeleteCommandError(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, hostsDeleteCommand)

	// expects exactly one arg, but none are given
	err := app.Run([]string{app.Name, hostsDeleteCommand.Name})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "delete.error.no_args.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestDeleteCommandError2 tests the 'delete' command when extra arguments are given.
func TestDeleteCommandError2(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, hostsDeleteCommand)

	// expects exactly one arg, but multiple are given
	err := app.Run([]string{app.Name, hostsDeleteCommand.Name, "arg1", "arg2"})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "delete.error.extra_args.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestDeleteCommandSuccess tests the 'delete' command for a host that does
// not exist in the configuration.
func TestDeleteCommandSuccess(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, hostsDeleteCommand)

	// we should not fail if we try to delete a host that is not in
	// the configuration
	err := app.Run([]string{app.Name, hostsDeleteCommand.Name, "missing-host"})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "delete.success.golden"))
	test.ExpectNoError(t, err)
}

// TestDeleteCommandSuccess2 tests the 'delete' command for a host that does
// exist in the configuration.
func TestDeleteCommandSuccess2(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, hostsDeleteCommand)

	config.Config.Hosts["test-host"] = &config.HostConfig{
		Name:    "test-host",
		Address: "test-address",
	}

	// we should not fail if we try to delete a host that is in the
	// configuration
	err := app.Run([]string{app.Name, hostsDeleteCommand.Name, "test-host"})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "delete.success.golden"))
	test.ExpectNoError(t, err)
}

// TestDeleteCommandSuccess3 tests the 'delete' command for a host that does
// exist in the configuration and is the active host.
func TestDeleteCommandSuccess3(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, hostsDeleteCommand)

	host := &config.HostConfig{
		Name:    "test-host",
		Address: "test-address",
	}
	config.Config.Hosts["test-host"] = host
	config.Config.ActiveHost = host

	// we should not fail if we try to delete a host that is in the
	// configuration and is also the active host
	err := app.Run([]string{app.Name, hostsDeleteCommand.Name, "test-host"})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "delete.success.golden"))
	test.ExpectNoError(t, err)

	if config.Config.ActiveHost != nil {
		t.Error("deleting the active host should make ActiveHost nil, but did not")
	}
}
