package hosts

import (
	"testing"

	"github.com/vapor-ware/synse-cli/internal/test"
	"github.com/vapor-ware/synse-cli/config"
	"github.com/urfave/cli"
)

func TestActiveCommandError(t *testing.T) {

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, hostsActiveCommand)

	// Set the active host to nil
	config.Config.ActiveHost = nil

	err := app.Run([]string{app.Name, hostsActiveCommand.Name})
	if err == nil {
		t.Error("expected error, but got nil")
	}
	_, ok := err.(cli.ExitCoder)
	if !ok {
		t.Error("expected error to fulfill cli.ExitCoder interface, but does not")
	}
}

func TestActiveCommandSuccess(t *testing.T) {

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, hostsActiveCommand)

	// Set the active host to a HostConfig
	config.Config.ActiveHost = &config.HostConfig{
		Name: "test-host",
		Address: "test-address",
	}

	err := app.Run([]string{app.Name, hostsActiveCommand.Name})
	if err != nil {
		t.Errorf("expected no error but got: %v", err)
	}
}
