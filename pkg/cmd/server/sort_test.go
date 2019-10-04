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

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vapor-ware/synse-client-go/synse/scheme"
)

func TestDeviceSummaries_Sort(t *testing.T) {
	cases := []struct {
		name string
		in   []*scheme.Scan
		out  []*scheme.Scan
	}{
		{
			name: "Empty data",
			in:   []*scheme.Scan{},
			out:  []*scheme.Scan{},
		},
		{
			name: "Single record",
			in: []*scheme.Scan{
				{ID: "1"},
			},
			out: []*scheme.Scan{
				{ID: "1"},
			},
		},
		{
			name: "Records already sorted",
			in: []*scheme.Scan{
				{ID: "1"},
				{ID: "3"},
				{ID: "4"},
				{ID: "8"},
			},
			out: []*scheme.Scan{
				{ID: "1"},
				{ID: "3"},
				{ID: "4"},
				{ID: "8"},
			},
		},
		{
			name: "Multiple records unsorted",
			in: []*scheme.Scan{
				{ID: "4"},
				{ID: "1"},
				{ID: "5"},
				{ID: "3"},
				{ID: "6"},
				{ID: "9"},
				{ID: "11"},
				{ID: "6"},
			},
			out: []*scheme.Scan{
				{ID: "1"},
				{ID: "11"},
				{ID: "3"},
				{ID: "4"},
				{ID: "5"},
				{ID: "6"},
				{ID: "6"},
				{ID: "9"},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			sort.Sort(DeviceSummaries(c.in))
			for i, item := range c.in {
				assert.Equal(t, item.ID, c.out[i].ID)
			}
		})
	}
}

func TestReadings_Sort(t *testing.T) {
	cases := []struct {
		name string
		in   []*scheme.Read
		out  []*scheme.Read
	}{
		{
			name: "Empty data",
			in:   []*scheme.Read{},
			out:  []*scheme.Read{},
		},
		{
			name: "Single record",
			in: []*scheme.Read{
				{Device: "dev-1", Type: "temp"},
			},
			out: []*scheme.Read{
				{Device: "dev-1", Type: "temp"},
			},
		},
		{
			name: "Records already sorted",
			in: []*scheme.Read{
				{Device: "dev-1", Type: "led"},
				{Device: "dev-1", Type: "temp"},
				{Device: "dev-3", Type: "temp"},
				{Device: "dev-4", Type: "led"},
			},
			out: []*scheme.Read{
				{Device: "dev-1", Type: "led"},
				{Device: "dev-1", Type: "temp"},
				{Device: "dev-3", Type: "temp"},
				{Device: "dev-4", Type: "led"},
			},
		},
		{
			name: "Multiple records unsorted",
			in: []*scheme.Read{
				{Device: "dev-1", Type: "temp"},
				{Device: "dev-3", Type: "temp"},
				{Device: "dev-2", Type: "temp"},
				{Device: "dev-9", Type: "led"},
				{Device: "dev-5", Type: "temp"},
				{Device: "dev-2", Type: "temp"},
				{Device: "dev-1", Type: "led"},
				{Device: "dev-3", Type: "humidity"},
				{Device: "dev-5", Type: "led"},
				{Device: "dev-6", Type: "temp"},
			},
			out: []*scheme.Read{
				{Device: "dev-1", Type: "led"},
				{Device: "dev-1", Type: "temp"},
				{Device: "dev-2", Type: "temp"},
				{Device: "dev-2", Type: "temp"},
				{Device: "dev-3", Type: "humidity"},
				{Device: "dev-3", Type: "temp"},
				{Device: "dev-5", Type: "led"},
				{Device: "dev-5", Type: "temp"},
				{Device: "dev-6", Type: "temp"},
				{Device: "dev-9", Type: "led"},
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
		in   []*scheme.Transaction
		out  []*scheme.Transaction
	}{
		{
			name: "Empty data",
			in:   []*scheme.Transaction{},
			out:  []*scheme.Transaction{},
		},
		{
			name: "Single record",
			in: []*scheme.Transaction{
				{Device: "dev-1", ID: "123"},
			},
			out: []*scheme.Transaction{
				{Device: "dev-1", ID: "123"},
			},
		},
		{
			name: "Records already sorted",
			in: []*scheme.Transaction{
				{Device: "dev-1", ID: "123"},
				{Device: "dev-1", ID: "126"},
				{Device: "dev-2", ID: "426"},
				{Device: "dev-3", ID: "268"},
			},
			out: []*scheme.Transaction{
				{Device: "dev-1", ID: "123"},
				{Device: "dev-1", ID: "126"},
				{Device: "dev-2", ID: "426"},
				{Device: "dev-3", ID: "268"},
			},
		},
		{
			name: "Multiple records unsorted",
			in: []*scheme.Transaction{
				{Device: "dev-1", ID: "123"},
				{Device: "dev-2", ID: "152"},
				{Device: "dev-1", ID: "637"},
				{Device: "dev-5", ID: "236"},
				{Device: "dev-8", ID: "863"},
				{Device: "dev-6", ID: "234"},
				{Device: "dev-3", ID: "232"},
				{Device: "dev-2", ID: "137"},
				{Device: "dev-1", ID: "123"},
				{Device: "dev-5", ID: "294"},
			},
			out: []*scheme.Transaction{
				{Device: "dev-1", ID: "123"},
				{Device: "dev-1", ID: "123"},
				{Device: "dev-1", ID: "637"},
				{Device: "dev-2", ID: "137"},
				{Device: "dev-2", ID: "152"},
				{Device: "dev-3", ID: "232"},
				{Device: "dev-5", ID: "236"},
				{Device: "dev-5", ID: "294"},
				{Device: "dev-6", ID: "234"},
				{Device: "dev-8", ID: "863"},
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
