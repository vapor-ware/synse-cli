package server

import (
	"testing"

	"github.com/gotestyourself/gotestyourself/assert"
	"github.com/gotestyourself/gotestyourself/golden"

	"github.com/vapor-ware/synse-cli/internal/test"
	"github.com/vapor-ware/synse-cli/pkg/config"
)

const (
	// the mocked 200 OK JSON response for the Synse Server 'readcached' route
	readCachedRespOk = `
{
  "location":{
    "rack":"rack-1",
    "board":"board-1",
    "device":"device-1"
  },
  "kind":"temperature",
  "value":"65",
  "timestamp":"2018-11-01T12:41:50.333443322Z",
  "unit":{
    "symbol":"C",
    "name":"celsius"
  },
  "type":"temperature",
  "info":"mock temperature response"
}
{
  "location":{
    "rack":"rack-1",
    "board":"board-1",
    "device":"device-1"
  },
  "kind":"temperature",
  "value":"66",
  "timestamp":"2018-11-11T12:41:50.333443322Z",
  "unit":{
    "symbol":"C",
    "name":"celsius"
  },
  "type":"temperature",
  "info":"mock temperature response"
}`

	// the mocked 500 error JSON response for the Synse Server 'readcached' route
	readCachedRespErr = `
{
  "http_code":500,
  "error_id":0,
  "description":"unknown",
  "timestamp":"2018-03-14 15:34:42.243715",
  "context":"test error."
}`
)

// TestReadCachedCommandError tests the 'readcached' command when it is unable to
// connect to the Synse Server instance because the active host is nil.
func TestReadCachedCommandError(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		ServerCommand.Name,
		readCachedCommand.Name,
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "error.nil.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestReadCachedCommandError2 tests the 'readcached' command when it is unable to
// connect to the Synse Server instance because the active host is not a
// Synse Server instance.
func TestReadCachedCommandError2(t *testing.T) {
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
		readCachedCommand.Name,
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	// FIXME: this test fails on CI because the expected output is different
	//     -Get http://localhost:5151/synse/version: dial tcp [::1]:5151: getsockopt: connection refused
	//     +Get http://localhost:5151/synse/version: dial tcp 127.0.0.1:5151: connect: connection refused
	//assert.Assert(t, golden.String(app.ErrBuffer.String(), "error.bad_host.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestReadCachedCommandRequestError tests the 'readcached' command when it gets a
// 500 response from Synse Server.
func TestReadCachedCommandRequestError(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/readcached", 500, readCachedRespErr)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		ServerCommand.Name,
		readCachedCommand.Name,
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.ErrBuffer.String(), "error.500.golden"))
	test.ExpectExitCoderError(t, err)
}

// TestReadCachedCommandRequestSuccessYaml tests the 'readcached' command when it gets
// a 200 response from Synse Server, with YAML output.
func TestReadCachedCommandRequestSuccessYaml(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/readcached", 200, readCachedRespOk)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "yaml",
		ServerCommand.Name,
		readCachedCommand.Name,
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "readcached.success.yaml.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCachedCommandRequestSuccessSingleParamsYaml tests the 'readcached' command when it gets
// a 200 response from Synse Server, with YAML output, using a single query parameter.
func TestReadCachedCommandRequestSuccessSingleParamsYaml(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/readcached", 200, readCachedRespOk)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "yaml",
		ServerCommand.Name,
		readCachedCommand.Name,
		"--start", "2018-11-01T12:41:50.333443322Z",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "readcached.success.yaml.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCachedCommandRequestSuccessMultipleParamsYaml tests the 'readcached' command when it gets
// a 200 response from Synse Server, with YAML output, using multiple query parameters.
func TestReadCachedCommandRequestSuccessMultipleParamsYaml(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/readcached", 200, readCachedRespOk)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "yaml",
		ServerCommand.Name,
		readCachedCommand.Name,
		"--start", "2018-11-01T12:41:50.333443322Z",
		"--end", "2018-11-11T12:41:50.333443322Z",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "readcached.success.yaml.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCachedCommandRequestSuccessInvalidParamsYaml tests the 'readcached' command when it gets
// a 200 response from Synse Server, with YAML output, using an invalid query parameters.
func TestReadCachedCommandRequestSuccessInvalidParamsYaml(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/readcached", 200, readCachedRespOk)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "yaml",
		ServerCommand.Name,
		readCachedCommand.Name,
		"--start", "invalid",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "readcached.success.yaml.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCachedCommandRequestSuccessExtraArgsYaml tests the 'readcached' command when
// too many arguments are provided, with YAML output,
func TestReadCachedCommandRequestSuccessExtraArgsYaml(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/readcached", 200, readCachedRespOk)

	test.AddServerHost(server)

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "yaml",
		ServerCommand.Name,
		readCachedCommand.Name,
		"extra",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "readcached.success.yaml.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCachedCommandRequestSuccessJson tests the 'readcached' command when it gets
// a 200 response from Synse Server, with JSON output.
func TestReadCachedCommandRequestSuccessJson(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/readcached", 200, readCachedRespOk)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "json",
		ServerCommand.Name,
		readCachedCommand.Name,
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "readcached.success.json.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCachedCommandRequestSuccessSingleParamsJson tests the 'readcached' command when it gets
// a 200 response from Synse Server, with JSON output, using a single query parameter.
func TestReadCachedCommandRequestSuccessSingleParamsJson(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/readcached", 200, readCachedRespOk)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "json",
		ServerCommand.Name,
		readCachedCommand.Name,
		"--start", "2018-11-01T12:41:50.333443322Z",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "readcached.success.json.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCachedCommandRequestSuccessMultipleParamsJson tests the 'readcached' command when it gets
// a 200 response from Synse Server, with JSON output, using multiple query parameters.
func TestReadCachedCommandRequestSuccessMultipleParamsJson(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/readcached", 200, readCachedRespOk)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "json",
		ServerCommand.Name,
		readCachedCommand.Name,
		"--start", "2018-11-01T12:41:50.333443322Z",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "readcached.success.json.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCachedCommandRequestSuccessInvalidParamsJson tests the 'readcached' command when it gets
// a 200 response from Synse Server, with JSON output, using an invalid query parameters.
func TestReadCachedCommandRequestSuccessInvalidParamsJson(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/readcached", 200, readCachedRespOk)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "json",
		ServerCommand.Name,
		readCachedCommand.Name,
		"--start", "invalid",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "readcached.success.json.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCachedCommandRequestSuccessExtraArgsJson tests the 'readcached' command when
// too many arguments are provided, with JSON output,
func TestReadCachedCommandRequestSuccessExtraArgsJson(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/readcached", 200, readCachedRespOk)

	test.AddServerHost(server)

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "json",
		ServerCommand.Name,
		readCachedCommand.Name,
		"extra",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "readcached.success.json.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCachedCommandRequestSuccessPretty tests the 'readcached' command when it gets
// a 200 response from Synse Server, with pretty output.
func TestReadCachedCommandRequestSuccessPretty(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/readcached", 200, readCachedRespOk)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "pretty",
		ServerCommand.Name,
		readCachedCommand.Name,
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "readcached.success.pretty.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCachedCommandRequestSuccessSingpleParamsPretty tests the 'readcached' command when it gets
// a 200 response from Synse Server, with pretty output, using a single query parameter.
func TestReadCachedCommandRequestSuccessSingleParamsPretty(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/readcached", 200, readCachedRespOk)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "pretty",
		ServerCommand.Name,
		readCachedCommand.Name,
		"--start", "2018-11-01T12:41:50.333443322Z",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "readcached.success.pretty.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCachedCommandRequestSuccessMultipleParamsPretty tests the 'readcached' command when it gets
// a 200 response from Synse Server, with pretty output, using multiple query parameters.
func TestReadCachedCommandRequestSuccessMultipleParamsPretty(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/readcached", 200, readCachedRespOk)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "pretty",
		ServerCommand.Name,
		readCachedCommand.Name,
		"--start", "2018-11-01T12:41:50.333443322Z",
		"--end", "2018-11-11T12:41:50.333443322Z",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "readcached.success.pretty.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCachedCommandRequestSuccessInvalidParamsPretty tests the 'readcached' command when it gets
// a 200 response from Synse Server, with pretty output, using an invalid query parameters.
func TestReadCachedCommandRequestSuccessInvalidParamsPretty(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/readcached", 200, readCachedRespOk)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "pretty",
		ServerCommand.Name,
		readCachedCommand.Name,
		"--start", "invalid",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "readcached.success.pretty.golden"))
	test.ExpectNoError(t, err)
}

// TestReadCachedCommandRequestSuccessExtraArgsPretty tests the 'readcached' command when
// too many arguments are provided, with pretty output,
func TestReadCachedCommandRequestSuccessExtraArgsPretty(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/readcached", 200, readCachedRespOk)

	test.AddServerHost(server)

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ServerCommand)

	err := app.Run([]string{
		app.Name,
		"--format", "pretty",
		ServerCommand.Name,
		readCachedCommand.Name,
		"extra",
	})

	t.Logf("Standard Out: \n%s", app.OutBuffer.String())
	t.Logf("Standard Error: \n%s", app.ErrBuffer.String())

	assert.Assert(t, golden.String(app.OutBuffer.String(), "readcached.success.pretty.golden"))
	test.ExpectNoError(t, err)
}
