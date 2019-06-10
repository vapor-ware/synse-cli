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

func TestCmdRead_multipleFormats(t *testing.T) {
	defer resetFlags()

	result := test.Cmd(cmdRead).Args(
		"--json",
		"--yaml",
	).Run(t)
	result.AssertNoErr()
	result.AssertExited()
	result.AssertGolden("multiple-formats.golden")
}

func TestCmdRead_badClient(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseGrpcClient, func(ctx, cert string) (*grpc.ClientConn, synse.V3PluginClient, error) {
		return nil, nil, errors.New("test error message")
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdRead).Run(t)
	result.AssertNoErr()
	result.AssertExited()
	result.AssertGolden("bad-client.golden")
}

func TestCmdRead_requestError(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseGrpcClient, func(ctx, cert string) (*grpc.ClientConn, synse.V3PluginClient, error) {
		return test.NewFakeConn(), test.NewFakeGRPCClientV3Err(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdRead).Run(t)
	result.AssertNoErr()
	result.AssertExited()
	result.AssertGolden("request-err.golden")
}

func TestCmdRead_table(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseGrpcClient, func(ctx, cert string) (*grpc.ClientConn, synse.V3PluginClient, error) {
		return test.NewFakeConn(), test.NewFakeGRPCClientV3(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdRead).Run(t)
	result.AssertNoErr()
	result.AssertGolden("read.table.golden")
}

func TestCmdRead_tableNoHeader(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseGrpcClient, func(ctx, cert string) (*grpc.ClientConn, synse.V3PluginClient, error) {
		return test.NewFakeConn(), test.NewFakeGRPCClientV3(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdRead).Args(
		"--no-header",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("read.table-no-header.golden")
}

func TestCmdRead_json(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseGrpcClient, func(ctx, cert string) (*grpc.ClientConn, synse.V3PluginClient, error) {
		return test.NewFakeConn(), test.NewFakeGRPCClientV3(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdRead).Args(
		"--json",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("read.json.golden")
}

func TestCmdRead_yaml(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseGrpcClient, func(ctx, cert string) (*grpc.ClientConn, synse.V3PluginClient, error) {
		return test.NewFakeConn(), test.NewFakeGRPCClientV3(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdRead).Args(
		"--yaml",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("read.yaml.golden")
}
