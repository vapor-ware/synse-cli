package test

import (
	"context"

	synse "github.com/vapor-ware/synse-server-grpc/go"
	"google.golang.org/grpc"
)

// NewFakeConn returns a gRPC client connection that can be used for testing,
// since the mock will need to return a connection that gets closed.
func NewFakeConn() *grpc.ClientConn {
	conn, err := grpc.Dial("localhost:5555", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return conn
}

// FakeGRPCClientV3 implements the synse.V3PluginClient interface to allow
// for basic testing of gRPC-based commands without dependencies on external
// services.
type FakeGRPCClientV3 struct {
	cmdErr bool
}

// NewFakeGRPCClientV3 creates a new fake gRPC client whose methods will
// never return an error.
func NewFakeGRPCClientV3() synse.V3PluginClient {
	return &FakeGRPCClientV3{}
}

// NewFakeGRPCClientV3Err creates a new fake gRPC client whose methods will
// always return an error.
func NewFakeGRPCClientV3Err() synse.V3PluginClient {
	return &FakeGRPCClientV3{
		cmdErr: true,
	}
}

func (f *FakeGRPCClientV3) Devices(ctx context.Context, in *synse.V3DeviceSelector, opts ...grpc.CallOption) (synse.V3Plugin_DevicesClient, error) {
	if f.cmdErr {
		return nil, ErrFakeClient
	}
	return nil, nil
}

func (f *FakeGRPCClientV3) Health(ctx context.Context, in *synse.Empty, opts ...grpc.CallOption) (*synse.V3Health, error) {
	if f.cmdErr {
		return nil, ErrFakeClient
	}
	return &synse.V3Health{
		Status:    synse.HealthStatus_OK,
		Timestamp: "2019-04-22T13:30:00Z",
		Checks: []*synse.V3HealthCheck{
			{
				Name:      "test-check",
				Status:    synse.HealthStatus_OK,
				Timestamp: "2019-04-22T13:30:00Z",
				Type:      "periodic",
			},
		},
	}, nil
}

func (f *FakeGRPCClientV3) Metadata(ctx context.Context, in *synse.Empty, opts ...grpc.CallOption) (*synse.V3Metadata, error) {
	if f.cmdErr {
		return nil, ErrFakeClient
	}
	return &synse.V3Metadata{
		Name:        "test-plugin",
		Maintainer:  "vaporio",
		Tag:         "vaporio/test-plugin",
		Description: "a plugin",
		Id:          "987654",
	}, nil
}

func (f *FakeGRPCClientV3) Read(ctx context.Context, in *synse.V3ReadRequest, opts ...grpc.CallOption) (synse.V3Plugin_ReadClient, error) {
	if f.cmdErr {
		return nil, ErrFakeClient
	}
	return nil, nil
}

func (f *FakeGRPCClientV3) ReadCache(ctx context.Context, in *synse.V3Bounds, opts ...grpc.CallOption) (synse.V3Plugin_ReadCacheClient, error) {
	if f.cmdErr {
		return nil, ErrFakeClient
	}
	return nil, nil
}

func (f *FakeGRPCClientV3) Test(ctx context.Context, in *synse.Empty, opts ...grpc.CallOption) (*synse.V3TestStatus, error) {
	if f.cmdErr {
		return nil, ErrFakeClient
	}
	return &synse.V3TestStatus{
		Ok: true,
	}, nil
}

func (f *FakeGRPCClientV3) Transaction(ctx context.Context, in *synse.V3TransactionSelector, opts ...grpc.CallOption) (*synse.V3TransactionStatus, error) {
	if f.cmdErr {
		return nil, ErrFakeClient
	}
	return &synse.V3TransactionStatus{
		Id:      "123456",
		Created: "2019-04-22T13:30:00Z",
		Updated: "2019-04-22T13:30:00Z",
		Timeout: "30s",
		Status:  synse.WriteStatus_DONE,
	}, nil
}

func (f *FakeGRPCClientV3) Transactions(ctx context.Context, in *synse.Empty, opts ...grpc.CallOption) (synse.V3Plugin_TransactionsClient, error) {
	if f.cmdErr {
		return nil, ErrFakeClient
	}
	return nil, nil
}

func (f *FakeGRPCClientV3) Version(ctx context.Context, in *synse.Empty, opts ...grpc.CallOption) (*synse.V3Version, error) {
	if f.cmdErr {
		return nil, ErrFakeClient
	}
	return &synse.V3Version{
		PluginVersion: "3.2.1",
		SdkVersion:    "3.0.0",
	}, nil
}

func (f *FakeGRPCClientV3) WriteAsync(ctx context.Context, in *synse.V3WritePayload, opts ...grpc.CallOption) (synse.V3Plugin_WriteAsyncClient, error) {
	if f.cmdErr {
		return nil, ErrFakeClient
	}
	return nil, nil
}

func (f *FakeGRPCClientV3) WriteSync(ctx context.Context, in *synse.V3WritePayload, opts ...grpc.CallOption) (synse.V3Plugin_WriteSyncClient, error) {
	if f.cmdErr {
		return nil, ErrFakeClient
	}
	return nil, nil
}
