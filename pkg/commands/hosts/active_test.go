package hosts

import (
	"testing"

	"github.com/gotestyourself/gotestyourself/assert"
	"github.com/gotestyourself/gotestyourself/golden"

	"github.com/vapor-ware/synse-cli/internal/test"
	"github.com/vapor-ware/synse-cli/pkg/config"
)

// TestActiveCommandError tests issuing the 'active' command when no
// active host is set.
func TestActiveCommandError(t *testing.T) {
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, HostsCommand)

	// Set the active host to nil
	config.Config.ActiveHost = nil

	err := app.Run([]string{
		app.Name,
		HostsCommand.Name,
		hostsActiveCommand.Name,
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "active.error.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestActiveCommandErrorPretty tests issuing the 'active' command when
// there is an active host and the 'pretty' format is specified. We expect
// this to fail since 'active' does not yet support pretty formatting.
func TestActiveCommandErrorPretty(t *testing.T) {
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, HostsCommand)

	// Set the active host to a HostConfig
	config.Config.ActiveHost = &config.HostConfig{
		Name:    "test-host",
		Address: "test-address",
	}

	err := app.Run([]string{
		app.Name,
		"--format", "pretty",
		HostsCommand.Name,
		hostsActiveCommand.Name,
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "active.error.pretty.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestActiveCommandSuccessYaml tests issuing the 'active' command when
// there is an active host and the 'yaml' format is specified. We expect
// this to succeed.
func TestActiveCommandSuccessYaml(t *testing.T) {
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, HostsCommand)

	// Set the active host to a HostConfig
	config.Config.ActiveHost = &config.HostConfig{
		Name:    "test-host",
		Address: "test-address",
	}

	err := app.Run([]string{
		app.Name,
		"--format", "yaml",
		HostsCommand.Name,
		hostsActiveCommand.Name,
	})

	assert.Assert(t, golden.String(app.OutBuffer.String(), "active.success.yaml.golden"))
	test.ExpectNoError(t, err)
}

// TestActiveCommandSuccessJson tests issuing the 'active' command when
// there is an active host and the 'json' format is specified. We expect
// this to succeed.
func TestActiveCommandSuccessJson(t *testing.T) {
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, HostsCommand)

	// Set the active host to a HostConfig
	config.Config.ActiveHost = &config.HostConfig{
		Name:    "test-host",
		Address: "test-address",
	}

	err := app.Run([]string{
		app.Name,
		"--format", "json",
		HostsCommand.Name,
		hostsActiveCommand.Name,
	})

	assert.Assert(t, golden.String(app.OutBuffer.String(), "active.success.json.golden"))
	test.ExpectNoError(t, err)
}
