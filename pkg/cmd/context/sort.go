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

package context

import (
	"github.com/vapor-ware/synse-cli/pkg/config"
)

// Records implements sort.Interface based on the default
// of the Type, Name, and Address fields.
type Records []config.ContextRecord

func (r Records) Len() int {
	return len(r)
}

func (r Records) Less(i, j int) bool {
	if r[i].Type < r[j].Type {
		return true
	}
	if r[i].Type > r[j].Type {
		return false
	}
	if r[i].Name < r[j].Name {
		return true
	}
	if r[i].Name > r[j].Name {
		return false
	}
	return r[i].Context.Address < r[j].Context.Address
}

func (r Records) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
