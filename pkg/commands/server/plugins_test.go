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
	// the mocked 200 OK JSON response for the Synse Server 'plugins' route
	pluginsRespOK = `
[
  {
    "name": "foo",
    "network": "tcp",
    "address": "localhost:6000"
  },
  {
    "name": "bar",
    "network": "unix",
    "address": "/tmp/synse/proc/bar.sock"
  }
]`

	// the mocked 500 error JSON response for the Synse Server 'plugins' route
	pluginsRespErr = `
{
  "http_code":500,
  "error_id":0,
  "description":"unknown",
  "timestamp":"2018-03-14 15:34:42.243715",
  "context":"test error."
}`
)

// TestPluginsCommandError tests the 'plugins' command when it is unable to
// connect to the Synse Server instance because the active host is nil.
func TestPluginsCommandError(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		ServerCommand.Name,
		pluginsCommand.Name,
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "error.nil.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestPluginsCommandError2 tests the 'plugins' command when it is unable to
// connect to the Synse Server instance because the active host is not a
// Synse Server instance.
func TestPluginsCommandError2(t *testing.T) {
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
		pluginsCommand.Name,
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "error.bad_host.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestPluginsCommandRequestError tests the 'plugins' command when it gets a
// 500 response from Synse Server.
func TestPluginsCommandRequestError(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()
	mux.HandleFunc(
		"/synse/2.0/plugins",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
			fmt.Fprint(w, pluginsRespErr)
		},
	)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		ServerCommand.Name,
		pluginsCommand.Name,
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "error.500.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestPluginsCommandRequestErrorYaml tests the 'plugins' command when it gets
// a 200 response from Synse Server, with YAML output.
func TestPluginsCommandRequestErrorYaml(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()
	mux.HandleFunc(
		"/synse/2.0/plugins",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, pluginsRespOK)
		},
	)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "yaml",
		ServerCommand.Name,
		pluginsCommand.Name,
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "plugins.error.yaml.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestPluginsCommandRequestErrorJson tests the 'plugins' command when it gets
// a 200 response from Synse Server, with JSON output.
func TestPluginsCommandRequestErrorJson(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()
	mux.HandleFunc(
		"/synse/2.0/plugins",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, pluginsRespOK)
		},
	)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "json",
		ServerCommand.Name,
		pluginsCommand.Name,
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "plugins.error.json.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestPluginsCommandRequestSuccessPretty tests the 'plugins' command when it gets
// a 200 response from Synse Server, with pretty output.
func TestPluginsCommandRequestSuccessPretty(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()
	mux.HandleFunc(
		"/synse/2.0/plugins",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, pluginsRespOK)
		},
	)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "pretty",
		ServerCommand.Name,
		pluginsCommand.Name,
	})

	assert.Assert(t, golden.String(app.OutBuffer.String(), "plugins.success.pretty.golden"))
	test.ExpectNoError(t, err)
}
