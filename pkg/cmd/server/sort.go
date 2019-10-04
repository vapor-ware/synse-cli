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

package server

import "github.com/vapor-ware/synse-client-go/synse/scheme"

// DeviceSummaries implements sort.Interface for Scan responses from the
// Synse Server client. It sorts based on the device ID.
type DeviceSummaries []*scheme.Scan

func (s DeviceSummaries) Len() int {
	return len(s)
}

func (s DeviceSummaries) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

func (s DeviceSummaries) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Readings implements sort.Interface for Read responses from the
// Synse Server client. It sorts based on the ID of the device from
// which the reading originated, and the type of the reading.
type Readings []*scheme.Read

func (s Readings) Len() int {
	return len(s)
}

func (s Readings) Less(i, j int) bool {
	if s[i].Device < s[j].Device {
		return true
	}
	if s[i].Device > s[j].Device {
		return false
	}
	return s[i].Type < s[j].Type
}

func (s Readings) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Transactions implements sort.Interface for Transaction responses from the
// Synse Server client. It sorts based on the device ID and transaction ID.
type Transactions []*scheme.Transaction

func (s Transactions) Len() int {
	return len(s)
}

func (s Transactions) Less(i, j int) bool {
	if s[i].Device < s[j].Device {
		return true
	}
	if s[i].Device > s[j].Device {
		return false
	}
	return s[i].ID < s[j].ID
}

func (s Transactions) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
