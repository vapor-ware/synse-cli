package plugin

import (
	"testing"

	"github.com/gotestyourself/gotestyourself/assert"
	"github.com/gotestyourself/gotestyourself/golden"

	"github.com/vapor-ware/synse-cli/internal/test"
	"github.com/vapor-ware/synse-cli/pkg/client"
)

// TestDevicesCommandError tests the 'devices' command when the plugin
// network info is not specified.
func TestDevicesCommandError(t *testing.T) {
	client.Grpc.Reset()
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, PluginCommand)

	err := app.Run([]string{
		app.Name,
		PluginCommand.Name,
		pluginDevicesCommand.Name,
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "devices.error.none.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestDevicesCommandError2 tests the 'devices' command when the plugin
// network info is specified as unix, but no backend is present.
func TestDevicesCommandError2(t *testing.T) {
	client.Grpc.Reset()
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, PluginCommand)

	err := app.Run([]string{
		app.Name,
		PluginCommand.Name,
		"--unix", "tmp/nonexistent",
		pluginDevicesCommand.Name,
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "devices.error.unix.no_backend.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestDevicesCommandError3 tests the 'devices' command when the plugin
// network info is specified as tcp, but no backend is present.
func TestDevicesCommandError3(t *testing.T) {
	client.Grpc.Reset()
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, PluginCommand)

	err := app.Run([]string{
		app.Name,
		PluginCommand.Name,
		"--tcp", "localhost:5151",
		pluginDevicesCommand.Name,
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "devices.error.tcp.no_backend.golden"))
	test.ExpectExitCoderError(t, err)
}
