package plugin

import (
	"testing"

	"github.com/gotestyourself/gotestyourself/assert"
	"github.com/gotestyourself/gotestyourself/golden"

	"github.com/vapor-ware/synse-cli/internal/test"
	"github.com/vapor-ware/synse-cli/pkg/client"
)

// TestMetainfoCommandError tests the 'metainfo' command when the plugin
// network info is not specified.
func TestMetainfoCommandError(t *testing.T) {
	client.Grpc.Reset()
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, PluginCommand)

	err := app.Run([]string{
		app.Name,
		PluginCommand.Name,
		pluginMetainfoCommand.Name,
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "metainfo.error.none.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestMetainfoCommandError2 tests the 'metainfo' command when the plugin
// network info is specified as unix, but no backend is present.
func TestMetainfoCommandError2(t *testing.T) {
	client.Grpc.Reset()
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, PluginCommand)

	err := app.Run([]string{
		app.Name,
		PluginCommand.Name,
		"--unix", "tmp/nonexistent",
		pluginMetainfoCommand.Name,
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "metainfo.error.unix.no_backend.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestMetainfoCommandError3 tests the 'metainfo' command when the plugin
// network info is specified as tcp, but no backend is present.
func TestMetainfoCommandError3(t *testing.T) {
	client.Grpc.Reset()
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, PluginCommand)

	err := app.Run([]string{
		app.Name,
		PluginCommand.Name,
		"--tcp", "localhost:5151",
		pluginMetainfoCommand.Name,
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "metainfo.error.tcp.no_backend.golden"))
	test.ExpectExitCoderError(t, err)
}
