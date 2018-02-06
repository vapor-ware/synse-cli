package utils

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/client"
)

// check
func check(response *http.Response, err error) error {
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	if response.StatusCode != http.StatusOK {
		return cli.NewExitError(
			fmt.Sprintf("got HTTP code %v for request", response.StatusCode),
			1,
		)
	}
	return nil
}

func MakeURI(components ...string) string {
	return strings.Join(components, "/")
}

func DoGet(uri string, scheme interface{}) error {
	return check(client.New().Get(uri).ReceiveSuccess(scheme))
}

func DoGetUnversioned(uri string, scheme interface{}) error {
	return check(client.NewUnversioned().Get(uri).ReceiveSuccess(scheme))
}
