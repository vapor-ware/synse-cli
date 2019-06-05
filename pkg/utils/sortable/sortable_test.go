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

package sortable

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vapor-ware/synse-cli/pkg/config"
)

func TestContextRecords_Sort(t *testing.T) {
	cases := []struct {
		in  []config.ContextRecord
		out []config.ContextRecord
	}{
		// Single empty record
		{
			in: []config.ContextRecord{
				{},
			},
			out: []config.ContextRecord{
				{},
			},
		},
		// Single non-empty record
		{
			in: []config.ContextRecord{
				{Name: "foo"},
			},
			out: []config.ContextRecord{
				{Name: "foo"},
			},
		},
		// Records already sorted
		{
			in: []config.ContextRecord{
				{Name: "a", Type: "plugin"},
				{Name: "b", Type: "plugin"},
				{Name: "a", Type: "server"},
			},
			out: []config.ContextRecord{
				{Name: "a", Type: "plugin"},
				{Name: "b", Type: "plugin"},
				{Name: "a", Type: "server"},
			},
		},
		// Many records, unsorted
		{
			in: []config.ContextRecord{
				{Name: "a", Type: "plugin", Context: config.Context{Address: "1"}},
				{Name: "b", Type: "plugin", Context: config.Context{Address: "1"}},
				{Name: "c", Type: "server", Context: config.Context{Address: "2"}},
				{Name: "aa", Type: "plugin", Context: config.Context{Address: "4"}},
				{Name: "cc", Type: "plugin", Context: config.Context{Address: "3"}},
				{Name: "bc", Type: "server", Context: config.Context{Address: "2"}},
				{Name: "ac", Type: "server", Context: config.Context{Address: "12"}},
				{Name: "ba", Type: "server", Context: config.Context{Address: "14"}},
				{Name: "ca", Type: "plugin", Context: config.Context{Address: "51"}},
				{Name: "a", Type: "plugin", Context: config.Context{Address: "15"}},
				{Name: "a", Type: "server", Context: config.Context{Address: "2"}},
				{Name: "b", Type: "plugin", Context: config.Context{Address: "5"}},
				{Name: "cb", Type: "server", Context: config.Context{Address: "2"}},
				{Name: "bc", Type: "server", Context: config.Context{Address: "1"}},
				{Name: "ac", Type: "plugin", Context: config.Context{Address: "2"}},
				{Name: "ab", Type: "plugin", Context: config.Context{Address: "1"}},
				{Name: "b", Type: "server", Context: config.Context{Address: "1"}},
				{Name: "c", Type: "plugin", Context: config.Context{Address: "2"}},
			},
			out: []config.ContextRecord{
				{Name: "a", Type: "plugin", Context: config.Context{Address: "1"}},
				{Name: "a", Type: "plugin", Context: config.Context{Address: "15"}},
				{Name: "aa", Type: "plugin", Context: config.Context{Address: "4"}},
				{Name: "ab", Type: "plugin", Context: config.Context{Address: "1"}},
				{Name: "ac", Type: "plugin", Context: config.Context{Address: "2"}},
				{Name: "b", Type: "plugin", Context: config.Context{Address: "1"}},
				{Name: "b", Type: "plugin", Context: config.Context{Address: "5"}},
				{Name: "c", Type: "plugin", Context: config.Context{Address: "2"}},
				{Name: "ca", Type: "plugin", Context: config.Context{Address: "51"}},
				{Name: "cc", Type: "plugin", Context: config.Context{Address: "3"}},
				{Name: "a", Type: "server", Context: config.Context{Address: "2"}},
				{Name: "ac", Type: "server", Context: config.Context{Address: "12"}},
				{Name: "b", Type: "server", Context: config.Context{Address: "1"}},
				{Name: "ba", Type: "server", Context: config.Context{Address: "14"}},
				{Name: "bc", Type: "server", Context: config.Context{Address: "1"}},
				{Name: "bc", Type: "server", Context: config.Context{Address: "2"}},
				{Name: "c", Type: "server", Context: config.Context{Address: "2"}},
				{Name: "cb", Type: "server", Context: config.Context{Address: "2"}},
			},
		},
	}

	for _, c := range cases {
		sort.Sort(ContextRecords(c.in))
		assert.Equal(t, c.in, c.out)
	}
}
