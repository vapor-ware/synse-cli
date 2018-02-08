package server

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/vapor-ware/synse-cli/internal/test"
)

const (
	infoRackRespOK = `
{
  "rack":"rack-1",
  "boards":[
    "vec"
  ]
}`

	infoBoardRespOK = `
{
  "board":"vec",
  "location":{
    "rack":"rack-1"
  },
  "devices":[
    "eb100067acb0c054cf877759db376b03",
    "83cc1efe7e596e4ab6769e0c6e3edf88",
    "db1e5deb43d9d0af6d80885e74362913",
    "329a91c6781ce92370a3c38ba9bf35b2",
    "f97f284037b04badb6bb7aacd9654a4e",
    "eb9a56f95b5bd6d9b51996ccd0f2329c",
    "f52d29fecf05a195af13f14c7306cfed",
    "d29e0bd113a484dc48fd55bd3abad6bb"
  ]
}`

	infoDeviceRespOK = `
{
  "timestamp":"2018-02-08 15:58:51.063845404 +0000 UTC m=+3105.953837345",
  "uid":"83cc1efe7e596e4ab6769e0c6e3edf88",
  "type":"temperature",
  "model":"emul8-temp",
  "manufacturer":"Vapor IO",
  "protocol":"emulator",
  "info":"Synse Temperature Sensor 2",
  "comment":"",
  "location":{
    "rack":"rack-1",
    "board":"vec"
  },
  "output":[
    {
      "type":"temperature",
      "data_type":"float",
      "precision":2,
      "unit":{
        "name":"degrees celsius",
        "symbol":"C"
      },
      "range":{
        "min":0,
        "max":100
      }
    }
  ]
}`

)


func TestInfoCommandError(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, InfoCommand)

	err := app.Run([]string{app.Name, InfoCommand.Name})

	test.ExpectExitCoderError(t, err)
}

func TestInfoCommandError2(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, InfoCommand)

	err := app.Run([]string{app.Name, InfoCommand.Name, "rack-1", "board-1", "device-1", "extra"})

	test.ExpectExitCoderError(t, err)
}

func TestInfoCommandErrorRack(t *testing.T) {
	test.Setup()

	mux, server := test.TestServer()
	defer server.Close()
	mux.HandleFunc(
		"/synse/2.0/info/rack-1",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
		},
	)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, InfoCommand)

	err := app.Run([]string{app.Name, InfoCommand.Name, "rack-1"})

	test.ExpectExitCoderError(t, err)
}

func TestInfoCommandSuccessRack(t *testing.T) {
	test.Setup()

	mux, server := test.TestServer()
	defer server.Close()
	mux.HandleFunc(
		"/synse/2.0/info/rack-1",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, infoRackRespOK)
		},
	)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, InfoCommand)

	err := app.Run([]string{app.Name, InfoCommand.Name, "rack-1"})

	test.ExpectNoError(t, err)
}

func TestInfoCommandErrorBoard(t *testing.T) {
	test.Setup()

	mux, server := test.TestServer()
	defer server.Close()
	mux.HandleFunc(
		"/synse/2.0/info/rack-1/board-1",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
		},
	)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, InfoCommand)

	err := app.Run([]string{app.Name, InfoCommand.Name, "rack-1", "board-1"})

	test.ExpectExitCoderError(t, err)
}

func TestInfoCommandSuccessBoard(t *testing.T) {
	test.Setup()

	mux, server := test.TestServer()
	defer server.Close()
	mux.HandleFunc(
		"/synse/2.0/info/rack-1/board-1",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, infoBoardRespOK)
		},
	)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, InfoCommand)

	err := app.Run([]string{app.Name, InfoCommand.Name, "rack-1", "board-1"})

	test.ExpectNoError(t, err)
}

func TestInfoCommandErrorDevice(t *testing.T) {
	test.Setup()

	mux, server := test.TestServer()
	defer server.Close()
	mux.HandleFunc(
		"/synse/2.0/info/rack-1/board-1/device-1",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
		},
	)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, InfoCommand)

	err := app.Run([]string{app.Name, InfoCommand.Name, "rack-1", "board-1", "device-1"})

	test.ExpectExitCoderError(t, err)
}

func TestInfoCommandSuccessDevice(t *testing.T) {
	test.Setup()

	mux, server := test.TestServer()
	defer server.Close()
	mux.HandleFunc(
		"/synse/2.0/info/rack-1/board-1/device-1",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, infoDeviceRespOK)
		},
	)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, InfoCommand)

	err := app.Run([]string{app.Name, InfoCommand.Name, "rack-1", "board-1", "device-1"})

	test.ExpectNoError(t, err)
}
