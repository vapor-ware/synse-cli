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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vapor-ware/synse-cli/pkg/config"
)

func TestContextRowFunc_nil(t *testing.T) {
	var data config.ContextRecord

	res, err := contextRowFunc(data)
	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestContextRowFunc_err(t *testing.T) {
	var data *config.ContextRecord

	res, err := contextRowFunc(data)
	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestContextRowFunc_current(t *testing.T) {
	defer config.Purge()

	var data = config.ContextRecord{
		Name: "test",
		Type: "server",
		Context: config.Context{
			Address: "123",
		},
	}

	assert.NoError(t, config.AddContext(&data))
	assert.NoError(t, config.SetCurrentContext("test"))

	res, err := contextRowFunc(data)
	assert.NoError(t, err)
	assert.Equal(t, []interface{}{"*", "test", "server", "123"}, res)
}

func TestContextRowFunc_notCurrent(t *testing.T) {
	defer config.Purge()

	var data = config.ContextRecord{
		Name: "test",
		Type: "server",
		Context: config.Context{
			Address: "123",
		},
	}

	res, err := contextRowFunc(data)
	assert.NoError(t, err)
	assert.Equal(t, []interface{}{" ", "test", "server", "123"}, res)
}
