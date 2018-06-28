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
	// the mocked 200 OK JSON response for the Synse Server 'status' route
	statusRespOK = `
{
  "status": "ok",
  "timestamp": "2018-06-28T12:59:47.625842798Z"
}`

	// the mocked 500 error JSON response for the Synse Server 'status' route
	statusRespErr = `
{
  "http_code":500,
  "error_id":0,
  "description":"unknown",
  "timestamp":"2018-03-14 15:34:42.243715",
  "context":"test error."
}`
)

// TestStatusCommandError tests the 'status' command when it is unable to
// connect to the Synse Server instance because the active host is nil.
func TestStatusCommandError(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		ServerCommand.Name,
		statusCommand.Name,
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "error.nil.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestStatusCommandError2 tests the 'status' command when it is unable to
// connect to the Synse Server instance because the active host is not a
// Synse Server instance.
func TestStatusCommandError2(t *testing.T) {
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
		statusCommand.Name,
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	// FIXME: this test fails on CI because the expected output is different
	//     -Get http://localhost:5151/synse/version: dial tcp [::1]:5151: getsockopt: connection refused
	//     +Get http://localhost:5151/synse/version: dial tcp 127.0.0.1:5151: connect: connection refused
	//assert.Assert(t, golden.String(app.ErrBuffer.String(), "status.error.bad_host.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestStatusCommandRequestError tests the 'status' command when it gets a
// 500 response from Synse Server.
func TestStatusCommandRequestError(t *testing.T) {
	test.Setup()

	mux, server := test.UnversionedServer()
	defer server.Close()
	mux.HandleFunc("/synse/test", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		fmt.Fprint(w, statusRespErr)
	})

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		ServerCommand.Name,
		statusCommand.Name,
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "error.500.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestStatusCommandRequestErrorPretty tests the 'status' command when it gets
// a 200 response from Synse Server, with pretty output.
func TestStatusCommandRequestErrorPretty(t *testing.T) {
	test.Setup()

	mux, server := test.UnversionedServer()
	defer server.Close()
	mux.HandleFunc("/synse/test", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, statusRespOK)
	})

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "pretty",
		ServerCommand.Name,
		statusCommand.Name,
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "status.error.pretty.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestStatusCommandRequestSuccessYaml tests the 'status' command when it gets
// a 200 response from Synse Server, with YAML output.
func TestStatusCommandRequestSuccessYaml(t *testing.T) {
	test.Setup()

	mux, server := test.UnversionedServer()
	defer server.Close()
	mux.HandleFunc("/synse/test", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, statusRespOK)
	})

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "yaml",
		ServerCommand.Name,
		statusCommand.Name,
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "status.success.yaml.golden"))
	test.ExpectNoError(t, err)
}

// TestStatusCommandRequestSuccessJson tests the 'status' command when it gets
// a 200 response from Synse Server, with JSON output.
func TestStatusCommandRequestSuccessJson(t *testing.T) {
	test.Setup()

	mux, server := test.UnversionedServer()
	defer server.Close()
	mux.HandleFunc("/synse/test", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, statusRespOK)
	})

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "json",
		ServerCommand.Name,
		statusCommand.Name,
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "status.success.json.golden"))
	test.ExpectNoError(t, err)
}
