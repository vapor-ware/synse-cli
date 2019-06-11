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

package plugin

import (
	"testing"

	"bou.ke/monkey"
	"github.com/pkg/errors"
	"github.com/vapor-ware/synse-cli/internal/test"
	"github.com/vapor-ware/synse-cli/pkg/utils"
	synse "github.com/vapor-ware/synse-server-grpc/go"
	"google.golang.org/grpc"
)

func TestCmdWrite_extraArgs(t *testing.T) {
	defer resetFlags()

	result := test.Cmd(cmdWrite).Args(
		"foo",
		"bar",
		"baz",
		"extra",
	).Run(t)
	result.AssertErr()
	result.AssertGolden("write.extra-args.golden")
}

func TestCmdWrite_multipleFormats(t *testing.T) {
	defer resetFlags()

	result := test.Cmd(cmdWrite).Args(
		"device",
		"action",
		"--json",
		"--yaml",
	).Run(t)
	result.AssertNoErr()
	result.AssertExited()
	result.AssertGolden("multiple-formats.golden")
}

func TestCmdWrite_badClient(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseGrpcClient, func(ctx, cert string) (*grpc.ClientConn, synse.V3PluginClient, error) {
		return nil, nil, errors.New("test error message")
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdWrite).Args(
		"device",
		"action",
	).Run(t)
	result.AssertNoErr()
	result.AssertExited()
	result.AssertGolden("bad-client.golden")
}

func TestCmdWrite_requestError(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseGrpcClient, func(ctx, cert string) (*grpc.ClientConn, synse.V3PluginClient, error) {
		return test.NewFakeConn(), test.NewFakeGRPCClientV3Err(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdWrite).Args(
		"device",
		"action",
	).Run(t)
	result.AssertNoErr()
	result.AssertExited()
	result.AssertGolden("request-err.golden")
}

func TestCmdWriteAsync_table(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseGrpcClient, func(ctx, cert string) (*grpc.ClientConn, synse.V3PluginClient, error) {
		return test.NewFakeConn(), test.NewFakeGRPCClientV3(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdWrite).Args(
		"device",
		"action",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("write-async.table.golden")
}

func TestCmdWriteAsync_tableNoHeader(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseGrpcClient, func(ctx, cert string) (*grpc.ClientConn, synse.V3PluginClient, error) {
		return test.NewFakeConn(), test.NewFakeGRPCClientV3(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdWrite).Args(
		"device",
		"action",
		"--no-header",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("write-async.table-no-header.golden")
}

func TestCmdWriteAsync_json(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseGrpcClient, func(ctx, cert string) (*grpc.ClientConn, synse.V3PluginClient, error) {
		return test.NewFakeConn(), test.NewFakeGRPCClientV3(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdWrite).Args(
		"device",
		"action",
		"--json",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("write-async.json.golden")
}

func TestCmdWriteAsync_yaml(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseGrpcClient, func(ctx, cert string) (*grpc.ClientConn, synse.V3PluginClient, error) {
		return test.NewFakeConn(), test.NewFakeGRPCClientV3(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdWrite).Args(
		"device",
		"action",
		"--yaml",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("write-async.yaml.golden")
}

func TestCmdWriteSync_table(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseGrpcClient, func(ctx, cert string) (*grpc.ClientConn, synse.V3PluginClient, error) {
		return test.NewFakeConn(), test.NewFakeGRPCClientV3(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdWrite).Args(
		"device",
		"action",
		"--wait",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("write-sync.table.golden")
}

func TestCmdWriteSync_tableNoHeader(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseGrpcClient, func(ctx, cert string) (*grpc.ClientConn, synse.V3PluginClient, error) {
		return test.NewFakeConn(), test.NewFakeGRPCClientV3(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdWrite).Args(
		"device",
		"action",
		"--wait",
		"--no-header",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("write-sync.table-no-header.golden")
}

func TestCmdWriteSync_json(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseGrpcClient, func(ctx, cert string) (*grpc.ClientConn, synse.V3PluginClient, error) {
		return test.NewFakeConn(), test.NewFakeGRPCClientV3(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdWrite).Args(
		"device",
		"action",
		"--wait",
		"--json",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("write-sync.json.golden")
}

func TestCmdWriteSync_yaml(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseGrpcClient, func(ctx, cert string) (*grpc.ClientConn, synse.V3PluginClient, error) {
		return test.NewFakeConn(), test.NewFakeGRPCClientV3(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdWrite).Args(
		"device",
		"action",
		"--wait",
		"--yaml",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("write-sync.yaml.golden")
}
