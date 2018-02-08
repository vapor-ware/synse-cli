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
    "board-1"
  ]
}`

	infoBoardRespOK = `
{
  "board":"board-1",
  "location":{
    "rack":"rack-1"
  },
  "devices":[
    "device-1",
    "device-2",
    "device-3",
    "device-4",
    "device-5",
    "device-6",
    "device-7",
    "device-8",
  ]
}`

	infoDeviceRespOK = `
{
  "timestamp":"2018-02-08 15:58:51.063845404 +0000 UTC m=+3105.953837345",
  "uid":"device-1",
  "type":"temperature",
  "model":"emul8-temp",
  "manufacturer":"Vapor IO",
  "protocol":"emulator",
  "info":"Synse Temperature Sensor",
  "comment":"",
  "location":{
    "rack":"rack-1",
    "board":"board-1"
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
