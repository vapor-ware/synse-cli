package plugin

import (
	"testing"

	"github.com/gotestyourself/gotestyourself/assert"
	"github.com/gotestyourself/gotestyourself/golden"

	"github.com/vapor-ware/synse-cli/internal/test"
	"github.com/vapor-ware/synse-cli/pkg/client"
)

// TestReadCachedCommandError tests the 'readcached' command when the plugin
// network info is not specified.
func TestReadCachedCommandError(t *testing.T) {
	client.Grpc.Reset()
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, PluginCommand)

	err := app.Run([]string{
		app.Name,
		PluginCommand.Name,
		pluginReadCachedCommand.Name,
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "readcached.error.none.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestReadCachedCommandError2 tests the 'readcached' command when the plugin
// network info is specified as unix, but no backend is present.
func TestReadCachedCommandError2(t *testing.T) {
	client.Grpc.Reset()
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, PluginCommand)

	err := app.Run([]string{
		app.Name,
		PluginCommand.Name,
		"--unix", "tmp/nonexistent",
		pluginReadCachedCommand.Name,
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "readcached.error.unix.no_backend.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestReadCachedCommandError3 tests the 'readcached' command when the plugin
// network info is specified as tcp, but no backend is present.
func TestReadCachedCommandError3(t *testing.T) {
	client.Grpc.Reset()
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, PluginCommand)

	err := app.Run([]string{
		app.Name,
		PluginCommand.Name,
		"--tcp", "localhost:5151",
		pluginReadCachedCommand.Name,
	})

	// FIXME: Refer to `rpc error` comment on #181.
	// assert.Assert(t, golden.String(app.ErrBuffer.String(), "readcached.error.tcp.no_backend.golden"))
	test.ExpectExitCoderError(t, err)
}
