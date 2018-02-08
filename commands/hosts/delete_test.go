package hosts

import (
	"testing"

	"github.com/vapor-ware/synse-cli/internal/test"
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/config"
)

func TestDeleteCommandError(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, hostsDeleteCommand)

	// expects exactly one arg, but none are given
	err := app.Run([]string{app.Name, hostsDeleteCommand.Name})
	if err == nil {
		t.Error("expected error, but got nil")
	}
	_, ok := err.(cli.ExitCoder)
	if !ok {
		t.Error("expected error to fulfill cli.ExitCoder interface, but does not")
	}
}

func TestDeleteCommandSuccess(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, hostsDeleteCommand)

	// we should not fail if we try to delete a host that is not in
	// the configuration
	err := app.Run([]string{app.Name, hostsDeleteCommand.Name, "missing-host"})
	if err != nil {
		t.Errorf("expected no error but got: %v", err)
	}
}

func TestDeleteCommandSuccess2(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, hostsDeleteCommand)

	config.Config.Hosts["test-host"] = &config.HostConfig{
		Name: "test-host",
		Address: "test-address",
	}

	// we should not fail if we try to delete a host that is in the
	// configuration
	err := app.Run([]string{app.Name, hostsDeleteCommand.Name, "test-host"})
	if err != nil {
		t.Errorf("expected no error but got: %v", err)
	}
}

func TestDeleteCommandSuccess3(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, hostsDeleteCommand)

	host := &config.HostConfig{
		Name: "test-host",
		Address: "test-address",
	}
	config.Config.Hosts["test-host"] = host
	config.Config.ActiveHost = host

	// we should not fail if we try to delete a host that is in the
	// configuration and is also the active host
	err := app.Run([]string{app.Name, hostsDeleteCommand.Name, "test-host"})
	if err != nil {
		t.Errorf("expected no error but got: %v", err)
	}

	if config.Config.ActiveHost != nil {
		t.Error("deleting the active host should make ActiveHost nil, but did not")
	}
}
