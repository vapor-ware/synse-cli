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

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vapor-ware/synse-cli/pkg/config"
)

func TestNewSynseHTTPClient_noContext(t *testing.T) {
	client, err := NewSynseHTTPClient("", "")

	assert.Nil(t, client)
	assert.Error(t, err)
}

func TestNewSynseHTTPClient_invalidURI(t *testing.T) {
	defer config.Purge()
	err := config.AddContext(&config.ContextRecord{
		Name: "testctx",
		Type: "server",
		Context: config.Context{
			Address: "foo",
		},
	})
	assert.NoError(t, err)

	client, err := NewSynseHTTPClient("", "")
	assert.Nil(t, client)
	assert.Error(t, err)
}

func TestNewSynseHTTPClient_noCurrentCtx(t *testing.T) {
	defer config.Purge()
	err := config.AddContext(&config.ContextRecord{
		Name: "testctx",
		Type: "server",
		Context: config.Context{
			Address: "localhost:5000",
		},
	})
	assert.NoError(t, err)

	err = config.SetCurrentContext("testctx")
	assert.NoError(t, err)

	client, err := NewSynseHTTPClient("", "")
	assert.NotNil(t, client)
	assert.NoError(t, err)
}

func TestNewSynseHTTPClient_namedContext(t *testing.T) {
	defer config.Purge()
	err := config.AddContext(&config.ContextRecord{
		Name: "testctx",
		Type: "server",
		Context: config.Context{
			Address: "localhost:5000",
		},
	})
	assert.NoError(t, err)

	client, err := NewSynseHTTPClient("testctx", "")
	assert.NotNil(t, client)
	assert.NoError(t, err)
}

func TestNewSynseHTTPClient_currentContextNotSet(t *testing.T) {
	client, err := NewSynseHTTPClient("", "")
	assert.Nil(t, client)
	assert.Error(t, err)
	assert.Equal(t, ErrNoCurrentServerCtx, err)
}

func TestNewSynseHTTPClient_namedContextNotFound(t *testing.T) {
	client, err := NewSynseHTTPClient("testctx", "")
	assert.Nil(t, client)
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidServerCtx, err)
}

func TestNewSynseHTTPClient_notAServerCtx(t *testing.T) {
	defer config.Purge()
	err := config.AddContext(&config.ContextRecord{
		Name: "testctx",
		Type: "plugin",
		Context: config.Context{
			Address: "foo",
		},
	})
	assert.NoError(t, err)

	client, err := NewSynseHTTPClient("testctx", "")
	assert.Nil(t, client)
	assert.Error(t, err)
	assert.Equal(t, ErrNotAServerCtx, err)
}
