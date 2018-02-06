package utils

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/client"
)

// check is a helper function to check the HTTP response from Synse Server.
// If the request failed with error or returned an error code, it will raise
// an error.
func check(response *http.Response, err error) error {
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	if response.StatusCode != http.StatusOK {
		// TODO (etd) - Synse Server returns JSON for 404 and 500 - we should check
		// that here/log it out.
		return cli.NewExitError(
			fmt.Sprintf("got HTTP code %v for request", response.StatusCode),
			1,
		)
	}
	return nil
}

// MakeURI joins the given components into a string, delimited with '/' which
// can then be used as the URI for API requests.
func MakeURI(components ...string) string {
	return strings.Join(components, "/")
}

// DoGet is a convenience function which performs a GET request against the
// Synse Server versioned API.
func DoGet(uri string, scheme interface{}) error {
	return check(client.New().Get(uri).ReceiveSuccess(scheme))
}

// DoGetUnversioned is a convenience function which performs a GET request against
// the Synse Server unversioned API.
func DoGetUnversioned(uri string, scheme interface{}) error {
	return check(client.NewUnversioned().Get(uri).ReceiveSuccess(scheme))
}
