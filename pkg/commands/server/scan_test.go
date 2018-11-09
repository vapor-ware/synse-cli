package server

import (
	"testing"

	"github.com/gotestyourself/gotestyourself/assert"
	"github.com/gotestyourself/gotestyourself/golden"

	"github.com/vapor-ware/synse-cli/internal/test"
	"github.com/vapor-ware/synse-cli/pkg/config"
)

const (
	// the mocked 200 OK JSON response for the Synse Server 'scan' route
	scanRespOK = `
{
  "racks":[
    {
      "id":"rack-1",
      "boards":[
        {
          "id":"board-1",
          "devices":[
            {
              "id":"device-1",
              "info":"Synse Temperature Sensor",
              "type":"temperature"
            },
            {
              "id":"device-2",
              "info":"Synse Fan",
              "type":"fan"
            },
            {
              "id":"device-3",
              "info":"Synse LED",
              "type":"led"
            },
            {
              "id":"device-4",
              "info":"Synse LED",
              "type":"led"
            },
            {
              "id":"device-5",
              "info":"Synse Temperature Sensor",
              "type":"temperature"
            },
            {
              "id":"device-6",
              "info":"Synse Temperature Sensor",
              "type":"temperature"
            },
            {
              "id":"device-7",
              "info":"Synse Temperature Sensor",
              "type":"temperature"
            },
            {
              "id":"device-8",
              "info":"Synse Temperature Sensor",
              "type":"temperature"
            }
          ]
        }
      ]
    }
  ]
}`

	// the mocked 500 error JSON response for the Synse Server 'scan' route
	scanRespErr = `
{
  "http_code":500,
  "error_id":0,
  "description":"unknown",
  "timestamp":"2018-03-14 15:34:42.243715",
  "context":"test error."
}`
)

// TestScanCommandError tests the 'scan' command when it is unable to
// connect to the Synse Server instance because the active host is nil.
func TestScanCommandError(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		ServerCommand.Name,
		scanCommand.Name,
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "error.nil.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestScanCommandError2 tests the 'scan' command when it is unable to
// connect to the Synse Server instance because the active host is not a
// Synse Server instance.
func TestScanCommandError2(t *testing.T) {
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
		scanCommand.Name,
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	// FIXME: this test fails on CI because the expected output is different
	//     -Get http://localhost:5151/synse/version: dial tcp [::1]:5151: getsockopt: connection refused
	//     +Get http://localhost:5151/synse/version: dial tcp 127.0.0.1:5151: connect: connection refused
	//assert.Assert(t, golden.String(app.ErrBuffer.String(), "error.bad_host.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestScanCommandRequestError tests the 'scan' command when it gets a
// 500 response from Synse Server.
func TestScanCommandRequestError(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/scan", 500, scanRespErr)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		ServerCommand.Name,
		scanCommand.Name,
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "error.500.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestScanCommandRequestSuccessYaml tests the 'scan' command when it gets
// a 200 response from Synse Server, with YAML output.
func TestScanCommandRequestSuccessYaml(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/scan", 200, scanRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "yaml",
		ServerCommand.Name,
		scanCommand.Name,
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "scan.success.yaml.golden"))
	test.ExpectNoError(t, err)
}

// TestScanCommandRequestSuccessJson tests the 'scan' command when it gets
// a 200 response from Synse Server, with JSON output.
func TestScanCommandRequestSuccessJson(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/scan", 200, scanRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "json",
		ServerCommand.Name,
		scanCommand.Name,
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "scan.success.json.golden"))
	test.ExpectNoError(t, err)
}

// TestScanCommandRequestSuccessPretty tests the 'scan' command when it gets
// a 200 response from Synse Server, with pretty output.
func TestScanCommandRequestSuccessPretty(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/scan", 200, scanRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "pretty",
		ServerCommand.Name,
		scanCommand.Name,
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "scan.success.pretty.golden"))
	test.ExpectNoError(t, err)
}
