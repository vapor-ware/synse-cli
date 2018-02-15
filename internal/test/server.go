package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/vapor-ware/synse-cli/pkg/config"
)

// Server creates a test HTTP server used for testing commands
// that use the Synse Server versioned endpoint.
func Server() (*http.ServeMux, *httptest.Server) {
	mux, server := UnversionedServer()

	mux.HandleFunc("/synse/version", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"version": "2.0.0", "api_version": "2.0"}`) // nolint
	})
	return mux, server
}

// UnversionedServer creates a test HTTP server used for testing
// commands that use the Synse Server unversioned endpoint.
func UnversionedServer() (*http.ServeMux, *httptest.Server) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	return mux, server
}

// AddServerHost is a test helper to add the address of the created
// test server as the active host in the CLI configuration.
func AddServerHost(server *httptest.Server) {
	// parse the URL to remove the 'http://' prefix
	parsedURL := server.URL[7:]

	// create a new host configuration and add it to the CLI config
	// as the active host
	host := config.NewHostConfig("test", parsedURL)
	config.Config.Hosts["test"] = host
	config.Config.ActiveHost = host
}
