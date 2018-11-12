package server

import (
	"testing"

	"github.com/gotestyourself/gotestyourself/assert"
	"github.com/gotestyourself/gotestyourself/golden"

	"github.com/vapor-ware/synse-cli/internal/test"
	"github.com/vapor-ware/synse-cli/pkg/config"
)

const (
	// the temperature mocked 200 OK JSON response for the Synse Server 'read' route
	temperatureReadRespOK = `
{
  "kind":"temperature",
  "data":[
    {
      "value":"65",
      "timestamp":"2018-06-28T12:41:50.333443322Z",
      "unit":{
        "symbol":"C",
        "name":"celsius"
      },
      "type":"temperature",
      "info":"mock temperature response"
    }
  ]
}`

	// the led mocked 200 OK JSON response for the Synse Server 'read' route
	ledReadRespOK = `
{
  "kind":"led",
  "data":[
    {
      "value":"off",
      "timestamp":"2018-06-28T12:41:50.333443322Z",
      "unit":null,
      "type":"state",
      "info":"mock led.state response"
    },
    {
      "value":"000000",
      "timestamp":"2018-06-28T12:41:50.333443322Z",
      "unit":null,
      "type":"color",
      "info":"mock led.color response"
    }
  ]
}`

	// the fan mocked 200 OK JSON response for the Synse Server 'read' route
	fanReadRespOK = `
{
  "kind":"fan",
  "data":[
    {
      "value":"0",
      "timestamp":"2018-06-28T12:41:50.333443322Z",
      "unit":{
        "symbol":"RPM",
        "name":"revolutions per minute"
      },
      "type":"speed",
      "info":"mock fan.speed response"
    }
  ]
}`
)

// TestReadCommandError tests the 'read' command when it is unable to
// connect to the Synse Server instance because the active host is nil.
func TestReadCommandError(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		ServerCommand.Name,
		readCommand.Name,
		"rack-1", "board-1", "device-1",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "error.nil.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestReadCommandError2 tests the 'read' command when it is unable to
// connect to the Synse Server instance because the active host is not a
// Synse Server instance.
func TestReadCommandError2(t *testing.T) {
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
		readCommand.Name,
		"rack-1", "board-1", "device-1",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	// FIXME: this test fails on CI because the expected output is different
	//     -Get http://localhost:5151/synse/version: dial tcp [::1]:5151: getsockopt: connection refused
	//     +Get http://localhost:5151/synse/version: dial tcp 127.0.0.1:5151: connect: connection refused
	//assert.Assert(t, golden.String(app.ErrBuffer.String(), "error.bad_host.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestReadCommandError4 tests the 'read' command when too many
// arguments are provided.
func TestReadCommandError4(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		ServerCommand.Name,
		readCommand.Name,
		"rack-1", "board-1", "device-1", "extra",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "read.error.extra_args.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestReadCommandRequestNoArgsSuccessYaml tests the 'read' command using no
// arguments when it gets a 200 response from Synse Server, with YAML output.
func TestReadCommandRequestNoArgsSuccessYaml(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/scan", 200, scanRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-1", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-2", 200, fanReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-3", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-4", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-5", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-6", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-7", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-8", 200, temperatureReadRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "yaml",
		ServerCommand.Name,
		readCommand.Name,
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "read.success.yaml.no_args.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCommandRequestNoArgsSuccessJson tests the 'read' command using no
// arguments when it gets a 200 response from Synse Server, with JSON output.
func TestReadCommandRequestNoArgsSuccessJson(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/scan", 200, scanRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-1", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-2", 200, fanReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-3", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-4", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-5", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-6", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-7", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-8", 200, temperatureReadRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "json",
		ServerCommand.Name,
		readCommand.Name,
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "read.success.json.no_args.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCommandRequestNoArgsSuccessPretty tests the 'read' command using no
// arguments when it gets a 200 response from Synse Server, with pretty output.
func TestReadCommandRequestNoArgsSuccessPretty(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/scan", 200, scanRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-1", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-2", 200, fanReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-3", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-4", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-5", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-6", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-7", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-8", 200, temperatureReadRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "pretty",
		ServerCommand.Name,
		readCommand.Name,
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "read.success.pretty.no_args.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCommandRequestRackLevelError tests the 'read' command when it gets
// a 500 rack-level error response from Synse Server.
func TestReadCommandRequestRackLevelError(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/info/rack-1", 500, infoRespErr)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		ServerCommand.Name,
		readCommand.Name,
		"rack-1",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "error.500.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestReadCommandRequestRackLevelSuccessYaml tests the 'read' command when it gets
// a 200 rack-level response from Synse Server, with YAML output.
func TestReadCommandRequestRackLevelSuccessYaml(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/scan", 200, scanRespOK)
	test.Serve(t, mux, "/synse/v2/info/rack-1", 200, infoRackRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-1", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-2", 200, fanReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-3", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-4", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-5", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-6", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-7", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-8", 200, temperatureReadRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "yaml",
		ServerCommand.Name,
		readCommand.Name,
		"rack-1",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "read.success.yaml.rack.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCommandRequestRackLevelSuccessJson tests the 'read' command when it gets
// a 200 rack-level response from Synse Server, with JSON output.
func TestReadCommandRequestRackLevelSuccessJson(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/scan", 200, scanRespOK)
	test.Serve(t, mux, "/synse/v2/info/rack-1", 200, infoRackRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-1", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-2", 200, fanReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-3", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-4", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-5", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-6", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-7", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-8", 200, temperatureReadRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "json",
		ServerCommand.Name,
		readCommand.Name,
		"rack-1",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "read.success.json.rack.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCommandRequestRackLevelSuccessPretty tests the 'read' command when it gets
// a 200 rack-level response from Synse Server, with pretty output.
func TestReadCommandRequestRackLevelSuccessPretty(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/scan", 200, scanRespOK)
	test.Serve(t, mux, "/synse/v2/info/rack-1", 200, infoRackRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-1", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-2", 200, fanReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-3", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-4", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-5", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-6", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-7", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-8", 200, temperatureReadRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "pretty",
		ServerCommand.Name,
		readCommand.Name,
		"rack-1",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "read.success.pretty.rack.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCommandRequestBoardLevelError tests the 'read' command when it gets
// a 500 board-level response from Synse Server.
func TestReadCommandRequestBoardLevelError(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/info/rack-1/board-1", 500, infoRespErr)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		ServerCommand.Name,
		readCommand.Name,
		"rack-1", "board-1",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "error.500.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestReadCommandRequestBoardLevelSuccessYaml tests the 'read' command when it gets
// a 200 board-level response from Synse Server, with YAML output.
func TestReadCommandRequestBoardLevelSuccessYaml(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/scan", 200, scanRespOK)
	test.Serve(t, mux, "/synse/v2/info/rack-1/board-1", 200, infoBoardRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-1", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-2", 200, fanReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-3", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-4", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-5", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-6", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-7", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-8", 200, temperatureReadRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "yaml",
		ServerCommand.Name,
		readCommand.Name,
		"rack-1", "board-1",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "read.success.yaml.board.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCommandRequestBoardLevelSuccessJson tests the 'read' command when it gets
// a 200 board-level response from Synse Server, with JSON output.
func TestReadCommandRequestBoardLevelSuccessJson(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/scan", 200, scanRespOK)
	test.Serve(t, mux, "/synse/v2/info/rack-1/board-1", 200, infoBoardRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-1", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-2", 200, fanReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-3", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-4", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-5", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-6", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-7", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-8", 200, temperatureReadRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "json",
		ServerCommand.Name,
		readCommand.Name,
		"rack-1", "board-1",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "read.success.json.board.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCommandRequestBoardLevelSuccessPretty tests the 'read' command when it gets
// a 200 board-level response from Synse Server, with pretty output.
func TestReadCommandRequestBoardLevelSuccessPretty(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/scan", 200, scanRespOK)
	test.Serve(t, mux, "/synse/v2/info/rack-1/board-1", 200, infoBoardRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-1", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-2", 200, fanReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-3", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-4", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-5", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-6", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-7", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-8", 200, temperatureReadRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "pretty",
		ServerCommand.Name,
		readCommand.Name,
		"rack-1", "board-1",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "read.success.pretty.board.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCommandRequestDeviceLevelError tests the 'read' command when it gets
// a 500 device-level response from Synse Server.
func TestReadCommandRequestDeviceLevelError(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/info/rack-1/board-1/device-1", 500, infoRespErr)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		ServerCommand.Name,
		readCommand.Name,
		"rack-1", "board-1", "device-1",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "error.500.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestReadCommandRequestDeviceLevelSuccessYaml tests the 'read' command when it gets
// a 200 device-level response from Synse Server, with YAML output.
func TestReadCommandRequestDeviceLevelSuccessYaml(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/scan", 200, scanRespOK)
	test.Serve(t, mux, "/synse/v2/info/rack-1/board-1/device-1", 200, infoDeviceRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-1", 200, temperatureReadRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "yaml",
		ServerCommand.Name,
		readCommand.Name,
		"rack-1", "board-1", "device-1",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "read.success.yaml.device.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCommandRequestDeviceLevelSuccessJson tests the 'read' command when it gets
// a 200 device-level response from Synse Server, with JSON output.
func TestReadCommandRequestDeviceLevelSuccessJson(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/scan", 200, scanRespOK)
	test.Serve(t, mux, "/synse/v2/info/rack-1/board-1/device-1", 200, infoDeviceRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-1", 200, temperatureReadRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "json",
		ServerCommand.Name,
		readCommand.Name,
		"rack-1", "board-1", "device-1",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "read.success.json.device.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCommandRequestDeviceLevelSuccessPretty tests the 'read' command when it gets
// a 200 device-level response from Synse Server, with pretty output.
func TestReadCommandRequestDeviceLevelSuccessPretty(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/scan", 200, scanRespOK)
	test.Serve(t, mux, "/synse/v2/info/rack-1/board-1/device-1", 200, infoDeviceRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-1", 200, temperatureReadRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "pretty",
		ServerCommand.Name,
		readCommand.Name,
		"rack-1", "board-1", "device-1",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "read.success.pretty.device.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCommandRequestRackLevelSingleFlagSuccessYaml tests the 'read'
// command using a single 'type' flag at rack-level when it gets a 200 response
// from Synse Server, with YAML output.
func TestReadCommandRequestRackLevelSingleFlagSuccessYaml(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/scan", 200, scanRespOK)
	test.Serve(t, mux, "/synse/v2/info/rack-1", 200, infoBoardRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-1", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-2", 200, fanReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-3", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-4", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-5", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-6", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-7", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-8", 200, temperatureReadRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "yaml",
		ServerCommand.Name,
		readCommand.Name,
		"rack-1",
		"--type", "temperature",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "read.success.yaml.single_flag.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCommandRequestRackLevelSingleFlagSuccessJson tests the 'read'
// command using a single 'type' flag at rack-level when it gets a 200 response
// from Synse Server, with JSON output.
func TestReadCommandRequestRackLevelSingleFlagSuccessJson(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/scan", 200, scanRespOK)
	test.Serve(t, mux, "/synse/v2/info/rack-1", 200, infoBoardRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-1", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-2", 200, fanReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-3", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-4", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-5", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-6", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-7", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-8", 200, temperatureReadRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "json",
		ServerCommand.Name,
		readCommand.Name,
		"rack-1",
		"--type", "temperature",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "read.success.json.single_flag.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCommandRequestRackLevelSingleFlagSuccessPretty tests the 'read'
// command using a single 'type' flag at rack-level when it gets a 200 response
// from Synse Server, with pretty output.
func TestReadCommandRequestRackLevelSingleFlagSuccessPretty(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/scan", 200, scanRespOK)
	test.Serve(t, mux, "/synse/v2/info/rack-1", 200, infoBoardRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-1", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-2", 200, fanReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-3", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-4", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-5", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-6", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-7", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-8", 200, temperatureReadRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "pretty",
		ServerCommand.Name,
		readCommand.Name,
		"rack-1",
		"--type", "temperature",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "read.success.pretty.single_flag.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCommandRequestBoardLevelSingleFlagSuccessYaml tests the 'read'
// command using a single 'type' flag at board-level when it gets a 200 response
// from Synse Server, with YAML output.
func TestReadCommandRequestBoardLevelSingleFlagSuccessYaml(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/scan", 200, scanRespOK)
	test.Serve(t, mux, "/synse/v2/info/rack-1/board-1", 200, infoBoardRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-1", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-2", 200, fanReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-3", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-4", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-5", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-6", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-7", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-8", 200, temperatureReadRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "yaml",
		ServerCommand.Name,
		readCommand.Name,
		"rack-1", "board-1",
		"--type", "temperature",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "read.success.yaml.single_flag.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCommandRequestBoardLevelSingleFlagSuccessJson tests the 'read'
// command using a single 'type' flag at board-level when it gets a 200 response
// from Synse Server, with JSON output.
func TestReadCommandRequestBoardLevelSingleFlagSuccessJson(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/scan", 200, scanRespOK)
	test.Serve(t, mux, "/synse/v2/info/rack-1/board-1", 200, infoBoardRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-1", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-2", 200, fanReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-3", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-4", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-5", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-6", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-7", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-8", 200, temperatureReadRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "json",
		ServerCommand.Name,
		readCommand.Name,
		"rack-1", "board-1",
		"--type", "temperature",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "read.success.json.single_flag.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCommandRequestBoardLevelSingleFlagSuccessPretty tests the 'read'
// command using a single 'type' flag at board-level when it gets a 200 response
// from Synse Server, with pretty output.
func TestReadCommandRequestBoardLevelSingleFlagSuccessPretty(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/scan", 200, scanRespOK)
	test.Serve(t, mux, "/synse/v2/info/rack-1/board-1", 200, infoBoardRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-1", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-2", 200, fanReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-3", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-4", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-5", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-6", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-7", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-8", 200, temperatureReadRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "pretty",
		ServerCommand.Name,
		readCommand.Name,
		"rack-1", "board-1",
		"--type", "temperature",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "read.success.pretty.single_flag.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCommandRequestRackLevelMutipleFlagsSuccessYaml tests the 'read'
// command using multiple 'type' flags at rack-level when it gets a 200 response
// from Synse Server, with YAML output.
// FIXME: The type we get from `scan` is different from the type we get from
// `read`. That affects the way we do filtering devices. For example, if we
// want to filter out type `speed`, we have to use `fan` flag instead. Refer
// to #192 for more information. All the tests below this point will be
// associated with this issue.
func TestReadCommandRequestRackLevelMultipleFlagsSuccessYaml(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/scan", 200, scanRespOK)
	test.Serve(t, mux, "/synse/v2/info/rack-1", 200, infoBoardRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-1", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-2", 200, fanReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-3", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-4", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-5", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-6", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-7", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-8", 200, temperatureReadRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "yaml",
		ServerCommand.Name,
		readCommand.Name,
		"rack-1",
		"--type", "fan",
		"--type", "led",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "read.success.yaml.multiple_flags.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCommandRequestRackLevelMultipleFlagsSuccessJson tests the 'read'
// command using multiple 'type' flags at rack-level when it gets a 200 response
// from Synse Server, with JSON output.
func TestReadCommandRequestRackLevelMultipleFlagsSuccessJson(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/scan", 200, scanRespOK)
	test.Serve(t, mux, "/synse/v2/info/rack-1", 200, infoBoardRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-1", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-2", 200, fanReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-3", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-4", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-5", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-6", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-7", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-8", 200, temperatureReadRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "json",
		ServerCommand.Name,
		readCommand.Name,
		"rack-1",
		"--type", "fan",
		"--type", "led",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "read.success.json.multiple_flags.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCommandRequestRackLevelMultipleFlagsSuccessPretty tests the 'read'
// command using multiple 'type' flags at rack-level when it gets a 200 response
// from Synse Server, with pretty output.
func TestReadCommandRequestMultipleFlagsSuccessPretty(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/scan", 200, scanRespOK)
	test.Serve(t, mux, "/synse/v2/info/rack-1", 200, infoBoardRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-1", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-2", 200, fanReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-3", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-4", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-5", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-6", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-7", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-8", 200, temperatureReadRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "pretty",
		ServerCommand.Name,
		readCommand.Name,
		"rack-1",
		"--type", "fan",
		"--type", "led",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "read.success.pretty.multiple_flags.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCommandRequestBoardLevelMultipleFlagsSuccessYaml tests the 'read'
// command using multiple 'type' flags at board-level when it gets a 200 response
// from Synse Server, with YAML output.
func TestReadCommandRequestBoardLevelMultipleFlagsSuccessYaml(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/scan", 200, scanRespOK)
	test.Serve(t, mux, "/synse/v2/info/rack-1/board-1", 200, infoBoardRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-1", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-2", 200, fanReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-3", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-4", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-5", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-6", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-7", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-8", 200, temperatureReadRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "yaml",
		ServerCommand.Name,
		readCommand.Name,
		"rack-1", "board-1",
		"--type", "fan",
		"--type", "led",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "read.success.yaml.multiple_flags.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCommandRequestBoardLevelMultipleFlagsSuccessJson tests the 'read'
// command using multiple 'type' flags at board-level when it gets a 200 response
// from Synse Server, with JSON output.
func TestReadCommandRequestBoardLevelMultipleFlagsSuccessJson(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/scan", 200, scanRespOK)
	test.Serve(t, mux, "/synse/v2/info/rack-1/board-1", 200, infoBoardRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-1", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-2", 200, fanReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-3", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-4", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-5", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-6", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-7", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-8", 200, temperatureReadRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "json",
		ServerCommand.Name,
		readCommand.Name,
		"rack-1", "board-1",
		"--type", "fan",
		"--type", "led",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "read.success.json.multiple_flags.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCommandRequestBoardLevelMultipleFlagsSuccessPretty tests the 'read'
// command using multiple 'type' flags at board-level when it gets a 200 response
// from Synse Server, with pretty output.
func TestReadCommandRequestBoardLevelMultipleFlagsSuccessPretty(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/scan", 200, scanRespOK)
	test.Serve(t, mux, "/synse/v2/info/rack-1/board-1", 200, infoBoardRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-1", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-2", 200, fanReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-3", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-4", 200, ledReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-5", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-6", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-7", 200, temperatureReadRespOK)
	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-8", 200, temperatureReadRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "pretty",
		ServerCommand.Name,
		readCommand.Name,
		"rack-1", "board-1",
		"--type", "fan",
		"--type", "led",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "read.success.pretty.multiple_flags.golden"))
	test.ExpectNoError(t, err)
}
