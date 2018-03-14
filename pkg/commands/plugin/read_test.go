package plugin

import (
	"testing"

	"github.com/gotestyourself/gotestyourself/assert"
	"github.com/gotestyourself/gotestyourself/golden"

	"github.com/vapor-ware/synse-cli/internal/test"
	"github.com/vapor-ware/synse-cli/pkg/client"
)

// TestReadCommandError tests the 'read' command when the plugin
// network info is not specified.
func TestReadCommandError(t *testing.T) {
	client.Grpc.Reset()
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, PluginCommand)

	err := app.Run([]string{
		app.Name,
		PluginCommand.Name,
		pluginReadCommand.Name,
		"rack", "board", "device",
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "read.error.none.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestReadCommandError2 tests the 'read' command when the plugin
// network info is specified as unix, but no backend is present.
func TestReadCommandError2(t *testing.T) {
	client.Grpc.Reset()
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, PluginCommand)

	err := app.Run([]string{
		app.Name,
		PluginCommand.Name,
		"--unix", "tmp/nonexistent",
		pluginReadCommand.Name,
		"rack", "board", "device",
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "read.error.unix.no_backend.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestReadCommandError3 tests the 'read' command when the plugin
// network info is specified as tcp, but no backend is present.
func TestReadCommandError3(t *testing.T) {
	client.Grpc.Reset()
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, PluginCommand)

	err := app.Run([]string{
		app.Name,
		PluginCommand.Name,
		"--tcp", "localhost:5151",
		pluginReadCommand.Name,
		"rack", "board", "device",
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "read.error.tcp.no_backend.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestReadCommandError4 tests the 'read' command when no arguments are given.
func TestReadCommandError4(t *testing.T) {
	client.Grpc.Reset()
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, PluginCommand)

	err := app.Run([]string{
		app.Name,
		PluginCommand.Name,
		"--tcp", "localhost:5151",
		pluginReadCommand.Name,
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "read.error.no_args.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestReadCommandError5 tests the 'read' command when extra arguments are given.
func TestReadCommandError5(t *testing.T) {
	client.Grpc.Reset()
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, PluginCommand)

	err := app.Run([]string{
		app.Name,
		PluginCommand.Name,
		"--tcp", "localhost:5151",
		pluginReadCommand.Name,
		"rack", "board", "device", "extra",
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "read.error.extra_args.golden"))
	test.ExpectExitCoderError(t, err)
}
