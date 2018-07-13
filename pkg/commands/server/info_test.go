package server

import (
	"testing"

	"github.com/gotestyourself/gotestyourself/assert"
	"github.com/gotestyourself/gotestyourself/golden"

	"github.com/vapor-ware/synse-cli/internal/test"
	"github.com/vapor-ware/synse-cli/pkg/config"
)

const (
	// the mocked 200 OK JSON response for the Synse Server rack 'info' route
	infoRackRespOK = `
{
  "rack":"rack-1",
  "boards":[
    "board-1"
  ]
}`

	// the mocked 200 OK JSON response for the Synse Server board 'info' route
	infoBoardRespOK = `
{
  "board":"board-1",
  "location":{
    "rack":"rack-1"
  },
  "devices":[
    "device-1",
    "device-2",
    "device-3",
    "device-4",
    "device-5",
    "device-6",
    "device-7",
    "device-8"
  ]
}`

	// the mocked 200 OK JSON response for the Synse Server device 'info' route
	infoDeviceRespOK = `
{
  "timestamp":"2018-06-28T12:59:47.625842798Z",
  "uid":"device-1",
  "kind":"pressure",
  "metadata":{
    "model":"emul8-pressure"
  },
  "plugin":"emulator-plugin",
  "info":"Synse Pressure Sensor 1",
  "location":{
    "rack":"rack-1",
    "board":"board-1"
  },
  "output":[
    {
      "name":"pressure",
      "type":"pressure",
      "precision":3,
      "scaling_factor":1.5,
      "unit":{
        "name":"pascals",
        "symbol":"Pa"
      }
    }
  ]
}`

	// the mocked 500 error JSON response for the Synse Server 'info' route
	infoRespErr = `
{
  "http_code":500,
  "error_id":0,
  "description":"unknown",
  "timestamp":"2018-03-14 15:34:42.243715",
  "context":"test error."
}`
)

// TestInfoCommandError tests the 'info' command when it is unable to
// connect to the Synse Server instance because the active host is nil.
func TestInfoCommandError(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		ServerCommand.Name,
		infoCommand.Name,
		"rack-1", "board-1", "device-1",
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "error.nil.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestInfoCommandError2 tests the 'info' command when it is unable to
// connect to the Synse Server instance because the active host is not a
// Synse Server instance.
func TestInfoCommandError2(t *testing.T) {
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
		infoCommand.Name,
		"rack-1", "board-1", "device-1",
	})

	// FIXME: this test fails on CI because the expected output is different
	//     -Get http://localhost:5151/synse/version: dial tcp [::1]:5151: getsockopt: connection refused
	//     +Get http://localhost:5151/synse/version: dial tcp 127.0.0.1:5151: connect: connection refused
	//assert.Assert(t, golden.String(app.ErrBuffer.String(), "error.bad_host.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestInfoCommandError3 tests the 'info' command when no arguments
// are provided, but some are required.
func TestInfoCommandError3(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		ServerCommand.Name,
		infoCommand.Name,
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "info.error.no_args.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestInfoCommandError4 tests the 'info' command when too many
// arguments are provided.
func TestInfoCommandError4(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		ServerCommand.Name,
		infoCommand.Name,
		"rack-1", "board-1", "device-1", "extra",
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "info.error.extra_args.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestInfoCommandRequestErrorRack tests the 'info' command when it gets a
// 500 response from Synse Server when querying for rack info.
func TestInfoCommandRequestErrorRack(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/2.0/info/rack-1", 500, infoRespErr)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		ServerCommand.Name,
		infoCommand.Name,
		"rack-1",
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "error.500.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestInfoCommandRequestErrorBoard tests the 'info' command when it gets a
// 500 response from Synse Server when querying for board info.
func TestInfoCommandRequestErrorBoard(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/2.0/info/rack-1/board-1", 500, infoRespErr)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		ServerCommand.Name,
		infoCommand.Name,
		"rack-1", "board-1",
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "error.500.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestInfoCommandRequestErrorDevice tests the 'info' command when it gets a
// 500 response from Synse Server when querying for device info.
func TestInfoCommandRequestErrorDevice(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/2.0/info/rack-1/board-1/device-1", 500, infoRespErr)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		ServerCommand.Name,
		infoCommand.Name,
		"rack-1", "board-1", "device-1",
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "error.500.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestInfoCommandRequestErrorRackPretty tests the 'info' command when it gets
// a 200 response from Synse Server rack request, with pretty output.
func TestInfoCommandRequestErrorRackPretty(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/2.0/info/rack-1", 200, infoRackRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "pretty",
		ServerCommand.Name,
		infoCommand.Name,
		"rack-1",
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "info.error.pretty.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestInfoCommandRequestErrorBoardPretty tests the 'info' command when it gets
// a 200 response from Synse Server board request, with pretty output.
func TestInfoCommandRequestErrorBoardPretty(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/2.0/info/rack-1/board-1", 200, infoBoardRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "pretty",
		ServerCommand.Name,
		infoCommand.Name,
		"rack-1", "board-1",
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "info.error.pretty.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestInfoCommandRequestErrorDevicePretty tests the 'info' command when it gets
// a 200 response from Synse Server device request, with pretty output.
func TestInfoCommandRequestErrorDevicePretty(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/2.0/info/rack-1/board-1/device-1", 200, infoDeviceRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "pretty",
		ServerCommand.Name,
		infoCommand.Name,
		"rack-1", "board-1", "device-1",
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "info.error.pretty.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestInfoCommandRequestSuccessRackYaml tests the 'info' command when it gets
// a 200 response from Synse Server rack request, with YAML output.
func TestInfoCommandRequestSuccessRackYaml(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/2.0/info/rack-1", 200, infoRackRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "yaml",
		ServerCommand.Name,
		infoCommand.Name,
		"rack-1",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "info.success.yaml.rack.golden"))
	test.ExpectNoError(t, err)
}

// TestInfoCommandRequestSuccessRackJson tests the 'info' command when it gets
// a 200 response from Synse Server rack request, with JSON output.
func TestInfoCommandRequestSuccessRackJson(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/2.0/info/rack-1", 200, infoRackRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "json",
		ServerCommand.Name,
		infoCommand.Name,
		"rack-1",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "info.success.json.rack.golden"))
	test.ExpectNoError(t, err)
}

// TestInfoCommandRequestSuccessBoardYaml tests the 'info' command when it gets
// a 200 response from Synse Server board request, with YAML output.
func TestInfoCommandRequestSuccessBoardYaml(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/2.0/info/rack-1/board-1", 200, infoBoardRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "yaml",
		ServerCommand.Name,
		infoCommand.Name,
		"rack-1", "board-1",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "info.success.yaml.board.golden"))
	test.ExpectNoError(t, err)
}

// TestInfoCommandRequestSuccessBoardJson tests the 'info' command when it gets
// a 200 response from Synse Server board request, with JSON output.
func TestInfoCommandRequestSuccessBoardJson(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/2.0/info/rack-1/board-1", 200, infoBoardRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "json",
		ServerCommand.Name,
		infoCommand.Name,
		"rack-1", "board-1",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "info.success.json.board.golden"))
	test.ExpectNoError(t, err)
}

// TestInfoCommandRequestSuccessDeviceYaml tests the 'info' command when it gets
// a 200 response from Synse Server device request, with YAML output.
func TestInfoCommandRequestSuccessDeviceYaml(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/2.0/info/rack-1/board-1/device-1", 200, infoDeviceRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "yaml",
		ServerCommand.Name,
		infoCommand.Name,
		"rack-1", "board-1", "device-1",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "info.success.yaml.device.golden"))
	test.ExpectNoError(t, err)
}

// TestInfoCommandRequestSuccessDeviceJson tests the 'info' command when it gets
// a 200 response from Synse Server device request, with JSON output.
func TestInfoCommandRequestSuccessDeviceJson(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/2.0/info/rack-1/board-1/device-1", 200, infoDeviceRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "json",
		ServerCommand.Name,
		infoCommand.Name,
		"rack-1", "board-1", "device-1",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "info.success.json.device.golden"))
	test.ExpectNoError(t, err)
}
