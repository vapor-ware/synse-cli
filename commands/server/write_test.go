package server

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/vapor-ware/synse-cli/internal/test"
)

const writeRespOK = `
[
  {
    "context":{
      "action":"color",
      "raw":[
        "000000"
      ]
    },
    "transaction":"b9u6ut6q5i6g020lau70"
  }
]`

func TestWriteCommandError(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, WriteCommand)

	err := app.Run([]string{app.Name, WriteCommand.Name, "rack-1", "board-1", "device-1"})

	test.ExpectExitCoderError(t, err)
}

func TestWriteCommandError2(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, WriteCommand)

	err := app.Run([]string{app.Name, WriteCommand.Name, "rack-1", "board-1", "device-1", "color", "000000", "extra"})

	test.ExpectExitCoderError(t, err)
}

func TestWriteCommandError3(t *testing.T) {
	test.Setup()

	mux, server := test.TestServer()
	defer server.Close()
	mux.HandleFunc(
		"/synse/2.0/write/rack-1/board-1/device-1",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
		},
	)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, WriteCommand)

	err := app.Run([]string{app.Name, WriteCommand.Name, "rack-1", "board-1", "device-1", "color", "000000"})

	test.ExpectExitCoderError(t, err)
}

func TestWriteCommandSuccess(t *testing.T) {
	test.Setup()

	mux, server := test.TestServer()
	defer server.Close()
	mux.HandleFunc(
		"/synse/2.0/write/rack-1/board-1/device-1",
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				t.Errorf("expected POST request, but was %v", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, writeRespOK)
		},
	)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, WriteCommand)

	err := app.Run([]string{app.Name, WriteCommand.Name, "rack-1", "board-1", "device-1", "color", "000000"})

	test.ExpectNoError(t, err)
}
