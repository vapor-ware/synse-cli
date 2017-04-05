package commands

import (
	// "reflect"
	"testing"

)

func TestTestApi(t *testing.T) {
	cases := []struct {
		inputUrl string
		expected error
	}{
		{"core.vapor.io", nil},
		{"demo.vapor.io", nil},
		{"someplace.noturl", nil},
	}

	for _, test := range cases {
		
	}
}
