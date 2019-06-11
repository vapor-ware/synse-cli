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
	"github.com/vapor-ware/synse-cli/internal/test"
	"github.com/vapor-ware/synse-cli/pkg/config"
)

func TestCmdAdd_notEnoughArgs(t *testing.T) {
	defer func() {
		config.Purge()
		resetFlags()
	}()

	result := test.Cmd(cmdAdd).Run(t)
	result.AssertErr()
	result.AssertGolden("add.no-args.golden")

	assert.Len(t, config.GetContexts(), 0)
	assert.Len(t, config.GetCurrentContext(), 0)
}

func TestCmdAdd_badCtxType(t *testing.T) {
	defer func() {
		config.Purge()
		resetFlags()
	}()

	result := test.Cmd(cmdAdd).Args(
		"bad-type",
		"test-name",
		"test-address",
	).Run(t)
	result.AssertNoErr()
	result.AssertExited()
	result.AssertGolden("add.bad-type.golden")

	assert.Len(t, config.GetContexts(), 0)
	assert.Len(t, config.GetCurrentContext(), 0)
}

func TestCmdAdd_duplicate(t *testing.T) {
	defer func() {
		config.Purge()
		resetFlags()
	}()

	assert.NoError(t, config.AddContext(&config.ContextRecord{
		Name: "test-ctx",
		Type: "server",
	}))

	result := test.Cmd(cmdAdd).Args(
		"server",
		"test-ctx",
		"test-address",
	).Run(t)
	result.AssertNoErr()
	result.AssertExited()
	result.AssertGolden("add.duplicate-add.golden")

	assert.Len(t, config.GetContexts(), 1)
	assert.Len(t, config.GetCurrentContext(), 0)
}

func TestCmdAdd_addContext(t *testing.T) {
	defer func() {
		config.Purge()
		resetFlags()
	}()

	result := test.Cmd(cmdAdd).Args(
		"server",
		"test-name",
		"test-address",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("empty.golden")

	assert.Len(t, config.GetContexts(), 1)
	assert.Len(t, config.GetCurrentContext(), 0)

	ctx := config.GetContexts()[0]
	assert.Equal(t, ctx.Name, "test-name")
	assert.Equal(t, ctx.Type, "server")
	assert.Equal(t, ctx.Context.Address, "test-address")
	assert.Equal(t, ctx.Context.ClientCert, "")
}

func TestCmdAdd_addAndSetContext(t *testing.T) {
	defer func() {
		config.Purge()
		resetFlags()
	}()

	result := test.Cmd(cmdAdd).Args(
		"server",
		"test-name",
		"test-address",
		"--set",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("empty.golden")

	assert.Len(t, config.GetContexts(), 1)
	assert.Len(t, config.GetCurrentContext(), 1)

	ctx := config.GetContexts()[0]
	assert.Equal(t, ctx.Name, "test-name")
	assert.Equal(t, ctx.Type, "server")
	assert.Equal(t, ctx.Context.Address, "test-address")
	assert.Equal(t, ctx.Context.ClientCert, "")

	serverCtx, ok := config.GetCurrentContext()["server"]
	assert.True(t, ok)
	assert.Equal(t, serverCtx.Name, "test-name")
	assert.Equal(t, serverCtx.Type, "server")
	assert.Equal(t, serverCtx.Context.Address, "test-address")
	assert.Equal(t, serverCtx.Context.ClientCert, "")

	pluginCtx, ok := config.GetCurrentContext()["plugin"]
	assert.False(t, ok)
	assert.Nil(t, pluginCtx)
}
