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
