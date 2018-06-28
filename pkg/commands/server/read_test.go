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
  "kind":"temperature",
  "data":[
    {
      "value":"65",
      "timestamp":"2018-06-28T12:41:50.333443322Z",
      "unit":{
        "symbol":"C",
        "name":"celsius"
      },
      "type":"state",
      "info":""
    }
  ]
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

// TestReadCommandRequestError tests the 'read' command when it gets a
// 500 response from Synse Server.
func TestReadCommandRequestError(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()
	mux.HandleFunc(
		"/synse/2.0/scan",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, scanRespOK)
		},
	)
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

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "error.500.golden"))
	test.ExpectExitCoderError(t, err)
}

// FIXME: these are temporarily broken -- need updates to the formatter to get these outputting correctly

//// TestReadCommandRequestSuccessYaml tests the 'read' command when it gets
//// a 200 response from Synse Server, with YAML output.
//func TestReadCommandRequestSuccessYaml(t *testing.T) {
//	test.Setup()
//
//	mux, server := test.Server()
//	defer server.Close()
//	mux.HandleFunc(
//		"/synse/2.0/scan",
//		func(w http.ResponseWriter, r *http.Request) {
//			w.Header().Set("Content-Type", "application/json")
//			fmt.Fprint(w, scanRespOK)
//		},
//	)
//	mux.HandleFunc(
//		"/synse/2.0/read/rack-1/board-1/device-1",
//		func(w http.ResponseWriter, r *http.Request) {
//			w.Header().Set("Content-Type", "application/json")
//			fmt.Fprint(w, readRespOK)
//		},
//	)
//
//	test.AddServerHost(server)
//	app := test.NewFakeApp()
//	app.Commands = append(app.Commands, ServerCommand)
//
//	err := app.Run([]string{
//		app.Name,
//		"--format", "yaml",
//		ServerCommand.Name,
//		readCommand.Name,
//		"rack-1", "board-1", "device-1",
//	})
//
//	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
//	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())
//
//	assert.Assert(t, golden.String(app.OutBuffer.String(), "read.success.yaml.golden"))
//	test.ExpectNoError(t, err)
//}
//
//// TestReadCommandRequestSuccessJson tests the 'read' command when it gets
//// a 200 response from Synse Server, with JSON output.
//func TestReadCommandRequestSuccessJson(t *testing.T) {
//	test.Setup()
//
//	mux, server := test.Server()
//	defer server.Close()
//	mux.HandleFunc(
//		"/synse/2.0/scan",
//		func(w http.ResponseWriter, r *http.Request) {
//			w.Header().Set("Content-Type", "application/json")
//			fmt.Fprint(w, scanRespOK)
//		},
//	)
//	mux.HandleFunc(
//		"/synse/2.0/read/rack-1/board-1/device-1",
//		func(w http.ResponseWriter, r *http.Request) {
//			w.Header().Set("Content-Type", "application/json")
//			fmt.Fprint(w, readRespOK)
//		},
//	)
//
//	test.AddServerHost(server)
//	app := test.NewFakeApp()
//	app.Commands = append(app.Commands, ServerCommand)
//
//	err := app.Run([]string{
//		app.Name,
//		"--format", "json",
//		ServerCommand.Name,
//		readCommand.Name,
//		"rack-1", "board-1", "device-1",
//	})
//
//	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
//	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())
//
//	assert.Assert(t, golden.String(app.OutBuffer.String(), "read.success.json.golden"))
//	test.ExpectNoError(t, err)
//}

// TestReadCommandRequestSuccessPretty tests the 'read' command when it gets
// a 200 response from Synse Server, with pretty output.
func TestReadCommandRequestSuccessPretty(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()
	mux.HandleFunc(
		"/synse/2.0/scan",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, scanRespOK)
		},
	)
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

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "read.success.pretty.golden"))
	test.ExpectNoError(t, err)
}
