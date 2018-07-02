package plugin

import (
	"testing"

	"github.com/gotestyourself/gotestyourself/assert"
	"github.com/gotestyourself/gotestyourself/golden"

	"github.com/vapor-ware/synse-cli/internal/test"
	"github.com/vapor-ware/synse-cli/pkg/client"
)

// TestCapabilitiesCommandError tests the `capabilities` command when the plugin
// network info is not specified.
func TestCapabilitiesCommandError(t *testing.T) {
	client.Grpc.Reset()
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, PluginCommand)

	err := app.Run([]string{
		app.Name,
		PluginCommand.Name,
		pluginCapabilitiesCommand.Name,
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "capabilities.error.none.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestCapabilitiesCommandError2 tests the `capabilities` command when the plugin
// network info is specified as unix, but no backend is present.
func TestCapabilitiesCommandError2(t *testing.T) {
	client.Grpc.Reset()
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, PluginCommand)

	err := app.Run([]string{
		app.Name,
		PluginCommand.Name,
		"--unix", "tmp/nonexistent",
		pluginCapabilitiesCommand.Name,
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "capabilities.error.unix.no_backend.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestCapabilitiesCommandError3 tests the `capabilities` command when the plugin
// network info is specified as tcp, but no backend is present.
func TestCapabilitiesCommandError3(t *testing.T) {
	client.Grpc.Reset()
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, PluginCommand)

	err := app.Run([]string{
		app.Name,
		PluginCommand.Name,
		"--tcp", "localhost:5151",
		pluginCapabilitiesCommand.Name,
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "capabilities.error.tcp.no_backend.golden"))
	test.ExpectExitCoderError(t, err)
}
