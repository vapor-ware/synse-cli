package server

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/vapor-ware/synse-cli/internal/test"
)

const statusRespOK = `
{
  "status": "ok",
  "timestamp": "2018-01-01 01:01:01.000000"
}`

func TestStatusCommandError(t *testing.T) {
	test.Setup()

	mux, server := test.UnversionedServer()
	defer server.Close()
	mux.HandleFunc("/synse/test", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
	})

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, StatusCommand)

	err := app.Run([]string{app.Name, StatusCommand.Name})

	test.ExpectExitCoderError(t, err)
}

func TestStatusCommandSuccess(t *testing.T) {
	test.Setup()

	mux, server := test.UnversionedServer()
	defer server.Close()
	mux.HandleFunc("/synse/test", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, statusRespOK)
	})

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, StatusCommand)

	err := app.Run([]string{app.Name, StatusCommand.Name})

	test.ExpectNoError(t, err)
}
