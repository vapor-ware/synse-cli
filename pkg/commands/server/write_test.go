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
	// the mocked 200 OK JSON response for the Synse Server 'write' route
	writeRespOK = `
[
  {
    "context":{
      "action":"color",
      "raw":[
        "000000"
      ]
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

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "error.bad_host.golden"))
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

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "write.error.extra_args.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestWriteCommandRequestError tests the 'write' command when it gets a
// 500 response from Synse Server.
func TestWriteCommandRequestError(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()
	mux.HandleFunc(
		"/synse/2.0/write/rack-1/board-1/device-1",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
			fmt.Fprint(w, writeRespErr)
		},
	)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		ServerCommand.Name,
		writeCommand.Name,
		"rack-1", "board-1", "device-1", "color", "000000",
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "error.500.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestWriteCommandRequestErrorYaml tests the 'write' command when it gets
// a 200 response from Synse Server, with YAML output.
func TestWriteCommandRequestErrorYaml(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()
	mux.HandleFunc(
		"/synse/2.0/write/rack-1/board-1/device-1",
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				t.Errorf("expected POST request, but was %v", r.Method)
			}
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, writeRespOK)
		},
	)

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

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "write.error.yaml.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestWriteCommandRequestErrorJson tests the 'write' command when it gets
// a 200 response from Synse Server, with JSON output.
func TestWriteCommandRequestErrorJson(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()
	mux.HandleFunc(
		"/synse/2.0/write/rack-1/board-1/device-1",
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				t.Errorf("expected POST request, but was %v", r.Method)
			}
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, writeRespOK)
		},
	)

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

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "write.error.json.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestWriteCommandRequestSuccessPretty tests the 'write' command when it gets
// a 200 response from Synse Server, with pretty output.
func TestWriteCommandRequestSuccessPretty(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()
	mux.HandleFunc(
		"/synse/2.0/write/rack-1/board-1/device-1",
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				t.Errorf("expected POST request, but was %v", r.Method)
			}
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, writeRespOK)
		},
	)

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

	assert.Assert(t, golden.String(app.OutBuffer.String(), "write.success.pretty.golden"))
	test.ExpectNoError(t, err)
}
