// Client extends the https://github.com/dghubble/sling package to provide a
// RESTful client to the openDCRE endpoints. The base url path is constructed
// from the configured openDCRE url as well as the type and version of the API.
// All new queires within vesh should be using an instance of this client.
package client

import (
	//"os"
	"fmt"
	"strconv"

	"github.com/dghubble/sling"
)

// Empty variable to store the content of the VESH_HOST env variable.
var VeshHostPtr = ""
var theClient *sling.Sling

// constructUrl builds the full url string from the host base, endpoint type
// (openDCRE), and API version number. Endpoint paths can be extended off of
// this base.
func constructUrl() string {
	var vaporBase = fmt.Sprint(VeshHostPtr)
	var vaporPort = strconv.Itoa(5000)
	var defaultPath = "opendcre/1.3/" //Add a version number here
	var CompleteBase = fmt.Sprintf("http://%s:%s/%s", vaporBase, vaporPort, defaultPath)
	//fmt.Println(CompleteBase) For bug testing only
	return CompleteBase
}

type ErrorResponse struct { // FIXME: This should go somewhere else
	HttpCode int    `json:"http_code"`
	Message  string `json:"message"`
}

func New() *sling.Sling {
	if theClient == nil {
		theClient = sling.New().Base(constructUrl())
	}

	return theClient.New()
}
