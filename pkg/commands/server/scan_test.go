package server

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/vapor-ware/synse-cli/internal/test"
)

const scanRespOK = `
{
  "racks":[
    {
      "id":"rack-1",
      "boards":[
        {
          "id":"board-1",
          "devices":[
            {
              "id":"device-1",
              "info":"Synse Temperature Sensor",
              "type":"temperature"
            },
            {
              "id":"device-2",
              "info":"Synse Fan",
              "type":"fan"
            },
            {
              "id":"device-3",
              "info":"Synse LED",
              "type":"led"
            },
            {
              "id":"device-4",
              "info":"Synse LED",
              "type":"led"
            },
            {
              "id":"device-5",
              "info":"Synse Temperature Sensor",
              "type":"temperature"
            },
            {
              "id":"device-6",
              "info":"Synse Temperature Sensor",
              "type":"temperature"
            },
            {
              "id":"device-7",
              "info":"Synse Temperature Sensor",
              "type":"temperature"
            },
            {
              "id":"device-8",
              "info":"Synse Temperature Sensor",
              "type":"temperature"
            }
          ]
        }
      ]
    }
  ]
}`

func TestScanCommandError(t *testing.T) {
	test.Setup()

	mux, server := test.Server()
	defer server.Close()
	mux.HandleFunc(
		"/synse/2.0/scan",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
		},
	)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, scanCommand)

	err := app.Run([]string{app.Name, scanCommand.Name})

	test.ExpectExitCoderError(t, err)
}

func TestScanCommandSuccess(t *testing.T) {
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

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, scanCommand)

	err := app.Run([]string{app.Name, scanCommand.Name})

	test.ExpectNoError(t, err)
}
