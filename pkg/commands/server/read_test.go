package server

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gotestyourself/gotestyourself/assert"
	"github.com/gotestyourself/gotestyourself/golden"

	"github.com/vapor-ware/synse-cli/internal/test"
	"github.com/vapor-ware/synse-cli/pkg/config"
)

const (
	// the mocked 200 OK JSON response for the Synse Server 'read' route
	readRespOK = `
{
  "type":"temperature",
  "data":{
    "temperature":{
      "value":51.0,
      "timestamp":"2018-02-08 15:54:26.255253838 +0000 UTC m=+2841.145248278",
      "unit":{
        "symbol":"C",
        "name":"degrees celsius"
      }
    }
  }
}`

	// the mocked 500 error JSON response for the Synse Server 'read' route
	readRespErr = `
{
  "http_code":500,
  "error_id":0,
  "description":"unknown",
  "timestamp":"2018-03-14 15:34:42.243715",
  "context":"test error."
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

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "error.bad_host.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestReadCommandError3 tests the 'read' command when no arguments
// are provided, but some are required.
func TestReadCommandError3(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		ServerCommand.Name,
		readCommand.Name,
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "read.error.no_args.golden"))
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

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "read.error.extra_args.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestReadCommandRequestError tests the 'read' command when it gets a
// 500 response from Synse Server.
func TestReadCommandRequestError(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()
	mux.HandleFunc(
		"/synse/2.0/read/rack-1/board-1/device-1",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
			fmt.Fprint(w, readRespErr)
		},
	)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		ServerCommand.Name,
		readCommand.Name,
		"rack-1", "board-1", "device-1",
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "error.500.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestReadCommandRequestErrorYaml tests the 'read' command when it gets
// a 200 response from Synse Server, with YAML output.
func TestReadCommandRequestErrorYaml(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()
	mux.HandleFunc(
		"/synse/2.0/read/rack-1/board-1/device-1",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, readRespOK)
		},
	)

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

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "read.error.yaml.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestReadCommandRequestErrorJson tests the 'read' command when it gets
// a 200 response from Synse Server, with JSON output.
func TestReadCommandRequestErrorJson(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()
	mux.HandleFunc(
		"/synse/2.0/read/rack-1/board-1/device-1",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, readRespOK)
		},
	)

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

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "read.error.json.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestReadCommandRequestSuccessPretty tests the 'read' command when it gets
// a 200 response from Synse Server, with pretty output.
func TestReadCommandRequestSuccessPretty(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()
	mux.HandleFunc(
		"/synse/2.0/read/rack-1/board-1/device-1",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, readRespOK)
		},
	)

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

	assert.Assert(t, golden.String(app.OutBuffer.String(), "read.success.pretty.golden"))
	test.ExpectNoError(t, err)
}
