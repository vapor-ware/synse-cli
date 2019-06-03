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

func TestDoc(t *testing.T) {
	cases := []struct {
		description string
		raw         string
		expected    string
	}{
		{
			description: "empty string",
			raw:         ``,
			expected:    "",
		},
		{
			description: "empty string with newlines",
			raw: `

`,
			expected: "\n",
		},
		{
			description: "string without offset tabs",
			raw:         `a simple test string`,
			expected:    "a simple test string",
		},
		{
			description: "string with offset tabs",
			raw: `
				a string offset for nice
				inline formatting
			`,
			expected: "a string offset for nice\ninline formatting\n",
		},
		{
			description: "string with offset tabs, multiple levels",
			raw: `
				multiple levels of
					string offset where
						the levels should be normalized
			`,
			expected: "multiple levels of\n\tstring offset where\n\t\tthe levels should be normalized\n",
		},
		{
			description: "string with console underscore marker",
			raw:         `a string with <underscore>underscore</>`,
			expected:    "a string with \x1b[4munderscore\x1b[0m",
		},
		{
			description: "string with console bold marker",
			raw:         `a string with <bold>bold</>`,
			expected:    "a string with \x1b[1mbold\x1b[0m",
		},
		{
			description: "string with console color marker",
			raw:         `a string with <cyan>color</>`,
			expected:    "a string with \x1b[0;36mcolor\x1b[0m",
		},
	}

	for _, c := range cases {
		actual := Doc(c.raw)
		assert.Equal(t, c.expected, actual, c.description)
	}
}
