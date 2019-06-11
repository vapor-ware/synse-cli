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

package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVersion(t *testing.T) {
	v := GetVersion()

	assert.Equal(t, "", v.BuildDate)
	assert.Equal(t, "", v.Commit)
	assert.Equal(t, "", v.Tag)
	assert.Equal(t, "", v.GoVersion)
	assert.Equal(t, "", v.Version)
	assert.NotEmpty(t, v.Arch)
	assert.NotEmpty(t, v.OS)
	assert.NotEmpty(t, v.GoCompiler)
}

func TestGetVersion2(t *testing.T) {
	Commit = "123"
	Tag = "tag-1"

	v := GetVersion()

	assert.Equal(t, "", v.BuildDate)
	assert.Equal(t, "123", v.Commit)
	assert.Equal(t, "tag-1", v.Tag)
	assert.Equal(t, "", v.GoVersion)
	assert.Equal(t, "", v.Version)
	assert.NotEmpty(t, v.Arch)
	assert.NotEmpty(t, v.OS)
	assert.NotEmpty(t, v.GoCompiler)
}
