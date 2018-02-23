package server

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/vapor-ware/synse-cli/internal/test"
)

const versionRespOK = `
{
  "version": "2.0.0",
  "api_version": "2.0"
}`

func TestVersionCommandError(t *testing.T) {
	test.Setup()

	mux, server := test.UnversionedServer()
	defer server.Close()
	mux.HandleFunc("/synse/version", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
	})

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, versionCommand)

	err := app.Run([]string{app.Name, versionCommand.Name})

	test.ExpectExitCoderError(t, err)
}

func TestVersionCommandSuccess(t *testing.T) {
	test.Setup()

	mux, server := test.UnversionedServer()
	defer server.Close()
	mux.HandleFunc("/synse/version", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, versionRespOK)
	})

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, versionCommand)

	err := app.Run([]string{app.Name, versionCommand.Name})

	test.ExpectNoError(t, err)
}
