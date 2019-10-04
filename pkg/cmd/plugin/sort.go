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

package plugin

import (
	synse "github.com/vapor-ware/synse-server-grpc/go"
)

// Devices implements sort.Interface for Devices results from
// the gRPC API. It sorts based on device ID.
type Devices []*synse.V3Device

func (s Devices) Len() int {
	return len(s)
}

func (s Devices) Less(i, j int) bool {
	return s[i].Id < s[j].Id
}

func (s Devices) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Readings implements sort.Interface for Readings results from
// the gRPC API. It sorts based on device ID and reading type.
type Readings []*synse.V3Reading

func (s Readings) Len() int {
	return len(s)
}

func (s Readings) Less(i, j int) bool {
	if s[i].Id < s[j].Id {
		return true
	}
	if s[i].Id > s[j].Id {
		return false
	}
	return s[i].Type < s[j].Type
}

func (s Readings) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Transactions implements sort.Interface for Transactions results from
// the gRPC API. It sorts based transaction ID.
type Transactions []*synse.V3TransactionStatus

func (s Transactions) Len() int {
	return len(s)
}

func (s Transactions) Less(i, j int) bool {
	return s[i].Id < s[j].Id
}

func (s Transactions) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
