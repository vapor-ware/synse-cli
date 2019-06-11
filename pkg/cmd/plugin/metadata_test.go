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

func TestCmdMetadata_multipleFormats(t *testing.T) {
	defer resetFlags()

	result := test.Cmd(cmdMetadata).Args(
		"--json",
		"--yaml",
	).Run(t)
	result.AssertNoErr()
	result.AssertExited()
	result.AssertGolden("multiple-formats.golden")
}

func TestCmdMetadata_badClient(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseGrpcClient, func(ctx, cert string) (*grpc.ClientConn, synse.V3PluginClient, error) {
		return nil, nil, errors.New("test error message")
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdMetadata).Run(t)
	result.AssertNoErr()
	result.AssertExited()
	result.AssertGolden("bad-client.golden")
}

func TestCmdMetadata_requestError(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseGrpcClient, func(ctx, cert string) (*grpc.ClientConn, synse.V3PluginClient, error) {
		return test.NewFakeConn(), test.NewFakeGRPCClientV3Err(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdMetadata).Run(t)
	result.AssertNoErr()
	result.AssertExited()
	result.AssertGolden("request-err.golden")
}

func TestCmdMetadata_table(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseGrpcClient, func(ctx, cert string) (*grpc.ClientConn, synse.V3PluginClient, error) {
		return test.NewFakeConn(), test.NewFakeGRPCClientV3(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdMetadata).Run(t)
	result.AssertNoErr()
	result.AssertGolden("metadata.table.golden")
}

func TestCmdMetadata_tableNoHeader(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseGrpcClient, func(ctx, cert string) (*grpc.ClientConn, synse.V3PluginClient, error) {
		return test.NewFakeConn(), test.NewFakeGRPCClientV3(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdMetadata).Args(
		"--no-header",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("metadata.table-no-header.golden")
}

func TestCmdMetadata_json(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseGrpcClient, func(ctx, cert string) (*grpc.ClientConn, synse.V3PluginClient, error) {
		return test.NewFakeConn(), test.NewFakeGRPCClientV3(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdMetadata).Args(
		"--json",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("metadata.json.golden")
}

func TestCmdMetadata_yaml(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseGrpcClient, func(ctx, cert string) (*grpc.ClientConn, synse.V3PluginClient, error) {
		return test.NewFakeConn(), test.NewFakeGRPCClientV3(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdMetadata).Args(
		"--yaml",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("metadata.yaml.golden")
}
