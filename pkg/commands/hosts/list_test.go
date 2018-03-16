package hosts

import (
	"testing"

	"github.com/gotestyourself/gotestyourself/assert"
	"github.com/gotestyourself/gotestyourself/golden"

	"github.com/vapor-ware/synse-cli/internal/test"
	"github.com/vapor-ware/synse-cli/pkg/config"
)

// TestListCommandSuccess tests the 'list' command, successfully listing
// the configured hosts when there are no hosts configured.
func TestListCommandSuccess(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, HostsCommand)

	// should succeed with no hosts
	err := app.Run([]string{
		app.Name,
		HostsCommand.Name,
		hostsListCommand.Name,
	})

	assert.Assert(t, golden.String(app.OutBuffer.String(), "list.empty.golden"))
	test.ExpectNoError(t, err)
}

// TestListCommandSuccess2 tests the 'list' command, successfully listing
// the configured hosts when there are hosts configured, but no active hosts.
func TestListCommandSuccess2(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, HostsCommand)

	config.Config.Hosts = map[string]*config.HostConfig{
		"test1": {Name: "test1", Address: "addr"},
		"test2": {Name: "test2", Address: "addr"},
		"test3": {Name: "test3", Address: "addr"},
		"test4": {Name: "test4", Address: "addr"},
	}

	// should succeed with hosts, but no active host
	err := app.Run([]string{
		app.Name,
		HostsCommand.Name,
		hostsListCommand.Name,
	})

	assert.Assert(t, golden.String(app.OutBuffer.String(), "list.success.no_active.golden"))
	test.ExpectNoError(t, err)
}

// TestListCommandSuccess3 tests the 'list' command, successfully listing
// the configured hosts when there are hosts configured, and one is the active
// host.
func TestListCommandSuccess3(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, HostsCommand)

	config.Config.Hosts = map[string]*config.HostConfig{
		"test1": {Name: "test1", Address: "addr"},
		"test2": {Name: "test2", Address: "addr"},
		"test3": {Name: "test3", Address: "addr"},
		"test4": {Name: "test4", Address: "addr"},
	}
	config.Config.ActiveHost = config.Config.Hosts["test1"]

	// should succeed with hosts and an active host
	err := app.Run([]string{
		app.Name,
		HostsCommand.Name,
		hostsListCommand.Name,
	})

	assert.Assert(t, golden.String(app.OutBuffer.String(), "list.success.active.golden"))
	test.ExpectNoError(t, err)
}
