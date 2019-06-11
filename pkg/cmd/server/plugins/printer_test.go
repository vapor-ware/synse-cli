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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vapor-ware/synse-client-go/synse/scheme"
)

func TestPluginSummaryRowFunc_nil(t *testing.T) {
	var data *scheme.PluginMeta

	res, err := serverPluginSummaryRowFunc(data)
	assert.Error(t, err)
	assert.Equal(t, ErrNilData, err)
	assert.Nil(t, res)
}

func TestPluginSummaryRowFunc_errDataType(t *testing.T) {
	res, err := serverPluginSummaryRowFunc("")
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidRowData, err)
	assert.Nil(t, res)
}

func TestPluginSummaryRowFunc_active(t *testing.T) {
	data := &scheme.PluginMeta{
		Active: true,
		ID:     "123",
		Tag:    "test/plugin",
		Version: scheme.VersionOptions{
			PluginVersion: "0.0.1",
		},
	}

	res, err := serverPluginSummaryRowFunc(data)
	assert.NoError(t, err)
	assert.Len(t, res, 5)
	assert.Equal(t, res[0], "✓")
	assert.Equal(t, res[1], "123")
	assert.Equal(t, res[2], "0.0.1")
	assert.Equal(t, res[3], "test/plugin")
	assert.Equal(t, res[4], "")
}

func TestPluginSummaryRowFunc_inactive(t *testing.T) {
	data := &scheme.PluginMeta{
		Active: false,
		ID:     "123",
		Tag:    "test/plugin",
		Version: scheme.VersionOptions{
			PluginVersion: "0.0.1",
		},
	}

	res, err := serverPluginSummaryRowFunc(data)
	assert.NoError(t, err)
	assert.Len(t, res, 5)
	assert.Equal(t, res[0], "✗")
	assert.Equal(t, res[1], "123")
	assert.Equal(t, res[2], "0.0.1")
	assert.Equal(t, res[3], "test/plugin")
	assert.Equal(t, res[4], "")
}

func TestPluginRowFunc_nil(t *testing.T) {
	var data *scheme.Plugin

	res, err := serverPluginRowFunc(data)
	assert.Error(t, err)
	assert.Equal(t, ErrNilData, err)
	assert.Nil(t, res)
}

func TestPluginRowFunc_errDataType(t *testing.T) {
	res, err := serverPluginRowFunc("")
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidRowData, err)
	assert.Nil(t, res)
}

func TestPluginRowFunc_active(t *testing.T) {
	data := &scheme.Plugin{
		PluginMeta: scheme.PluginMeta{
			Active: true,
			ID:     "123",
			Tag:    "test/plugin",
			Version: scheme.VersionOptions{
				PluginVersion: "0.0.1",
			},
		},
		Network: scheme.NetworkOptions{
			Protocol: "tcp",
			Address:  "localhost:5001",
		},
		Health: scheme.HealthOptions{
			Status:    "OK",
			Timestamp: "2019-04-22T13:30:00Z",
		},
	}

	res, err := serverPluginRowFunc(data)
	assert.NoError(t, err)
	assert.Len(t, res, 6)
	assert.Equal(t, res[0], "✓")
	assert.Equal(t, res[1], "123")
	assert.Equal(t, res[2], "test/plugin")
	assert.Equal(t, res[3], "tcp://localhost:5001")
	assert.Equal(t, res[4], "OK")
	assert.Equal(t, res[5], "2019-04-22T13:30:00Z")
}

func TestPluginRowFunc_inactive(t *testing.T) {
	data := &scheme.Plugin{
		PluginMeta: scheme.PluginMeta{
			Active: false,
			ID:     "123",
			Tag:    "test/plugin",
			Version: scheme.VersionOptions{
				PluginVersion: "0.0.1",
			},
		},
		Network: scheme.NetworkOptions{
			Protocol: "tcp",
			Address:  "localhost:5001",
		},
		Health: scheme.HealthOptions{
			Status:    "OK",
			Timestamp: "2019-04-22T13:30:00Z",
		},
	}

	res, err := serverPluginRowFunc(data)
	assert.NoError(t, err)
	assert.Len(t, res, 6)
	assert.Equal(t, res[0], "✗")
	assert.Equal(t, res[1], "123")
	assert.Equal(t, res[2], "test/plugin")
	assert.Equal(t, res[3], "tcp://localhost:5001")
	assert.Equal(t, res[4], "OK")
	assert.Equal(t, res[5], "2019-04-22T13:30:00Z")
}

func TestPluginHealthRowFunc_nil(t *testing.T) {
	var data *scheme.PluginHealth

	res, err := serverPluginHealthRowFunc(data)
	assert.Error(t, err)
	assert.Equal(t, ErrNilData, err)
	assert.Nil(t, res)
}

func TestPluginHealthRowFunc_errDataType(t *testing.T) {
	res, err := serverPluginHealthRowFunc("")
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidRowData, err)
	assert.Nil(t, res)
}

func TestPluginHealthRowFunc(t *testing.T) {
	data := &scheme.PluginHealth{
		Status:    "OK",
		Healthy:   []string{"123"},
		Unhealthy: []string{},
		Active:    1,
		Inactive:  0,
	}

	res, err := serverPluginHealthRowFunc(data)
	assert.NoError(t, err)
	assert.Len(t, res, 5)
	assert.Equal(t, res[0], "OK")
	assert.Equal(t, res[1], 1)
	assert.Equal(t, res[2], 0)
	assert.Equal(t, res[3], 1)
	assert.Equal(t, res[4], 0)
}
