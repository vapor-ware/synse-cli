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

package plugins

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vapor-ware/synse-client-go/synse/scheme"
)

func TestPluginSummaries_Sort(t *testing.T) {
	cases := []struct {
		name string
		in   []*scheme.PluginMeta
		out  []*scheme.PluginMeta
	}{
		{
			name: "Empty data",
			in:   []*scheme.PluginMeta{},
			out:  []*scheme.PluginMeta{},
		},
		{
			name: "Single record",
			in: []*scheme.PluginMeta{
				{ID: "1"},
			},
			out: []*scheme.PluginMeta{
				{ID: "1"},
			},
		},
		{
			name: "Records already sorted",
			in: []*scheme.PluginMeta{
				{ID: "1"},
				{ID: "3"},
				{ID: "4"},
				{ID: "6"},
				{ID: "9"},
			},
			out: []*scheme.PluginMeta{
				{ID: "1"},
				{ID: "3"},
				{ID: "4"},
				{ID: "6"},
				{ID: "9"},
			},
		},
		{
			name: "Multiple records unsorted",
			in: []*scheme.PluginMeta{
				{ID: "9"},
				{ID: "3"},
				{ID: "4"},
				{ID: "6"},
				{ID: "1"},
				{ID: "4"},
				{ID: "7"},
				{ID: "17"},
			},
			out: []*scheme.PluginMeta{
				{ID: "1"},
				{ID: "17"},
				{ID: "3"},
				{ID: "4"},
				{ID: "4"},
				{ID: "6"},
				{ID: "7"},
				{ID: "9"},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			sort.Sort(PluginSummaries(c.in))
			for i, item := range c.in {
				assert.Equal(t, item.ID, c.out[i].ID)
			}
		})
	}
}
