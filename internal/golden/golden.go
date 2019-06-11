// Synse CLI
// Copyright (c) 2019 Vapor IO
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

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
		assert.Equal(t, string(expected), string(actual))
	}
}
