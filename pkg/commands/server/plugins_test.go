package server

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/vapor-ware/synse-cli/internal/test"
)

const pluginsRespOK = `
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

func TestPluginsCommandError(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()
	mux.HandleFunc(
		"/synse/2.0/plugins",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
		},
	)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, pluginsCommand)

	err := app.Run([]string{app.Name, pluginsCommand.Name})

	test.ExpectExitCoderError(t, err)
}

func TestPluginsCommandSuccess(t *testing.T) {
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
	app.Commands = append(app.Commands, pluginsCommand)

	err := app.Run([]string{app.Name, pluginsCommand.Name})

	test.ExpectNoError(t, err)
}
