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

import "strings"

// NormalizeTags takes a slice specifying tags which may be comma-separated
// and produces a slice of tags where each element is an individual tag.
func NormalizeTags(tags []string) []string {
	var final []string
	for _, tag := range tags {
		final = append(final, strings.Split(tag, ",")...)
	}
	return final
}
