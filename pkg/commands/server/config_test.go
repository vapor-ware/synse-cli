package server

import (
	"net/http"
	"testing"

	"github.com/gotestyourself/gotestyourself/assert"
	"github.com/gotestyourself/gotestyourself/golden"

	"github.com/vapor-ware/synse-cli/internal/test"
	"github.com/vapor-ware/synse-cli/pkg/config"
)

const (
	// the mocked 200 OK JSON response for the Synse Server 'config' route
	configRespOK = `
{
  "locale":"en_US",
  "pretty_json":true,
  "logging":"debug",
  "cache":{
    "meta":{
      "ttl":20
    },
    "transaction":{
      "ttl":20
    }
  },
  "grpc":{
    "timeout":20
  },
  "plugin":{
    "tcp":[
      "emulator-plugin:5001"
    ],
    "unix":[],
    "discovery":{
      "kubernetes":{
        "endpoints":{
          "labels":{}
        }
      }
    }
  }
}`

	// the mocked 500 error JSON response for the Synse Server 'config' route
	configRespErr = `
{
  "http_code":500,
  "error_id":0,
  "description":"unknown",
  "timestamp":"2018-03-14 15:34:42.243715",
  "context":"test error."
}`
)

// TestConfigCommandError tests the 'config' command when it is unable to
// connect to the Synse Server instance because the active host is nil.
func TestConfigCommandError(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		ServerCommand.Name,
		configCommand.Name,
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "error.nil.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestConfigCommandError2 tests the 'config' command when it is unable to
// connect to the Synse Server instance because the active host is not a
// Synse Server instance.
func TestConfigCommandError2(t *testing.T) {
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
		configCommand.Name,
	})

	// FIXME: this test fails on CI because the expected output is different
	//     -Get http://localhost:5151/synse/version: dial tcp [::1]:5151: getsockopt: connection refused
	//     +Get http://localhost:5151/synse/version: dial tcp 127.0.0.1:5151: connect: connection refused
	//assert.Assert(t, golden.String(app.ErrBuffer.String(), "error.bad_host.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestConfigCommandRequestError tests the 'config' command when it gets a
// 500 response from Synse Server.
func TestConfigCommandRequestError(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()
	mux.HandleFunc("/synse/2.0/config", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		test.Fprint(t, w, configRespErr)
	})

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		ServerCommand.Name,
		configCommand.Name,
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "error.500.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestConfigCommandRequestErrorPretty tests the 'config' command when it gets
// a 200 response from Synse Server, with pretty output.
func TestConfigCommandRequestErrorPretty(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()
	mux.HandleFunc("/synse/2.0/config", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		test.Fprint(t, w, configRespOK)
	})

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "pretty",
		ServerCommand.Name,
		configCommand.Name,
	})

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "config.error.pretty.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestConfigCommandRequestSuccessYaml tests the 'config' command when it gets
// a 200 response from Synse Server, with YAML output.
func TestConfigCommandRequestSuccessYaml(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()
	mux.HandleFunc("/synse/2.0/config", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		test.Fprint(t, w, configRespOK)
	})

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "yaml",
		ServerCommand.Name,
		configCommand.Name,
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "config.success.yaml.golden"))
	test.ExpectNoError(t, err)
}

// TestConfigCommandRequestSuccessJson tests the 'config' command when it gets
// a 200 response from Synse Server, with JSON output.
func TestConfigCommandRequestSuccessJson(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()
	mux.HandleFunc("/synse/2.0/config", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		test.Fprint(t, w, configRespOK)
	})

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "json",
		ServerCommand.Name,
		configCommand.Name,
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "config.success.json.golden"))
	test.ExpectNoError(t, err)
}
