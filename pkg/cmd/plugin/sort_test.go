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
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	synse "github.com/vapor-ware/synse-server-grpc/go"
)

func TestDevices_Sort(t *testing.T) {
	cases := []struct {
		name string
		in   []*synse.V3Device
		out  []*synse.V3Device
	}{
		{
			name: "Empty data",
			in:   []*synse.V3Device{},
			out:  []*synse.V3Device{},
		},
		{
			name: "Single record",
			in: []*synse.V3Device{
				{Id: "1"},
			},
			out: []*synse.V3Device{
				{Id: "1"},
			},
		},
		{
			name: "Records already sorted",
			in: []*synse.V3Device{
				{Id: "1"},
				{Id: "2"},
				{Id: "5"},
				{Id: "8"},
			},
			out: []*synse.V3Device{
				{Id: "1"},
				{Id: "2"},
				{Id: "5"},
				{Id: "8"},
			},
		},
		{
			name: "Multiple records unsorted",
			in: []*synse.V3Device{
				{Id: "1"},
				{Id: "6"},
				{Id: "3"},
				{Id: "2"},
				{Id: "4"},
				{Id: "2"},
				{Id: "11"},
				{Id: "7"},
				{Id: "5"},
			},
			out: []*synse.V3Device{
				{Id: "1"},
				{Id: "11"},
				{Id: "2"},
				{Id: "2"},
				{Id: "3"},
				{Id: "4"},
				{Id: "5"},
				{Id: "6"},
				{Id: "7"},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			sort.Sort(Devices(c.in))
			for i, item := range c.in {
				assert.Equal(t, item.Id, c.out[i].Id)
			}
		})
	}
}

func TestReadings_Sort(t *testing.T) {
	cases := []struct {
		name string
		in   []*synse.V3Reading
		out  []*synse.V3Reading
	}{
		{
			name: "Empty data",
			in:   []*synse.V3Reading{},
			out:  []*synse.V3Reading{},
		},
		{
			name: "Single record",
			in: []*synse.V3Reading{
				{Id: "1", Type: "temp"},
			},
			out: []*synse.V3Reading{
				{Id: "1", Type: "temp"},
			},
		},
		{
			name: "Records already sorted",
			in: []*synse.V3Reading{
				{Id: "1", Type: "led"},
				{Id: "1", Type: "temp"},
				{Id: "2", Type: "temp"},
				{Id: "3", Type: "led"},
			},
			out: []*synse.V3Reading{
				{Id: "1", Type: "led"},
				{Id: "1", Type: "temp"},
				{Id: "2", Type: "temp"},
				{Id: "3", Type: "led"},
			},
		},
		{
			name: "Multiple records unsorted",
			in: []*synse.V3Reading{
				{Id: "1", Type: "led"},
				{Id: "2", Type: "humid"},
				{Id: "1", Type: "temp"},
				{Id: "5", Type: "fan"},
				{Id: "6", Type: "temp"},
				{Id: "3", Type: "pressure"},
				{Id: "6", Type: "temp"},
				{Id: "4", Type: "temp"},
				{Id: "2", Type: "led"},
				{Id: "8", Type: "temp"},
			},
			out: []*synse.V3Reading{
				{Id: "1", Type: "led"},
				{Id: "1", Type: "temp"},
				{Id: "2", Type: "humid"},
				{Id: "2", Type: "led"},
				{Id: "3", Type: "pressure"},
				{Id: "4", Type: "temp"},
				{Id: "5", Type: "fan"},
				{Id: "6", Type: "temp"},
				{Id: "6", Type: "temp"},
				{Id: "8", Type: "temp"},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			sort.Sort(Readings(c.in))
			assert.Equal(t, c.in, c.out)
		})
	}
}

func TestTransactions_Sort(t *testing.T) {
	cases := []struct {
		name string
		in   []*synse.V3TransactionStatus
		out  []*synse.V3TransactionStatus
	}{
		{
			name: "Empty data",
			in:   []*synse.V3TransactionStatus{},
			out:  []*synse.V3TransactionStatus{},
		},
		{
			name: "Single record",
			in: []*synse.V3TransactionStatus{
				{Id: "1"},
			},
			out: []*synse.V3TransactionStatus{
				{Id: "1"},
			},
		},
		{
			name: "Records already sorted",
			in: []*synse.V3TransactionStatus{
				{Id: "1"},
				{Id: "3"},
				{Id: "4"},
				{Id: "7"},
			},
			out: []*synse.V3TransactionStatus{
				{Id: "1"},
				{Id: "3"},
				{Id: "4"},
				{Id: "7"},
			},
		},
		{
			name: "Multiple records unsorted",
			in: []*synse.V3TransactionStatus{
				{Id: "5"},
				{Id: "3"},
				{Id: "7"},
				{Id: "1"},
				{Id: "8"},
				{Id: "4"},
				{Id: "9"},
				{Id: "4"},
				{Id: "18"},
				{Id: "13"},
			},
			out: []*synse.V3TransactionStatus{
				{Id: "1"},
				{Id: "13"},
				{Id: "18"},
				{Id: "3"},
				{Id: "4"},
				{Id: "4"},
				{Id: "5"},
				{Id: "7"},
				{Id: "8"},
				{Id: "9"},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			sort.Sort(Transactions(c.in))
			assert.Equal(t, c.in, c.out)
		})
	}
}
