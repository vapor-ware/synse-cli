// Client extends the https://github.com/dghubble/sling package to provide a
// RESTful client to the openDCRE endpoints. The base url path is constructed
// from the configured openDCRE url as well as the type and version of the API.
// All new queires within vesh should be using an instance of this client.
package client

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/dghubble/sling"
)

var theClient *sling.Sling

// constructUrl builds the full url string from the host base, endpoint type
// (openDCRE), and API version number. Endpoint paths can be extended off of
// this base.
func constructUrl(host string) string {
	var vaporPort = 5000
	var defaultPath = "opendcre/1.3/" //Add a version number here
	var CompleteBase = fmt.Sprintf(
		"http://%s:%d/%s", host, vaporPort, defaultPath)
	return CompleteBase
}

type ErrorResponse struct { // FIXME: This should go somewhere else
	HttpCode int    `json:"http_code"`
	Message  string `json:"message"`
}

type LogMiddleware struct {
	c http.Client
}

func track(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func (d LogMiddleware) Do(req *http.Request) (*http.Response, error) {
	log.WithFields(log.Fields{
		"method": req.Method,
		"url":    fmt.Sprintf("%v", req.URL),
		// We're not doing anything with headers or forms yet. Once we do, turn
		// these on.
		// ----
		// "header": fmt.Sprintf("%v", req.Header),
		// "form": fmt.Sprintf("%v", req.Form),
	}).Debug("request: start")

	start := time.Now()
	resp, err := d.c.Do(req)
	elapsed := time.Since(start)

	status := ""
	if resp != nil {
		status = resp.Status
	}

	log.WithFields(log.Fields{
		"duration": elapsed,
		"url":      fmt.Sprintf("%v", req.URL),
		"status":   status,
	}).Debug("request: complete")

	return resp, err
}

func Config(host string) {
	theClient = sling.New().Doer(&LogMiddleware{}).Base(constructUrl(host))
}

func New() *sling.Sling {
	if theClient == nil {
		panic("You must configure the client first.")
	}

	return theClient.New()
}
