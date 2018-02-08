package hosts

import (
	"testing"

	"github.com/vapor-ware/synse-cli/config"
	"github.com/vapor-ware/synse-cli/internal/test"
)

func TestListCommandSuccess(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, hostsListCommand)

	// should succeed with no hosts
	err := app.Run([]string{app.Name, hostsListCommand.Name})

	test.ExpectNoError(t, err)
}

func TestListCommandSuccess2(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, hostsListCommand)

	config.Config.Hosts = map[string]*config.HostConfig{
		"test1": {Name: "test1", Address: "addr"},
		"test2": {Name: "test2", Address: "addr"},
		"test3": {Name: "test3", Address: "addr"},
		"test4": {Name: "test4", Address: "addr"},
	}

	// should succeed with hosts, but no active host
	err := app.Run([]string{app.Name, hostsListCommand.Name})

	test.ExpectNoError(t, err)
}

func TestListCommandSuccess3(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, hostsListCommand)

	config.Config.Hosts = map[string]*config.HostConfig{
		"test1": {Name: "test1", Address: "addr"},
		"test2": {Name: "test2", Address: "addr"},
		"test3": {Name: "test3", Address: "addr"},
		"test4": {Name: "test4", Address: "addr"},
	}
	config.Config.ActiveHost = config.Config.Hosts["test1"]

	// should succeed with hosts and an active host
	err := app.Run([]string{app.Name, hostsListCommand.Name})

	test.ExpectNoError(t, err)
}
