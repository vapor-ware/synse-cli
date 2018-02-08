package hosts

import (
	"testing"

	"github.com/vapor-ware/synse-cli/config"
	"github.com/vapor-ware/synse-cli/internal/test"
)

func TestChangeCommandError(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, hostsChangeCommand)

	// expects exactly one arg, here passing none
	err := app.Run([]string{app.Name, hostsChangeCommand.Name})

	test.ExpectExitCoderError(t, err)
}

func TestChangeCommandError2(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, hostsChangeCommand)

	// specifying a host that does not exist - should cause failure
	err := app.Run([]string{app.Name, hostsChangeCommand.Name, "missing-host"})

	test.ExpectExitCoderError(t, err)
}

func TestChangeCommandSuccess(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, hostsChangeCommand)

	config.Config.Hosts["test-host"] = &config.HostConfig{
		Name:    "test-host",
		Address: "test-address",
	}

	// before running, check that the active host is nil
	if config.Config.ActiveHost != nil {
		t.Error("expected active host to be nil at test start")
	}

	err := app.Run([]string{app.Name, hostsChangeCommand.Name, "test-host"})

	test.ExpectNoError(t, err)

	// after running, check that the active host is now set
	if config.Config.ActiveHost != config.Config.Hosts["test-host"] {
		t.Error("active host was not set to the 'test-host'")
	}
}
