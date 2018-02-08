package hosts

import (
	"testing"

	"github.com/vapor-ware/synse-cli/config"
	"github.com/vapor-ware/synse-cli/internal/test"
)

func TestDeleteCommandError(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, hostsDeleteCommand)

	// expects exactly one arg, but none are given
	err := app.Run([]string{app.Name, hostsDeleteCommand.Name})

	test.ExpectExitCoderError(t, err)
}

func TestDeleteCommandSuccess(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, hostsDeleteCommand)

	// we should not fail if we try to delete a host that is not in
	// the configuration
	err := app.Run([]string{app.Name, hostsDeleteCommand.Name, "missing-host"})

	test.ExpectNoError(t, err)
}

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

	test.ExpectNoError(t, err)
}

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

	test.ExpectNoError(t, err)

	if config.Config.ActiveHost != nil {
		t.Error("deleting the active host should make ActiveHost nil, but did not")
	}
}
