package server

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/vapor-ware/synse-cli/internal/test"
)

const transactionRespOK = `
{
  "id":"b9u6ss6q5i6g020lau6g",
  "context":{
    "action":"color",
    "raw":[
      "000000"
    ]
  },
  "state":"ok",
  "status":"done",
  "created":"2018-02-08 15:36:16.081873199 +0000 UTC m=+1750.971865639",
  "updated":"2018-02-08 15:36:16.081873199 +0000 UTC m=+1750.971865639",
  "message":""
}`

func TestTransactionCommandError(t *testing.T) {
	test.Setup()

	app := test.NewFakeApp()
	app.Commands = append(app.Commands, TransactionCommand)

	err := app.Run([]string{app.Name, TransactionCommand.Name})

	test.ExpectExitCoderError(t, err)
}

func TestTransactionCommandError2(t *testing.T) {
	test.Setup()

	mux, server := test.TestServer()
	defer server.Close()
	mux.HandleFunc(
		"/synse/2.0/transaction/b9u6ss6q5i6g020lau6g",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
		},
	)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, TransactionCommand)

	err := app.Run([]string{app.Name, TransactionCommand.Name, "b9u6ss6q5i6g020lau6g"})

	test.ExpectExitCoderError(t, err)
}

func TestTransactionCommandSuccess(t *testing.T) {
	test.Setup()

	mux, server := test.TestServer()
	defer server.Close()
	mux.HandleFunc(
		"/synse/2.0/transaction/b9u6ss6q5i6g020lau6g",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, transactionRespOK)
		},
	)

	test.AddServerHost(server)
	app := test.NewFakeApp()
	app.Commands = append(app.Commands, TransactionCommand)

	err := app.Run([]string{app.Name, TransactionCommand.Name, "b9u6ss6q5i6g020lau6g"})

	test.ExpectNoError(t, err)
}
