package plugin

import (
	"testing"

	"github.com/gotestyourself/gotestyourself/assert"
	"github.com/gotestyourself/gotestyourself/golden"

	"github.com/vapor-ware/synse-cli/internal/test"
	"github.com/vapor-ware/synse-cli/pkg/client"
)

// TestHealthCommandError tests the `health` command when the plugin
// network info is not specified.
func TestHealthCommandError(t *testing.T) {
	client.Grpc.Reset()
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, PluginCommand)

	err := app.Run([]string{
		app.Name,
		PluginCommand.Name,
		pluginHealthCommand.Name,
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "health.error.none.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestHealthCommandError2 tests the `health` command when the plugin
// network info is specified as unix, but no backend is present.
func TestHealthCommandError2(t *testing.T) {
	client.Grpc.Reset()
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, PluginCommand)

	err := app.Run([]string{
		app.Name,
		PluginCommand.Name,
		"--unix", "tmp/nonexistent",
		pluginHealthCommand.Name,
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "health.error.unix.no_backend.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestHealthCommandError3 tests the `health` command when the plugin
// network info is specified as tcp, but no backend is present.
func TestHealthCommandError3(t *testing.T) {
	client.Grpc.Reset()
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, PluginCommand)

	err := app.Run([]string{
		app.Name,
		PluginCommand.Name,
		"--tcp", "localhost:5151",
		pluginHealthCommand.Name,
	})

	// FIXME: Refer to `rpc error` comment on #181.
	// assert.Assert(t, golden.String(app.ErrBuffer.String(), "health.error.tcp.no_backend.golden"))
	test.ExpectExitCoderError(t, err)
}
