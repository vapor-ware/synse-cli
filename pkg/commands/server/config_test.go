package server

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/vapor-ware/synse-cli/internal/test"
)

const configRespOK = `
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
  }
}`

func TestConfigCommandError(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()
	mux.HandleFunc("/synse/2.0/config", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
	})

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, configCommand)

	err := app.Run([]string{app.Name, configCommand.Name})

	test.ExpectExitCoderError(t, err)
}

func TestConfigCommandSuccess(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()
	mux.HandleFunc("/synse/2.0/config", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, configRespOK)
	})

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, configCommand)

	err := app.Run([]string{app.Name, configCommand.Name})

	test.ExpectNoError(t, err)
}
