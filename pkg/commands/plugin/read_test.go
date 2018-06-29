package plugin

import (
	"bytes"
	"testing"

	"github.com/gotestyourself/gotestyourself/assert"
	"github.com/gotestyourself/gotestyourself/golden"

	"github.com/vapor-ware/synse-cli/internal/test"
	"github.com/vapor-ware/synse-cli/pkg/client"
	"github.com/vapor-ware/synse-server-grpc/go"
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

// TestGetValue tests the getValue function.
func TestGetValue(t *testing.T) {
	tests := []struct {
		in  *synse.Reading
		out interface{}
	}{
		{
			in: &synse.Reading{
				Value: &synse.Reading_StringValue{StringValue: "test"},
			},
			out: "test",
		},
		{
			in: &synse.Reading{
				Value: &synse.Reading_BoolValue{BoolValue: true},
			},
			out: true,
		},
		{
			in: &synse.Reading{
				Value: &synse.Reading_Float32Value{Float32Value: 0.6046603},
			},
			out: float32(0.6046603),
		},
		{
			in: &synse.Reading{
				Value: &synse.Reading_Float64Value{Float64Value: 0.6046602879796196},
			},
			out: float64(0.6046602879796196),
		},
		{
			in: &synse.Reading{
				Value: &synse.Reading_Int32Value{Int32Value: 1298498081},
			},
			out: int32(1298498081),
		},
		{
			in: &synse.Reading{
				Value: &synse.Reading_Int64Value{Int64Value: 5577006791947779410},
			},
			out: int64(5577006791947779410),
		},
		{
			in: &synse.Reading{
				Value: &synse.Reading_Uint32Value{Uint32Value: 2596996162},
			},
			out: uint32(2596996162),
		},
		{
			in: &synse.Reading{
				Value: &synse.Reading_Uint64Value{Uint64Value: 5577006791947779410},
			},
			out: uint64(5577006791947779410),
		},
	}

	for _, test := range tests {
		r := getValue(test.in)
		if r != test.out {
			t.Errorf("getValue(%v) => %v Out:%v", test.in, r, test.out)
		}
	}
}

// TestGetValueBytes tests getValue function where input is a slice of byte.
func TestGetValueBytes(t *testing.T) {
	tests := []struct {
		in  *synse.Reading
		out []byte
	}{
		{
			in: &synse.Reading{
				Value: &synse.Reading_BytesValue{BytesValue: []byte("foo")},
			},
			out: []byte("foo"),
		},
	}

	for _, test := range tests {
		r := getValue(test.in).([]byte)
		if !bytes.Equal(r, test.out) {
			t.Errorf("getValue(%v) => %v Out:%v", test.in, r, test.out)
		}
	}
}
