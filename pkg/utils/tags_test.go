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

package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeTags(t *testing.T) {
	cases := []struct {
		description string
		tags        []string
		expected    []string
	}{
		{
			description: "empty tags list",
			tags:        []string{},
			expected:    nil,
		},
		{
			description: "single tag string, not comma separated",
			tags:        []string{"test/1"},
			expected:    []string{"test/1"},
		},
		{
			description: "single tag string, comma separated",
			tags:        []string{"test/1,test/2"},
			expected:    []string{"test/1", "test/2"},
		},
		{
			description: "multiple tag strings, not comma separated",
			tags:        []string{"test/1", "test/2"},
			expected:    []string{"test/1", "test/2"},
		},
		{
			description: "multiple tag strings, some comma separated",
			tags:        []string{"test/1", "test/2,test/3"},
			expected:    []string{"test/1", "test/2", "test/3"},
		},
		{
			description: "multiple tag strings, all comma separated",
			tags:        []string{"test/1,test/2", "test/3,test/4", "test/5,test/6,test/7"},
			expected:    []string{"test/1", "test/2", "test/3", "test/4", "test/5", "test/6", "test/7"},
		},
	}

	for _, c := range cases {
		actual := NormalizeTags(c.tags)
		assert.Equal(t, c.expected, actual, c.description)
	}
}
