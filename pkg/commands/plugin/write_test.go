package plugin

import (
	"testing"

	"github.com/gotestyourself/gotestyourself/assert"
	"github.com/gotestyourself/gotestyourself/golden"

	"github.com/vapor-ware/synse-cli/internal/test"
	"github.com/vapor-ware/synse-cli/pkg/client"
)

// TestWriteCommandError tests the 'write' command when the plugin
// network info is not specified.
func TestWriteCommandError(t *testing.T) {
	client.Grpc.Reset()
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, PluginCommand)

	err := app.Run([]string{
		app.Name,
		PluginCommand.Name,
		pluginWriteCommand.Name,
		"rack", "board", "device", "action", "data",
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "write.error.none.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestWriteCommandError2 tests the 'write' command when the plugin
// network info is specified as unix, but no backend is present.
func TestWriteCommandError2(t *testing.T) {
	client.Grpc.Reset()
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, PluginCommand)

	err := app.Run([]string{
		app.Name,
		PluginCommand.Name,
		"--unix", "tmp/nonexistent",
		pluginWriteCommand.Name,
		"rack", "board", "device", "action", "data",
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "write.error.unix.no_backend.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestWriteCommandError3 tests the 'write' command when the plugin
// network info is specified as tcp, but no backend is present.
func TestWriteCommandError3(t *testing.T) {
	client.Grpc.Reset()
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, PluginCommand)

	err := app.Run([]string{
		app.Name,
		PluginCommand.Name,
		"--tcp", "localhost:5151",
		pluginWriteCommand.Name,
		"rack", "board", "device", "action", "data",
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "write.error.tcp.no_backend.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestWriteCommandError4 tests the 'write' command when no arguments are given.
func TestWriteCommandError4(t *testing.T) {
	client.Grpc.Reset()
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, PluginCommand)

	err := app.Run([]string{
		app.Name,
		PluginCommand.Name,
		"--tcp", "localhost:5151",
		pluginWriteCommand.Name,
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "write.error.no_args.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestWriteCommandError5 tests the 'write' command when extra arguments are given.
func TestWriteCommandError5(t *testing.T) {
	client.Grpc.Reset()
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, PluginCommand)

	err := app.Run([]string{
		app.Name,
		PluginCommand.Name,
		"--tcp", "localhost:5151",
		pluginWriteCommand.Name,
		"rack", "board", "device", "action", "data", "extra",
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "write.error.extra_args.golden"))
	test.ExpectExitCoderError(t, err)
}
