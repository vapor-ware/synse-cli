package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/vapor-ware/synse-cli/config"
)

func TestServer() (*http.ServeMux, *httptest.Server) {
	mux, server := TestUnversionedServer()

	mux.HandleFunc("/synse/version", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"version": "2.0.0", "api_version": "2.0"}`)
	})
	return mux, server
}

func TestUnversionedServer() (*http.ServeMux, *httptest.Server) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	return mux, server
}

func AddServerHost(server *httptest.Server) {
	// parse the URL to remove the 'http://' prefix
	parsedURL := server.URL[7:]

	// create a new host configuration and add it to the CLI config
	// as the active host
	host := config.NewHostConfig("test", parsedURL)
	config.Config.Hosts["test"] = host
	config.Config.ActiveHost = host
}
