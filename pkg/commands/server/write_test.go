package server

import (
	"testing"

	"github.com/gotestyourself/gotestyourself/assert"
	"github.com/gotestyourself/gotestyourself/golden"

	"github.com/vapor-ware/synse-cli/internal/test"
	"github.com/vapor-ware/synse-cli/pkg/config"
)

const (
	// the mocked 200 OK JSON response for the Synse Server 'write' route
	writeRespOK = `
[
  {
    "context":{
      "action":"color",
      "data":"000000"
    },
    "transaction":"b9u6ut6q5i6g020lau70"
  }
]`

	// the mocked 500 error JSON response for the Synse Server 'write' route
	writeRespErr = `
{
  "http_code":500,
  "error_id":0,
  "description":"unknown",
  "timestamp":"2018-03-14 15:34:42.243715",
  "context":"test error."
}`
)

// TestWriteCommandError tests the 'write' command when it is unable to
// connect to the Synse Server instance because the active host is nil.
func TestWriteCommandError(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		ServerCommand.Name,
		writeCommand.Name,
		"rack-1", "board-1", "device-1", "color", "000000",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "error.nil.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestWriteCommandError2 tests the 'write' command when it is unable to
// connect to the Synse Server instance because the active host is not a
// Synse Server instance.
func TestWriteCommandError2(t *testing.T) {
	test.Setup()
	config.Config.ActiveHost = &config.HostConfig{
		Name:    "test-host",
		Address: "localhost:5151",
	}

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		ServerCommand.Name,
		writeCommand.Name,
		"rack-1", "board-1", "device-1", "color", "000000",
	})

	// FIXME: this test fails on CI because the expected output is different
	//     -Get http://localhost:5151/synse/version: dial tcp [::1]:5151: getsockopt: connection refused
	//     +Get http://localhost:5151/synse/version: dial tcp 127.0.0.1:5151: connect: connection refused
	//assert.Assert(t, golden.String(app.ErrBuffer.String(), "error.bad_host.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestWriteCommandError3 tests the 'write' command when no arguments
// are provided, but some are required.
func TestWriteCommandError3(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		ServerCommand.Name,
		writeCommand.Name,
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "write.error.no_args.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestWriteCommandError4 tests the 'write' command when too many
// arguments are provided.
func TestWriteCommandError4(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		ServerCommand.Name,
		writeCommand.Name,
		"rack-1", "board-1", "device-1", "color", "000000", "extra",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "write.error.extra_args.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestWriteCommandRequestError tests the 'write' command when it gets a
// 500 response from Synse Server.
func TestWriteCommandRequestError(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/write/rack-1/board-1/device-1", 500, writeRespErr)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		ServerCommand.Name,
		writeCommand.Name,
		"rack-1", "board-1", "device-1", "color", "000000",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "error.500.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestWriteCommandRequestSuccessYaml tests the 'write' command when it gets
// a 200 response from Synse Server, with YAML output.
func TestWriteCommandRequestSuccessYaml(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/write/rack-1/board-1/device-1", 200, writeRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "yaml",
		ServerCommand.Name,
		writeCommand.Name,
		"rack-1", "board-1", "device-1", "color", "000000",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "write.success.yaml.golden"))
	test.ExpectNoError(t, err)
}

// TestWriteCommandRequestSuccessJson tests the 'write' command when it gets
// a 200 response from Synse Server, with JSON output.
func TestWriteCommandRequestSuccessJson(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/write/rack-1/board-1/device-1", 200, writeRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "json",
		ServerCommand.Name,
		writeCommand.Name,
		"rack-1", "board-1", "device-1", "color", "000000",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "write.success.json.golden"))
	test.ExpectNoError(t, err)
}

// TestWriteCommandRequestSuccessPretty tests the 'write' command when it gets
// a 200 response from Synse Server, with pretty output.
func TestWriteCommandRequestSuccessPretty(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/write/rack-1/board-1/device-1", 200, writeRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "pretty",
		ServerCommand.Name,
		writeCommand.Name,
		"rack-1", "board-1", "device-1", "color", "000000",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "write.success.pretty.golden"))
	test.ExpectNoError(t, err)
}
