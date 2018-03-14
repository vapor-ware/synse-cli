package plugin

import (
	"testing"

	"github.com/gotestyourself/gotestyourself/assert"
	"github.com/gotestyourself/gotestyourself/golden"

	"github.com/vapor-ware/synse-cli/internal/test"
	"github.com/vapor-ware/synse-cli/pkg/client"
)

// TestTransactionCommandError tests the 'transaction' command when the plugin
// network info is not specified.
func TestTransactionCommandError(t *testing.T) {
	client.Grpc.Reset()
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, PluginCommand)

	err := app.Run([]string{
		app.Name,
		PluginCommand.Name,
		pluginTransactionCommand.Name,
		"transaction-id",
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "transaction.error.none.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestTransactionCommandError2 tests the 'transaction' command when the plugin
// network info is specified as unix, but no backend is present.
func TestTransactionCommandError2(t *testing.T) {
	client.Grpc.Reset()
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, PluginCommand)

	err := app.Run([]string{
		app.Name,
		PluginCommand.Name,
		"--unix", "tmp/nonexistent",
		pluginTransactionCommand.Name,
		"transaction-id",
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "transaction.error.unix.no_backend.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestTransactionCommandError3 tests the 'transaction' command when the plugin
// network info is specified as tcp, but no backend is present.
func TestTransactionCommandError3(t *testing.T) {
	client.Grpc.Reset()
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, PluginCommand)

	err := app.Run([]string{
		app.Name,
		PluginCommand.Name,
		"--tcp", "localhost:5151",
		pluginTransactionCommand.Name,
		"transaction-id",
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "transaction.error.tcp.no_backend.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestTransactionCommandError4 tests the 'transaction' command when no arguments are given.
func TestTransactionCommandError4(t *testing.T) {
	client.Grpc.Reset()
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, PluginCommand)

	err := app.Run([]string{
		app.Name,
		PluginCommand.Name,
		"--tcp", "localhost:5151",
		pluginTransactionCommand.Name,
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "transaction.error.no_args.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestTransactionCommandError5 tests the 'transaction' command when extra arguments are given.
func TestTransactionCommandError5(t *testing.T) {
	client.Grpc.Reset()
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, PluginCommand)

	err := app.Run([]string{
		app.Name,
		PluginCommand.Name,
		"--tcp", "localhost:5151",
		pluginTransactionCommand.Name,
		"transaction-id", "extra",
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "transaction.error.extra_args.golden"))
	test.ExpectExitCoderError(t, err)
}
