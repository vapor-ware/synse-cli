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

package server

import (
	"fmt"
	"testing"

	"bou.ke/monkey"
	"github.com/vapor-ware/synse-cli/internal/test"
	"github.com/vapor-ware/synse-cli/pkg/utils"
	"github.com/vapor-ware/synse-client-go/synse"
)

func TestCmdWrite_extraArgs(t *testing.T) {
	defer resetFlags()

	result := test.Cmd(cmdWrite).Args(
		"foo",
		"bar",
		"baz",
		"foo",
		"extra",
	).Run(t)
	result.AssertErr()
	result.AssertGolden("write.extra-args.golden")
}

func TestCmdWrite_multipleFormats(t *testing.T) {
	defer resetFlags()

	result := test.Cmd(cmdWrite).Args(
		"111-222-333",
		"foo",
		"bar",
		"--json",
		"--yaml",
	).Run(t)
	result.AssertNoErr()
	result.AssertExited()
	result.AssertGolden("multiple-formats.golden")
}

func TestCmdWriteAsync_badClient(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseHTTPClient, func(ctx string, cert string) (synse.Client, error) {
		return nil, fmt.Errorf("test error message")
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdWrite).Args(
		"111-222-333",
		"foo",
	).Run(t)
	result.AssertNoErr()
	result.AssertExited()
	result.AssertGolden("bad-client.golden")
}

func TestCmdWriteAsync_requestError(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseHTTPClient, func(ctx string, cert string) (synse.Client, error) {
		return test.NewFakeHTTPClientV3Err(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdWrite).Args(
		"111-222-333",
		"foo",
		"bar",
	).Run(t)
	result.AssertNoErr()
	result.AssertExited()
	result.AssertGolden("request-err.golden")
}

func TestCmdWriteAsync_table(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseHTTPClient, func(ctx string, cert string) (synse.Client, error) {
		return test.NewFakeHTTPClientV3(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdWrite).Args(
		"111-222-333",
		"foo",
		"bar",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("write.async.table.golden")
}

func TestCmdWriteAsync_tableNoHeader(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseHTTPClient, func(ctx string, cert string) (synse.Client, error) {
		return test.NewFakeHTTPClientV3(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdWrite).Args(
		"111-222-333",
		"foo",
		"bar",
		"--no-header",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("write.async.table-no-header.golden")
}

func TestCmdWriteAsync_json(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseHTTPClient, func(ctx string, cert string) (synse.Client, error) {
		return test.NewFakeHTTPClientV3(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdWrite).Args(
		"111-222-333",
		"foo",
		"bar",
		"--json",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("write.async.json.golden")
}

func TestCmdWriteAsync_yaml(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseHTTPClient, func(ctx string, cert string) (synse.Client, error) {
		return test.NewFakeHTTPClientV3(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdWrite).Args(
		"111-222-333",
		"foo",
		"bar",
		"--yaml",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("write.async.yaml.golden")
}

func TestCmdWriteSync_badClient(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseHTTPClient, func(ctx string, cert string) (synse.Client, error) {
		return nil, fmt.Errorf("test error message")
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdWrite).Args(
		"111-222-333",
		"foo",
		"--wait",
	).Run(t)
	result.AssertNoErr()
	result.AssertExited()
	result.AssertGolden("bad-client.golden")
}

func TestCmdWriteSync_requestError(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseHTTPClient, func(ctx string, cert string) (synse.Client, error) {
		return test.NewFakeHTTPClientV3Err(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdWrite).Args(
		"111-222-333",
		"foo",
		"bar",
		"--wait",
	).Run(t)
	result.AssertNoErr()
	result.AssertExited()
	result.AssertGolden("request-err.golden")
}

func TestCmdWriteSync_table(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseHTTPClient, func(ctx string, cert string) (synse.Client, error) {
		return test.NewFakeHTTPClientV3(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdWrite).Args(
		"111-222-333",
		"foo",
		"bar",
		"--wait",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("write.sync.table.golden")
}

func TestCmdWriteSync_tableNoHeader(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseHTTPClient, func(ctx string, cert string) (synse.Client, error) {
		return test.NewFakeHTTPClientV3(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdWrite).Args(
		"111-222-333",
		"foo",
		"bar",
		"--wait",
		"--no-header",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("write.sync.table-no-header.golden")
}

func TestCmdWriteSync_json(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseHTTPClient, func(ctx string, cert string) (synse.Client, error) {
		return test.NewFakeHTTPClientV3(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdWrite).Args(
		"111-222-333",
		"foo",
		"bar",
		"--wait",
		"--json",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("write.sync.json.golden")
}

func TestCmdWriteSync_yaml(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseHTTPClient, func(ctx string, cert string) (synse.Client, error) {
		return test.NewFakeHTTPClientV3(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdWrite).Args(
		"111-222-333",
		"foo",
		"bar",
		"--wait",
		"--yaml",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("write.sync.yaml.golden")
}
