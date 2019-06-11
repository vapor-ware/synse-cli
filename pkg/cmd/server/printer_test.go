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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vapor-ware/synse-client-go/synse/scheme"
)

func TestReadRowFunc_nil(t *testing.T) {
	var data *scheme.Read

	res, err := serverReadRowFunc(data)
	assert.Error(t, err)
	assert.Equal(t, ErrNilData, err)
	assert.Nil(t, res)
}

func TestReadRowFunc_errDataType(t *testing.T) {
	res, err := serverReadRowFunc("")
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidRowData, err)
	assert.Nil(t, res)
}

func TestReadRowFunc(t *testing.T) {
	var data = &scheme.Read{
		Device: "123",
		Value:  23,
		Unit: scheme.UnitOptions{
			Name:   "percent",
			Symbol: "%",
		},
		Timestamp: "2019-04-22T13:30:00Z",
		Type:      "foo",
	}

	res, err := serverReadRowFunc(data)
	assert.NoError(t, err)
	assert.Len(t, res, 5)
	assert.Equal(t, res[0], "123")
	assert.Equal(t, res[1], 23)
	assert.Equal(t, res[2], "%%")
	assert.Equal(t, res[3], "foo")
	assert.Equal(t, res[4], "2019-04-22T13:30:00Z")
}

func TestScanRowFunc_nil(t *testing.T) {
	var data *scheme.Scan

	res, err := serverScanRowFunc(data)
	assert.Error(t, err)
	assert.Equal(t, ErrNilData, err)
	assert.Nil(t, res)
}

func TestScanRowFunc_errDataType(t *testing.T) {
	res, err := serverScanRowFunc("")
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidRowData, err)
	assert.Nil(t, res)
}

func TestScanRowFunc(t *testing.T) {
	var data = &scheme.Scan{
		ID:   "123",
		Type: "foo",
		Info: "example info",
	}

	res, err := serverScanRowFunc(data)
	assert.NoError(t, err)
	assert.Len(t, res, 3)
	assert.Equal(t, res[0], "123")
	assert.Equal(t, res[1], "foo")
	assert.Equal(t, res[2], "example info")
}

func TestStatusRowFunc_nil(t *testing.T) {
	var data *scheme.Status

	res, err := serverStatusRowFunc(data)
	assert.Error(t, err)
	assert.Equal(t, ErrNilData, err)
	assert.Nil(t, res)
}

func TestStatusRowFunc_errDataType(t *testing.T) {
	res, err := serverStatusRowFunc("")
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidRowData, err)
	assert.Nil(t, res)
}

func TestStatusRowFunc(t *testing.T) {
	var data = &scheme.Status{
		Status:    "ok",
		Timestamp: "2019-04-22T13:30:00Z",
	}

	res, err := serverStatusRowFunc(data)
	assert.NoError(t, err)
	assert.Len(t, res, 2)
	assert.Equal(t, res[0], "ok")
	assert.Equal(t, res[1], "2019-04-22T13:30:00Z")
}

func TestTagsRowFunc_empty(t *testing.T) {
	var data string

	res, err := serverTagsRowFunc(data)
	assert.NoError(t, err)
	assert.Len(t, res, 1)
	assert.Equal(t, res[0], "")
}

func TestTagsRowFunc_errDataType(t *testing.T) {
	res, err := serverTagsRowFunc(123)
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidRowData, err)
	assert.Nil(t, res)
}

func TestTagsRowFunc(t *testing.T) {
	var data = "foo"

	res, err := serverTagsRowFunc(data)
	assert.NoError(t, err)
	assert.Len(t, res, 1)
	assert.Equal(t, res[0], "foo")
}

func TestTransactionRowFunc_nil(t *testing.T) {
	var data *scheme.Transaction

	res, err := serverTransactionRowFunc(data)
	assert.Error(t, err)
	assert.Equal(t, ErrNilData, err)
	assert.Nil(t, res)
}

func TestTransactionRowFunc_errDataType(t *testing.T) {
	res, err := serverTransactionRowFunc("")
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidRowData, err)
	assert.Nil(t, res)
}

func TestTransactionRowFunc(t *testing.T) {
	var data = &scheme.Transaction{
		ID:      "123",
		Status:  "OK",
		Created: "2019-04-22T13:30:00Z",
		Updated: "2019-04-22T13:30:00Z",
	}

	res, err := serverTransactionRowFunc(data)
	assert.NoError(t, err)
	assert.Len(t, res, 5)
	assert.Equal(t, res[0], "123")
	assert.Equal(t, res[1], "OK")
	assert.Equal(t, res[2], "")
	assert.Equal(t, res[3], "2019-04-22T13:30:00Z")
	assert.Equal(t, res[4], "2019-04-22T13:30:00Z")
}

func TestTransactionsRowFunc_empty(t *testing.T) {
	var data string

	res, err := serverTransactionsRowFunc(data)
	assert.NoError(t, err)
	assert.Len(t, res, 1)
	assert.Equal(t, res[0], "")
}

func TestTransactionsRowFunc_errDataType(t *testing.T) {
	res, err := serverTransactionsRowFunc(123)
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidRowData, err)
	assert.Nil(t, res)
}

func TestTransactionsRowFunc(t *testing.T) {
	var data = "foo"

	res, err := serverTransactionsRowFunc(data)
	assert.NoError(t, err)
	assert.Len(t, res, 1)
	assert.Equal(t, res[0], "foo")
}

func TestTransactionSummaryRowFunc_nil(t *testing.T) {
	var data *scheme.Write

	res, err := serverTransactionSummaryRowFunc(data)
	assert.Error(t, err)
	assert.Equal(t, ErrNilData, err)
	assert.Nil(t, res)
}

func TestTransactionSummaryRowFunc_errDataType(t *testing.T) {
	res, err := serverTransactionSummaryRowFunc("")
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidRowData, err)
	assert.Nil(t, res)
}

func TestTransactionSummaryRowFunc(t *testing.T) {
	var data = &scheme.Write{
		ID:      "123",
		Device:  "987",
		Timeout: "30s",
		Context: scheme.WriteData{
			Action: "foo",
			Data:   "bar",
		},
	}

	res, err := serverTransactionSummaryRowFunc(data)
	assert.NoError(t, err)
	assert.Len(t, res, 4)
	assert.Equal(t, res[0], "123")
	assert.Equal(t, res[1], "foo")
	assert.Equal(t, res[2], "bar")
	assert.Equal(t, res[3], "987")
}

func TestVersionRowFunc_nil(t *testing.T) {
	var data *scheme.Version

	res, err := serverVersionRowFunc(data)
	assert.Error(t, err)
	assert.Equal(t, ErrNilData, err)
	assert.Nil(t, res)
}

func TestVersionRowFunc_errDataType(t *testing.T) {
	res, err := serverVersionRowFunc("")
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidRowData, err)
	assert.Nil(t, res)
}

func TestVersionRowFunc(t *testing.T) {
	var data = &scheme.Version{
		Version:    "3.0.0",
		APIVersion: "v3",
	}

	res, err := serverVersionRowFunc(data)
	assert.NoError(t, err)
	assert.Len(t, res, 2)
	assert.Equal(t, res[0], "3.0.0")
	assert.Equal(t, res[1], "v3")
}
