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
          "id":"vec",
          "devices":[
            {
              "id":"f97f284037b04badb6bb7aacd9654a4e",
              "info":"Synse Temperature Sensor 5",
              "type":"temperature"
            },
            {
              "id":"eb9a56f95b5bd6d9b51996ccd0f2329c",
              "info":"Synse Fan",
              "type":"fan"
            },
            {
              "id":"f52d29fecf05a195af13f14c7306cfed",
              "info":"Synse LED",
              "type":"led"
            },
            {
              "id":"d29e0bd113a484dc48fd55bd3abad6bb",
              "info":"Synse backup LED",
              "type":"led"
            },
            {
              "id":"eb100067acb0c054cf877759db376b03",
              "info":"Synse Temperature Sensor 1",
              "type":"temperature"
            },
            {
              "id":"83cc1efe7e596e4ab6769e0c6e3edf88",
              "info":"Synse Temperature Sensor 2",
              "type":"temperature"
            },
            {
              "id":"db1e5deb43d9d0af6d80885e74362913",
              "info":"Synse Temperature Sensor 3",
              "type":"temperature"
            },
            {
              "id":"329a91c6781ce92370a3c38ba9bf35b2",
              "info":"Synse Temperature Sensor 4",
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

	mux, server := test.TestServer()
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
	app.Commands = append(app.Commands, ScanCommand)

	err := app.Run([]string{app.Name, ScanCommand.Name})

	test.ExpectExitCoderError(t, err)
}

func TestScanCommandSuccess(t *testing.T) {
	test.Setup()

	mux, server := test.TestServer()
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
	app.Commands = append(app.Commands, ScanCommand)

	err := app.Run([]string{app.Name, ScanCommand.Name})

	test.ExpectNoError(t, err)
}
