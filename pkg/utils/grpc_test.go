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

func TestNewSynseGrpcClient_noContext(t *testing.T) {
	conn, client, err := NewSynseGrpcClient("", "")

	assert.Nil(t, conn)
	assert.Nil(t, client)
	assert.Error(t, err)
}

func TestNewSynseGrpcClient_currentContext(t *testing.T) {
	defer config.Purge()
	err := config.AddContext(&config.ContextRecord{
		Name: "testctx",
		Type: "plugin",
		Context: config.Context{
			Address: "foo",
		},
	})
	assert.NoError(t, err)

	err = config.SetCurrentContext("testctx")
	assert.NoError(t, err)

	conn, client, err := NewSynseGrpcClient("", "")
	assert.NotNil(t, conn)
	assert.NotNil(t, client)
	assert.NoError(t, err)
}

func TestNewSynseGrpcClient_namedContext(t *testing.T) {
	defer config.Purge()
	err := config.AddContext(&config.ContextRecord{
		Name: "testctx",
		Type: "plugin",
		Context: config.Context{
			Address: "foo",
		},
	})
	assert.NoError(t, err)

	conn, client, err := NewSynseGrpcClient("testctx", "")
	assert.NotNil(t, conn)
	assert.NotNil(t, client)
	assert.NoError(t, err)
}

func TestNewSynseGrpcClient_currentContextNotSet(t *testing.T) {
	conn, client, err := NewSynseGrpcClient("", "")
	assert.Nil(t, conn)
	assert.Nil(t, client)
	assert.Error(t, err)
	assert.Equal(t, ErrNoCurrentPluginCtx, err)
}

func TestNewSynseGrpcClient_namedContextNotFound(t *testing.T) {
	conn, client, err := NewSynseGrpcClient("testctx", "")
	assert.Nil(t, conn)
	assert.Nil(t, client)
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidPluginCtx, err)
}

func TestNewSynseGrpcClient_notAPluginCtx(t *testing.T) {
	defer config.Purge()
	err := config.AddContext(&config.ContextRecord{
		Name: "testctx",
		Type: "server",
		Context: config.Context{
			Address: "foo",
		},
	})
	assert.NoError(t, err)

	conn, client, err := NewSynseGrpcClient("testctx", "")
	assert.Nil(t, conn)
	assert.Nil(t, client)
	assert.Error(t, err)
	assert.Equal(t, ErrNotAPluginCtx, err)
}

func TestNewSynseGrpcClient_invalidCert(t *testing.T) {
	defer config.Purge()
	err := config.AddContext(&config.ContextRecord{
		Name: "testctx",
		Type: "plugin",
		Context: config.Context{
			Address: "foo",
		},
	})
	assert.NoError(t, err)

	conn, client, err := NewSynseGrpcClient("testctx", "not-a-cert")
	assert.Nil(t, conn)
	assert.Nil(t, client)
	assert.Error(t, err)
}
