package server

import (
	"testing"

	"github.com/gotestyourself/gotestyourself/assert"
	"github.com/gotestyourself/gotestyourself/golden"

	"github.com/vapor-ware/synse-cli/internal/test"
)

const (
	// the mocked 200 OK JSON response for the 'plugins health' command.
	pluginsHealthRespOK = `
[
  {
    "tag":"vaporio\/emulator-plugin",
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

// TestPluginsHealthCommandRequestSuccessYaml tests the 'plugins health'
// command when it gets a 200 response from Synse Server, with YAML output.
func TestPluginsHealthCommandRequestSuccessYaml(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/2.0/plugins", 200, pluginsHealthRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "yaml",
		ServerCommand.Name,
		pluginsCommand.Name,
		pluginsHealthCommand.Name,
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "pluginshealth.success.yaml.golden"))
	test.ExpectNoError(t, err)
}

// TestPluginsHealthCommandRequestSuccessJson tests the 'plugins health'
// command when it gets a 200 response from Synse Server, with JSON output.
func TestPluginsHealthCommandRequestSuccessJson(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/2.0/plugins", 200, pluginsHealthRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "json",
		ServerCommand.Name,
		pluginsCommand.Name,
		pluginsHealthCommand.Name,
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "pluginshealth.success.json.golden"))
	test.ExpectNoError(t, err)
}

// TestPluginsHealthCommandSingleArgsRequestSuccessYaml tests the 'plugins health'
// command using a single argument when it gets a 200 response from Synse Server,
// with YAML output.
func TestPluginsHealthCommandSingleArgsRequestSuccessYaml(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/2.0/plugins", 200, pluginsHealthRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "yaml",
		ServerCommand.Name,
		pluginsCommand.Name,
		pluginsHealthCommand.Name,
		"vaporio/emulator-plugin",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "pluginshealth.success.yaml.single_args.golden"))
	test.ExpectNoError(t, err)
}

// TestPluginsHealthCommandMultipleArgsRequestSuccessYaml tests the 'plugins health'
// command using multiple arguments when it gets a 200 response from Synse Server,
// with YAML output.
func TestPluginsHealthCommandMultipleArgsRequestSuccessYaml(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/2.0/plugins", 200, pluginsHealthRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "yaml",
		ServerCommand.Name,
		pluginsCommand.Name,
		pluginsHealthCommand.Name,
		"vaporio/emulator-plugin",
		"vaporio/unix-plugin",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "pluginshealth.success.yaml.multiple_args.golden"))
	test.ExpectNoError(t, err)
}

// TestPluginsHealthCommandSingleArgsRequestSuccessJson tests the 'plugins health'
// command using a single argument when it gets a 200 response from Synse Server,
// with JSON output.
func TestPluginsHealthCommandSingleArgsRequestSuccessJson(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/2.0/plugins", 200, pluginsHealthRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "json",
		ServerCommand.Name,
		pluginsCommand.Name,
		pluginsHealthCommand.Name,
		"vaporio/unix-plugin",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "pluginshealth.success.json.single_args.golden"))
	test.ExpectNoError(t, err)
}

// TestPluginsHealthCommandMultipleArgsRequestSuccessJson tests the 'plugins health'
// command using multiple plugin tags arguments when it gets a 200 response from
// Synse Server, with JSON output.
func TestPluginsHealthCommandMultipleArgsRequestSuccessJson(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/2.0/plugins", 200, pluginsHealthRespOK)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "json",
		ServerCommand.Name,
		pluginsCommand.Name,
		pluginsHealthCommand.Name,
		"vaporio/emulator-plugin",
		"vaporio/unix-plugin",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "pluginshealth.success.json.multiple_args.golden"))
	test.ExpectNoError(t, err)
}
