package hosts

import (
	"testing"

	"github.com/vapor-ware/synse-cli/internal/test"
	"github.com/vapor-ware/synse-cli/config"
	"github.com/urfave/cli"
)

func TestAddCommandError(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, hostsAddCommand)

	err := app.Run([]string{app.Name, hostsAddCommand.Name})
	if err == nil {
		t.Error("expected error, but got nil")
	}
	_, ok := err.(cli.ExitCoder)
	if !ok {
		t.Error("expected error to fulfill cli.ExitCoder interface, but does not")
	}
}

func TestAddCommandError2(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, hostsAddCommand)

	// Before adding the "name / address" host, we want it to already exist.
	config.Config.Hosts["name"] = &config.HostConfig{
		Name: "name",
		Address: "address",
	}

	err := app.Run([]string{app.Name, hostsAddCommand.Name, "name", "address"})
	if err == nil {
		t.Error("expected error, but got nil")
	}
	_, ok := err.(cli.ExitCoder)
	if !ok {
		t.Error("expected error to fulfill cli.ExitCoder interface, but does not")
	}
}

func TestAddCommandSuccess(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, hostsAddCommand)

	err := app.Run([]string{app.Name, hostsAddCommand.Name, "name", "address"})
	if err != nil {
		t.Errorf("expected no error but got: %v", err)
	}
}
