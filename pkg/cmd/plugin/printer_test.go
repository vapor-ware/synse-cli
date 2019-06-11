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
	"testing"

	"github.com/stretchr/testify/assert"
	synse "github.com/vapor-ware/synse-server-grpc/go"
)

func TestTestRowFunc_nil(t *testing.T) {
	var data *synse.V3TestStatus

	res, err := pluginTestRowFunc(data)
	assert.Error(t, err)
	assert.Equal(t, ErrNilData, err)
	assert.Nil(t, res)
}

func TestTestRowFunc_errDataType(t *testing.T) {
	res, err := pluginTestRowFunc("")
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidRowData, err)
	assert.Nil(t, res)
}

func TestTestRowFunc_ok(t *testing.T) {
	var data = &synse.V3TestStatus{
		Ok: true,
	}

	res, err := pluginTestRowFunc(data)
	assert.NoError(t, err)
	assert.Len(t, res, 1)
	assert.Equal(t, res[0], "OK")
}

func TestTestRowFunc_notOk(t *testing.T) {
	var data = &synse.V3TestStatus{
		Ok: false,
	}

	res, err := pluginTestRowFunc(data)
	assert.NoError(t, err)
	assert.Len(t, res, 1)
	assert.Equal(t, res[0], "ERROR")
}

func TestVersionRowFunc_nil(t *testing.T) {
	var data *synse.V3Version

	res, err := pluginVersionRowFunc(data)
	assert.Error(t, err)
	assert.Equal(t, ErrNilData, err)
	assert.Nil(t, res)
}

func TestVersionRowFunc_errDataType(t *testing.T) {
	res, err := pluginVersionRowFunc("")
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidRowData, err)
	assert.Nil(t, res)
}

func TestVersionRowFunc_ok(t *testing.T) {
	var data = &synse.V3Version{
		PluginVersion: "3.2.1",
		SdkVersion:    "3.0.0",
		BuildDate:     "now",
		Os:            "foo",
		Arch:          "bar",
	}

	res, err := pluginVersionRowFunc(data)
	assert.NoError(t, err)
	assert.Len(t, res, 5)
	assert.Equal(t, res[0], "3.2.1")
	assert.Equal(t, res[1], "3.0.0")
	assert.Equal(t, res[2], "now")
	assert.Equal(t, res[3], "foo")
	assert.Equal(t, res[4], "bar")
}

func TestReadingRowFunc_nil(t *testing.T) {
	var data *synse.V3Reading

	res, err := pluginReadingRowFunc(data)
	assert.Error(t, err)
	assert.Equal(t, ErrNilData, err)
	assert.Nil(t, res)
}

func TestReadingRowFunc_errDataType(t *testing.T) {
	res, err := pluginReadingRowFunc("")
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidRowData, err)
	assert.Nil(t, res)
}

func TestReadingRowFunc_ok(t *testing.T) {
	var data = &synse.V3Reading{
		Id:         "123",
		Timestamp:  "2019-04-22T13:30:00Z",
		Type:       "test-type",
		DeviceType: "test-device",
		Unit: &synse.V3OutputUnit{
			Name:   "percent",
			Symbol: "%",
		},
		Value: &synse.V3Reading_Float64Value{
			Float64Value: 34,
		},
	}

	res, err := pluginReadingRowFunc(data)
	assert.NoError(t, err)
	assert.Len(t, res, 5)
	assert.Equal(t, res[0], "123")
	assert.Equal(t, res[1], &synse.V3Reading_Float64Value{Float64Value: 34})
	assert.Equal(t, res[2], "%%")
	assert.Equal(t, res[3], "test-type")
	assert.Equal(t, res[4], "2019-04-22T13:30:00Z")
}

func TestTransactionStatusRowFunc_nil(t *testing.T) {
	var data *synse.V3TransactionStatus

	res, err := pluginTransactionStatusRowFunc(data)
	assert.Error(t, err)
	assert.Equal(t, ErrNilData, err)
	assert.Nil(t, res)
}

func TestTransactionStatusRowFunc_errDataType(t *testing.T) {
	res, err := pluginTransactionStatusRowFunc("")
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidRowData, err)
	assert.Nil(t, res)
}

func TestTransactionStatusRowFunc_ok(t *testing.T) {
	var data = &synse.V3TransactionStatus{
		Created: "2019-04-22T13:30:00Z",
		Updated: "2019-04-22T13:30:00Z",
		Timeout: "30s",
		Id:      "123-456",
		Status:  synse.WriteStatus_DONE,
	}

	res, err := pluginTransactionStatusRowFunc(data)
	assert.NoError(t, err)
	assert.Len(t, res, 5)
	assert.Equal(t, res[0], "123-456")
	assert.Equal(t, res[1], synse.WriteStatus_DONE)
	assert.Equal(t, res[2], "")
	assert.Equal(t, res[3], "2019-04-22T13:30:00Z")
	assert.Equal(t, res[4], "2019-04-22T13:30:00Z")
}

func TestTransactionInfoRowFunc_nil(t *testing.T) {
	var data *synse.V3WriteTransaction

	res, err := pluginTransactionInfoRowFunc(data)
	assert.Error(t, err)
	assert.Equal(t, ErrNilData, err)
	assert.Nil(t, res)
}

func TestTransactionInfoRowFunc_errDataType(t *testing.T) {
	res, err := pluginTransactionInfoRowFunc("")
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidRowData, err)
	assert.Nil(t, res)
}

func TestTransactionInfoRowFunc_ok(t *testing.T) {
	var data = &synse.V3WriteTransaction{
		Device:  "987-654",
		Timeout: "30s",
		Id:      "123-456",
		Context: &synse.V3WriteData{
			Action: "foo",
			Data:   []byte("bar"),
		},
	}

	res, err := pluginTransactionInfoRowFunc(data)
	assert.NoError(t, err)
	assert.Len(t, res, 4)
	assert.Equal(t, res[0], "123-456")
	assert.Equal(t, res[1], "foo")
	assert.Equal(t, res[2], "bar")
	assert.Equal(t, res[3], "987-654")
}

func TestMetadataRowFunc_nil(t *testing.T) {
	var data *synse.V3Metadata

	res, err := pluginMetadataRowFunc(data)
	assert.Error(t, err)
	assert.Equal(t, ErrNilData, err)
	assert.Nil(t, res)
}

func TestMetadataRowFunc_errDataType(t *testing.T) {
	res, err := pluginMetadataRowFunc("")
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidRowData, err)
	assert.Nil(t, res)
}

func TestMetadataRowFunc_ok(t *testing.T) {
	var data = &synse.V3Metadata{
		Id:          "123-456",
		Name:        "foo",
		Maintainer:  "bar",
		Tag:         "bar/foo",
		Description: "a plugin",
	}

	res, err := pluginMetadataRowFunc(data)
	assert.NoError(t, err)
	assert.Len(t, res, 3)
	assert.Equal(t, res[0], "123-456")
	assert.Equal(t, res[1], "bar/foo")
	assert.Equal(t, res[2], "a plugin")
}

func TestDeviceRowFunc_nil(t *testing.T) {
	var data *synse.V3Device

	res, err := pluginDeviceRowFunc(data)
	assert.Error(t, err)
	assert.Equal(t, ErrNilData, err)
	assert.Nil(t, res)
}

func TestDeviceRowFunc_errDataType(t *testing.T) {
	res, err := pluginDeviceRowFunc("")
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidRowData, err)
	assert.Nil(t, res)
}

func TestDeviceRowFunc_ok(t *testing.T) {
	var data = &synse.V3Device{
		Id:        "123-456",
		Timestamp: "2019-04-22T13:30:00Z",
		Type:      "test-type",
		Plugin:    "999-888",
		Info:      "example device info",
		Alias:     "test-dev",
	}

	res, err := pluginDeviceRowFunc(data)
	assert.NoError(t, err)
	assert.Len(t, res, 5)
	assert.Equal(t, res[0], "123-456")
	assert.Equal(t, res[1], "test-dev")
	assert.Equal(t, res[2], "test-type")
	assert.Equal(t, res[3], "example device info")
	assert.Equal(t, res[4], "999-888")
}

func TestHealthRowFunc_nil(t *testing.T) {
	var data *synse.V3Health

	res, err := pluginHealthRowFunc(data)
	assert.Error(t, err)
	assert.Equal(t, ErrNilData, err)
	assert.Nil(t, res)
}

func TestHealthRowFunc_errDataType(t *testing.T) {
	res, err := pluginHealthRowFunc("")
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidRowData, err)
	assert.Nil(t, res)
}

func TestHealthRowFunc_ok(t *testing.T) {
	var data = &synse.V3Health{
		Timestamp: "2019-04-22T13:30:00Z",
		Status:    synse.HealthStatus_OK,
		Checks: []*synse.V3HealthCheck{
			{
				Name:      "foo",
				Status:    synse.HealthStatus_OK,
				Timestamp: "2019-04-22T13:30:00Z",
			},
		},
	}

	res, err := pluginHealthRowFunc(data)
	assert.NoError(t, err)
	assert.Len(t, res, 3)
	assert.Equal(t, res[0], synse.HealthStatus_OK)
	assert.Equal(t, res[1], "2019-04-22T13:30:00Z")
	assert.Equal(t, res[2], 1)
}
