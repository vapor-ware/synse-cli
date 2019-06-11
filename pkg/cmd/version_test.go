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

package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vapor-ware/synse-cli/internal/golden"
	"github.com/vapor-ware/synse-cli/pkg"
)

func TestDisplayVersion_simple(t *testing.T) {
	defer resetFlags()

	pkg.Version = "3.0.0"
	flagSimple = true

	out := bytes.Buffer{}
	err := displayVersion(&out)
	assert.NoError(t, err)
	golden.Check(t, out.Bytes(), "version.simple.golden")
}
