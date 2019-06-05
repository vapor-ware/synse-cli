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

func TestCmdSet_extraArgs(t *testing.T) {
	defer func() {
		config.Purge()
		resetFlags()
	}()

	result := test.Cmd(cmdSet).Args(
		"foo",
		"bar",
	).Run(t)
	result.AssertErr()
	result.AssertGolden("set.extra-args.golden")

	assert.Len(t, config.GetContexts(), 0)
	assert.Len(t, config.GetCurrentContext(), 0)
}

func TestCmdSet_existingCtx(t *testing.T) {
	defer func() {
		config.Purge()
		resetFlags()
	}()

	assert.NoError(t, config.AddContext(&config.ContextRecord{
		Name: "test-ctx",
		Type: "server",
		Context: config.Context{
			Address:    "0.0.0.0",
			ClientCert: "/tmp/test/dir",
		},
	}))

	assert.Len(t, config.GetContexts(), 1)
	assert.Len(t, config.GetCurrentContext(), 0)

	result := test.Cmd(cmdSet).Args(
		"test-ctx",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("empty.golden")

	assert.Len(t, config.GetContexts(), 1)
	assert.Len(t, config.GetCurrentContext(), 1)
}

func TestCmdSet_nonexistingCtx(t *testing.T) {
	defer func() {
		config.Purge()
		resetFlags()
	}()

	assert.NoError(t, config.AddContext(&config.ContextRecord{
		Name: "test-ctx",
		Type: "server",
		Context: config.Context{
			Address:    "0.0.0.0",
			ClientCert: "/tmp/test/dir",
		},
	}))

	assert.Len(t, config.GetContexts(), 1)
	assert.Len(t, config.GetCurrentContext(), 0)

	result := test.Cmd(cmdSet).Args(
		"foo",
	).Run(t)
	result.AssertNoErr()
	result.AssertExited()
	result.AssertGolden("set.nonexisting.golden")

	assert.Len(t, config.GetContexts(), 1)
	assert.Len(t, config.GetCurrentContext(), 0)
}
