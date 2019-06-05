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

func TestCmdList(t *testing.T) {
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
	result.AssertGolden("list.golden")

	assert.Len(t, config.GetContexts(), 2)
	assert.Len(t, config.GetCurrentContext(), 1)
}
