package golden

import (
	"flag"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var update = flag.Bool("update", false, "update .golden files")

// Get returns the content of the specified .golden file. If the '-update' flag is
// set (e.g. `go test ./pkg/... -update`), the specified .golden file is updated
// with the current output, then the output is returned.
func Get(t *testing.T, actual []byte, filename string) []byte {
	golden := filepath.Join("testdata", filename)
	if *update {
		if err := ioutil.WriteFile(golden, actual, 0644); err != nil {
			t.Fatal(err)
		}
	}

	expected, err := ioutil.ReadFile(golden)
	if err != nil {
		t.Fatal(err)
	}
	return expected
}

// Check is a test helper which checks if the actual return data matches the expected
// data.
func Check(t *testing.T, actual []byte, filename string) {
	expected := Get(t, actual, filename)

	if actual == nil {
		assert.Empty(t, expected)
	} else {
		assert.Equal(t, expected, actual)
	}
}
