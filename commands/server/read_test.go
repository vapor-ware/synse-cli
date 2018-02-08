package server

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/vapor-ware/synse-cli/internal/test"
)

const readRespOK = `
{
  "type":"temperature",
  "data":{
    "temperature":{
      "value":51.0,
      "timestamp":"2018-02-08 15:54:26.255253838 +0000 UTC m=+2841.145248278",
      "unit":{
        "symbol":"C",
        "name":"degrees celsius"
      }
    }
  }
}`

func TestReadCommandError(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ReadCommand)

	err := app.Run([]string{app.Name, ReadCommand.Name, "rack-1", "board-1"})

	test.ExpectExitCoderError(t, err)
}

func TestReadCommandError2(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ReadCommand)

	err := app.Run([]string{app.Name, ReadCommand.Name, "rack-1", "board-1", "device-1", "extra"})

	test.ExpectExitCoderError(t, err)
}

func TestReadCommandError3(t *testing.T) {
	test.Setup()

	mux, server := test.TestServer()
	defer server.Close()
	mux.HandleFunc(
		"/synse/2.0/read/rack-1/board-1/device-1",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
		},
	)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ReadCommand)

	err := app.Run([]string{app.Name, ReadCommand.Name, "rack-1", "board-1", "device-1"})

	test.ExpectExitCoderError(t, err)
}

func TestReadCommandSuccess(t *testing.T) {
	test.Setup()

	mux, server := test.TestServer()
	defer server.Close()
	mux.HandleFunc(
		"/synse/2.0/read/rack-1/board-1/device-1",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, readRespOK)
		},
	)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, ReadCommand)

	err := app.Run([]string{app.Name, ReadCommand.Name, "rack-1", "board-1", "device-1"})

	test.ExpectNoError(t, err)
}
