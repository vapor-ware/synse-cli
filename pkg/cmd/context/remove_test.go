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

func TestCmdRemove_noCtxs(t *testing.T) {
	defer func() {
		config.Purge()
		resetFlags()
	}()

	result := test.Cmd(cmdRemove).Run(t)
	result.AssertNoErr()
	result.AssertExited()
	result.AssertGolden("remove.no-ctx.golden")

	assert.Len(t, config.GetContexts(), 0)
	assert.Len(t, config.GetCurrentContext(), 0)
}

func TestCmdRemove_removeNonexistent(t *testing.T) {
	defer func() {
		config.Purge()
		resetFlags()
	}()

	result := test.Cmd(cmdRemove).Args("foo").Run(t)
	result.AssertNoErr()
	result.AssertGolden("empty.golden")

	assert.Len(t, config.GetContexts(), 0)
	assert.Len(t, config.GetCurrentContext(), 0)
}

func TestCmdRemove_removeExisting(t *testing.T) {
	defer func() {
		config.Purge()
		resetFlags()
	}()

	// Set a current server context
	assert.NoError(t, config.AddContext(&config.ContextRecord{
		Name: "server-ctx",
		Type: "server",
		Context: config.Context{
			Address:    "0.0.0.0",
			ClientCert: "/tmp/test/dir",
		},
	}))
	assert.NoError(t, config.SetCurrentContext("server-ctx"))

	// Add a plugin context
	assert.NoError(t, config.AddContext(&config.ContextRecord{
		Name: "plugin-ctx",
		Type: "plugin",
		Context: config.Context{
			Address:    "foo/bar",
			ClientCert: "/tmp/test/dir",
		},
	}))

	assert.Len(t, config.GetContexts(), 2)
	assert.Len(t, config.GetCurrentContext(), 1)

	result := test.Cmd(cmdRemove).Args("server-ctx").Run(t)
	result.AssertNoErr()
	result.AssertGolden("empty.golden")

	assert.Len(t, config.GetContexts(), 1)
	assert.Len(t, config.GetCurrentContext(), 0)
}

func TestCmdRemove_removeMultipleExisting(t *testing.T) {
	defer func() {
		config.Purge()
		resetFlags()
	}()

	// Set a current server context
	assert.NoError(t, config.AddContext(&config.ContextRecord{
		Name: "server-ctx",
		Type: "server",
		Context: config.Context{
			Address:    "0.0.0.0",
			ClientCert: "/tmp/test/dir",
		},
	}))
	assert.NoError(t, config.SetCurrentContext("server-ctx"))

	// Add a plugin context
	assert.NoError(t, config.AddContext(&config.ContextRecord{
		Name: "plugin-ctx",
		Type: "plugin",
		Context: config.Context{
			Address:    "foo/bar",
			ClientCert: "/tmp/test/dir",
		},
	}))

	assert.Len(t, config.GetContexts(), 2)
	assert.Len(t, config.GetCurrentContext(), 1)

	result := test.Cmd(cmdRemove).Args("server-ctx", "plugin-ctx").Run(t)
	result.AssertNoErr()
	result.AssertGolden("empty.golden")

	assert.Len(t, config.GetContexts(), 0)
	assert.Len(t, config.GetCurrentContext(), 0)
}

func TestCmdRemove_removeAll(t *testing.T) {
	defer func() {
		config.Purge()
		resetFlags()
	}()

	// Set a current server context
	assert.NoError(t, config.AddContext(&config.ContextRecord{
		Name: "server-ctx",
		Type: "server",
		Context: config.Context{
			Address:    "0.0.0.0",
			ClientCert: "/tmp/test/dir",
		},
	}))
	assert.NoError(t, config.SetCurrentContext("server-ctx"))

	// Add a plugin context
	assert.NoError(t, config.AddContext(&config.ContextRecord{
		Name: "plugin-ctx",
		Type: "plugin",
		Context: config.Context{
			Address:    "foo/bar",
			ClientCert: "/tmp/test/dir",
		},
	}))

	assert.Len(t, config.GetContexts(), 2)
	assert.Len(t, config.GetCurrentContext(), 1)

	result := test.Cmd(cmdRemove).Args("--all").Run(t)
	result.AssertNoErr()
	result.AssertGolden("empty.golden")

	assert.Len(t, config.GetContexts(), 0)
	assert.Len(t, config.GetCurrentContext(), 0)
}
