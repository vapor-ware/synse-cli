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

func TestCmdList_multipleFmtFlags(t *testing.T) {
	defer func() {
		config.Purge()
		resetFlags()
	}()

	result := test.Cmd(cmdList).Args(
		"--yaml",
		"--json",
	).Run(t)
	result.AssertNoErr()
	result.AssertExited()
	result.AssertGolden("list.multiple-fmt-flags.golden")

	assert.Len(t, config.GetContexts(), 0)
	assert.Len(t, config.GetCurrentContext(), 0)
}

func TestCmdList_noContexts(t *testing.T) {
	defer func() {
		config.Purge()
		resetFlags()
	}()

	result := test.Cmd(cmdList).Run(t)
	result.AssertNoErr()
	result.AssertGolden("empty.golden")

	assert.Len(t, config.GetContexts(), 0)
	assert.Len(t, config.GetCurrentContext(), 0)
}

func TestCmdList_table(t *testing.T) {
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

	result := test.Cmd(cmdList).Run(t)
	result.AssertNoErr()
	result.AssertGolden("list.table.golden")

	assert.Len(t, config.GetContexts(), 2)
	assert.Len(t, config.GetCurrentContext(), 1)
}

func TestCmdList_tableNoHeader(t *testing.T) {
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

	result := test.Cmd(cmdList).Args(
		"--no-header",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("list.table-no-header.golden")

	assert.Len(t, config.GetContexts(), 2)
	assert.Len(t, config.GetCurrentContext(), 1)
}

func TestCmdList_yaml(t *testing.T) {
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

	result := test.Cmd(cmdList).Args(
		"--yaml",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("list.yaml.golden")

	assert.Len(t, config.GetContexts(), 2)
	assert.Len(t, config.GetCurrentContext(), 1)
}

func TestCmdList_json(t *testing.T) {
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

	result := test.Cmd(cmdList).Args(
		"--json",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("list.json.golden")

	assert.Len(t, config.GetContexts(), 2)
	assert.Len(t, config.GetCurrentContext(), 1)
}
