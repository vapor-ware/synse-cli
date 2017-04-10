package client

import (
	"fmt"
	"net/http"
	"error"
	"testing"

)

func TestConstructUrl(t *testing.T) {
	var cases []struct {
		URL string
		vaporBase string
		vaporPort int
		defaultPath string
		CompleteBase string
		requestErr error
		httpCode *http.StatusCode
		httpErrMessage ErrorResponse.Message
		}{
			{"demo.vapor.io", 5000, "opendcre/1.3/", fmt.Sprintf("http://%s:%d/%s", vaporBase, vaporPort, defaultPath), 200, nil, ""},
			{"core.vapor.io", 5000, "opendcre/1.3/", fmt.Sprintf("http://%s:%d/%s", vaporBase, vaporPort, defaultPath), 200, nil, ""},
			{"someplace.noturl", 5000, "opendcre/1.3/", fmt.Sprintf("http://%s:%d/%s", vaporBase, vaporPort, defaultPath), 404, nil, ""},
		}
	}

func TestDo(t *testing.T) {
	var cases []struct {
		log *LogMiddleware
		req *http.Request
		expectedResult string
		expectedError error
	}{
		{log.LogMiddleware, http.Request, "", nil},
		{log.LogMiddleware, http.Request, fmt.Sprintf("%s", req.Body), nil},
	}
}

func TestNew(t *testing.T) {
	var cases []struct {
		client *sling.Sling
		expectedHttpCode int
	}{
		{sling.Sling, 200},
	}
}
