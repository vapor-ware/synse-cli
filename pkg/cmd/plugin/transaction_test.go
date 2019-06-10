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

func TestCmdTransaction_multipleFormats(t *testing.T) {
	defer resetFlags()

	result := test.Cmd(cmdTransaction).Args(
		"--json",
		"--yaml",
	).Run(t)
	result.AssertNoErr()
	result.AssertExited()
	result.AssertGolden("multiple-formats.golden")
}

func TestCmdTransaction_badClient(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseGrpcClient, func(ctx, cert string) (*grpc.ClientConn, synse.V3PluginClient, error) {
		return nil, nil, errors.New("test error message")
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdTransaction).Run(t)
	result.AssertNoErr()
	result.AssertExited()
	result.AssertGolden("bad-client.golden")
}

func TestCmdTransaction_requestError(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseGrpcClient, func(ctx, cert string) (*grpc.ClientConn, synse.V3PluginClient, error) {
		return test.NewFakeConn(), test.NewFakeGRPCClientV3Err(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdTransaction).Run(t)
	result.AssertNoErr()
	result.AssertExited()
	result.AssertGolden("request-err.golden")
}

func TestCmdTransactions_table(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseGrpcClient, func(ctx, cert string) (*grpc.ClientConn, synse.V3PluginClient, error) {
		return test.NewFakeConn(), test.NewFakeGRPCClientV3(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdTransaction).Run(t)
	result.AssertNoErr()
	result.AssertGolden("transactions.table.golden")
}

func TestCmdTransactions_tableNoHeader(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseGrpcClient, func(ctx, cert string) (*grpc.ClientConn, synse.V3PluginClient, error) {
		return test.NewFakeConn(), test.NewFakeGRPCClientV3(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdTransaction).Args(
		"--no-header",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("transactions.table-no-header.golden")
}

func TestCmdTransactions_json(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseGrpcClient, func(ctx, cert string) (*grpc.ClientConn, synse.V3PluginClient, error) {
		return test.NewFakeConn(), test.NewFakeGRPCClientV3(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdTransaction).Args(
		"--json",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("transactions.json.golden")
}

func TestCmdTransactions_yaml(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseGrpcClient, func(ctx, cert string) (*grpc.ClientConn, synse.V3PluginClient, error) {
		return test.NewFakeConn(), test.NewFakeGRPCClientV3(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdTransaction).Args(
		"--yaml",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("transactions.yaml.golden")
}

func TestCmdTransaction_table(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseGrpcClient, func(ctx, cert string) (*grpc.ClientConn, synse.V3PluginClient, error) {
		return test.NewFakeConn(), test.NewFakeGRPCClientV3(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdTransaction).Args(
		"123",
		"456",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("transaction.table.golden")
}

func TestCmdTransaction_tableNoHeader(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseGrpcClient, func(ctx, cert string) (*grpc.ClientConn, synse.V3PluginClient, error) {
		return test.NewFakeConn(), test.NewFakeGRPCClientV3(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdTransaction).Args(
		"123",
		"456",
		"--no-header",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("transaction.table-no-header.golden")
}

func TestCmdTransaction_json(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseGrpcClient, func(ctx, cert string) (*grpc.ClientConn, synse.V3PluginClient, error) {
		return test.NewFakeConn(), test.NewFakeGRPCClientV3(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdTransaction).Args(
		"123",
		"456",
		"--json",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("transaction.json.golden")
}

func TestCmdTransaction_yaml(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseGrpcClient, func(ctx, cert string) (*grpc.ClientConn, synse.V3PluginClient, error) {
		return test.NewFakeConn(), test.NewFakeGRPCClientV3(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdTransaction).Args(
		"123",
		"456",
		"--yaml",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("transaction.yaml.golden")
}
