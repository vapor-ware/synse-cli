package server

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gotestyourself/gotestyourself/assert"
	"github.com/gotestyourself/gotestyourself/golden"

	"github.com/vapor-ware/synse-cli/internal/test"
)

const (
	// the mocked 200 OK JSON response for the 'plugins info' command.
	pluginsInfoRespOK = `
[
  {
    "tag":"vaporio\/emulator-plugin",
    "name":"sample-tcp",
    "description":"A sample emulator plugin",
    "maintainer":"vaporio",
    "vcs":"github.com\/vapor-ware\/synse-emulator-plugin",
    "version":{
      "plugin_version":"2.0.0",
      "sdk_version":"1.0.0",
      "build_date":"2018-06-25T14:39:18",
      "git_commit":"4831f12",
      "git_tag":"1.0.2-8-g4831f12",
      "arch":"amd64",
      "os":"linux"
    },
    "network":{
      "protocol":"tcp",
      "address":"emulator-plugin:5001"
    },
    "health":{
      "timestamp":"2018-06-27T18:30:46.237254715Z",
      "status":"ok",
      "message":"",
      "checks":[
        {
          "name":"read buffer health",
          "status":"ok",
          "message":"",
          "timestamp":"2018-06-27T18:30:16.531781924Z",
          "type":"periodic"
        },
        {
          "name":"write buffer health",
          "status":"ok",
          "message":"",
          "timestamp":"2018-06-27T18:30:16.531781924Z",
          "type":"periodic"
        }
      ]
    }
  },
  {
    "tag":"vaporio\/unix-plugin",
    "name":"sample-unix",
    "description":"A sample unix plugin",
    "maintainer":"vaporio",
    "vcs":"github.com\/vapor-ware\/synse-unix-plugin",
    "version":{
      "plugin_version":"2.0.0",
      "sdk_version":"1.0.0",
      "build_date":"2018-06-25T14:39:18",
      "git_commit":"4831f12",
      "git_tag":"1.0.2-8-g4831f12",
      "arch":"amd64",
      "os":"linux"
    },
    "network":{
      "protocol":"unix",
      "address":"/tmp/synse/proc/bar.sock"
    },
    "health":{
      "timestamp":"2018-06-27T18:30:46.237254715Z",
      "status":"ok",
      "message":"",
      "checks":[
        {
          "name":"read buffer health",
          "status":"ok",
          "message":"",
          "timestamp":"2018-06-27T18:30:16.531781924Z",
          "type":"periodic"
        },
        {
          "name":"write buffer health",
          "status":"ok",
          "message":"",
          "timestamp":"2018-06-27T18:30:16.531781924Z",
          "type":"periodic"
        }
      ]
    }
  }
]`

	// The 500 error JSON response for the Synse Server 'plugins' route
	// is already mocked in plugins_test.go.
)

// TestPluginsInfoCommandRequestSuccessYaml tests the 'plugins info' command when it gets
// a 200 response from Synse Server, with YAML output.
func TestPluginsInfoCommandRequestSuccessYaml(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()
	mux.HandleFunc(
		"/synse/2.0/plugins",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, pluginsInfoRespOK)
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
		pluginsInfoCommand.Name,
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "pluginsinfo.success.yaml.golden"))
	test.ExpectNoError(t, err)
}

// TestPluginsInfoCommandRequestSuccessJson tests the 'plugins info' command when it gets
// a 200 response from Synse Server, with JSON output.
func TestPluginsInfoCommandRequestSuccessJson(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()
	mux.HandleFunc(
		"/synse/2.0/plugins",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, pluginsInfoRespOK)
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
		pluginsInfoCommand.Name,
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "pluginsinfo.success.json.golden"))
	test.ExpectNoError(t, err)
}

// TestPluginsInfoCommandSingleArgsRequestSuccessYaml tests the 'plugins info'
// command using a single argument when it gets a 200 response
// from Synse Server, with YAML output.
func TestPluginsInfoCommandSingleArgsRequestSuccessYaml(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()
	mux.HandleFunc(
		"/synse/2.0/plugins",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, pluginsInfoRespOK)
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
		pluginsInfoCommand.Name,
		"vaporio/emulator-plugin",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "pluginsinfo.success.yaml.single_args.golden"))
	test.ExpectNoError(t, err)
}

// TestPluginsInfoCommandMultipleArgsRequestSuccessYaml tests the 'plugins info'
// command using a multiple arguments when it gets a 200 response
// from Synse Server, with YAML output.
func TestPluginsInfoCommandMultipleArgsRequestSuccessYaml(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()
	mux.HandleFunc(
		"/synse/2.0/plugins",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, pluginsInfoRespOK)
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
		pluginsInfoCommand.Name,
		"vaporio/emulator-plugin", "vaporio/unix-plugin",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "pluginsinfo.success.yaml.multiple_args.golden"))
	test.ExpectNoError(t, err)
}

// TestPluginsInfoCommandSingleArgsRequestSuccessJson tests the 'plugins info'
// command using a single argument  when it gets a 200 response
// from Synse Server, with JSON output.
func TestPluginsInfoCommandSingleArgsRequestSuccessJson(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()
	mux.HandleFunc(
		"/synse/2.0/plugins",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, pluginsInfoRespOK)
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
		pluginsInfoCommand.Name,
		"vaporio/unix-plugin",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "pluginsinfo.success.json.single_args.golden"))
	test.ExpectNoError(t, err)
}

// TestPluginsInfoCommandMultipleArgsRequestSuccessJson tests the 'plugins info'
// command using multiple arguments when it gets a 200 response
// from Synse Server, with JSON output.
func TestPluginsInfoCommandMultipleArgsRequestSuccessJson(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()
	mux.HandleFunc(
		"/synse/2.0/plugins",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, pluginsInfoRespOK)
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
		pluginsInfoCommand.Name,
		"vaporio/unix-plugin", "vaporio/emulator-plugin",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "pluginsinfo.success.json.multiple_args.golden"))
	test.ExpectNoError(t, err)
}

// TestPluginsInfoCommandExtraArgsRequestSuccessYaml tests the 'plugins info'
// command using a [PLUGIN TAG] argument when extra arguments are provided.
func TestPluginsInfoCommandExtraArgsRequestSuccessYaml(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "yaml",
		ServerCommand.Name,
		pluginsCommand.Name,
		pluginsInfoCommand.Name,
		"vaporio/emulator-plugin", "vaporio/unix-plugin", "foo", "bar",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "pluginsinfo.error.extra_args.golden"))
	test.ExpectExitCoderError(t, err)
}
