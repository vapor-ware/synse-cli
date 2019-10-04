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

import "github.com/vapor-ware/synse-client-go/synse/scheme"

// PluginSummaries implements sort.Interface for PluginMeta responses
// from the Synse Server client. It sorts based on plugin ID.
type PluginSummaries []*scheme.PluginMeta

func (s PluginSummaries) Len() int {
	return len(s)
}

func (s PluginSummaries) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

func (s PluginSummaries) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
