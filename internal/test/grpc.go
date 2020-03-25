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

package test

import (
	"context"
	"io"

	synse "github.com/vapor-ware/synse-server-grpc/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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

type FakeClientStream struct {
	md         metadata.MD
	isComplete bool
}

func (f *FakeClientStream) Header() (metadata.MD, error) {
	return f.md, nil
}

func (f *FakeClientStream) Trailer() metadata.MD {
	return f.md
}

func (f *FakeClientStream) CloseSend() error {
	return nil
}

func (f *FakeClientStream) Context() context.Context {
	return context.Background()
}

func (f *FakeClientStream) SendMsg(m interface{}) error {
	return io.EOF
}

func (f *FakeClientStream) RecvMsg(m interface{}) error {
	return io.EOF
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

type FakeDevicesClient struct {
	FakeClientStream
}

func (f *FakeDevicesClient) Recv() (*synse.V3Device, error) {

	if f.isComplete {
		return nil, io.EOF
	}

	f.isComplete = true
	return &synse.V3Device{
		Timestamp: "2019-04-22T13:30:00Z",
		Id:        "111-222-333",
		Type:      "testdevice",
		Plugin:    "123345",
		Info:      "foo bar",
		Alias:     "faked",
		Tags: []*synse.V3Tag{
			{
				Namespace:  "fake",
				Annotation: "test",
				Label:      "device",
			},
		},
		Metadata: map[string]string{
			"123": "456",
			"abc": "def",
		},
	}, nil
}

func (f *FakeGRPCClientV3) Devices(ctx context.Context, in *synse.V3DeviceSelector, opts ...grpc.CallOption) (synse.V3Plugin_DevicesClient, error) {
	if f.cmdErr {
		return nil, ErrFakeClient
	}
	return &FakeDevicesClient{
		FakeClientStream: FakeClientStream{
			md: metadata.MD{},
		},
	}, nil
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

type FakeReadClient struct {
	FakeClientStream
}

func (f *FakeReadClient) Recv() (*synse.V3Reading, error) {

	if f.isComplete {
		return nil, io.EOF
	}

	f.isComplete = true
	return &synse.V3Reading{
		Id:         "123",
		Timestamp:  "2019-04-22T13:30:00Z",
		Type:       "faked",
		DeviceType: "faked",
		Context: map[string]string{
			"foo": "bar",
		},
		Unit: &synse.V3OutputUnit{},
		Value: &synse.V3Reading_Float64Value{
			Float64Value: 23.0,
		},
	}, nil
}

func (f *FakeGRPCClientV3) Read(ctx context.Context, in *synse.V3ReadRequest, opts ...grpc.CallOption) (synse.V3Plugin_ReadClient, error) {
	if f.cmdErr {
		return nil, ErrFakeClient
	}
	return &FakeReadClient{
		FakeClientStream: FakeClientStream{
			md: metadata.MD{},
		},
	}, nil
}

func (f *FakeGRPCClientV3) ReadCache(ctx context.Context, in *synse.V3Bounds, opts ...grpc.CallOption) (synse.V3Plugin_ReadCacheClient, error) {
	if f.cmdErr {
		return nil, ErrFakeClient
	}
	return &FakeReadClient{
		FakeClientStream: FakeClientStream{
			md: metadata.MD{},
		},
	}, nil
}

func (f *FakeGRPCClientV3) ReadStream(ctx context.Context, in *synse.V3StreamRequest, opts ...grpc.CallOption) (synse.V3Plugin_ReadStreamClient, error) {
	if f.cmdErr {
		return nil, ErrFakeClient
	}
	return &FakeReadClient{
		FakeClientStream: FakeClientStream{
			md: metadata.MD{},
		},
	}, nil
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

type FakeTransactionsClient struct {
	FakeClientStream
}

func (f *FakeTransactionsClient) Recv() (*synse.V3TransactionStatus, error) {
	if f.isComplete {
		return nil, io.EOF
	}

	f.isComplete = true
	return &synse.V3TransactionStatus{
		Id:      "123",
		Created: "2019-04-22T13:30:00Z",
		Updated: "2019-04-22T13:30:00Z",
		Timeout: "30s",
		Status:  synse.WriteStatus_DONE,
		Context: &synse.V3WriteData{
			Action: "foo",
			Data:   []byte("bar"),
		},
	}, nil
}

func (f *FakeGRPCClientV3) Transactions(ctx context.Context, in *synse.Empty, opts ...grpc.CallOption) (synse.V3Plugin_TransactionsClient, error) {
	if f.cmdErr {
		return nil, ErrFakeClient
	}
	return &FakeTransactionsClient{
		FakeClientStream: FakeClientStream{
			md: metadata.MD{},
		},
	}, nil
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

type FakeWriteClient struct {
	FakeClientStream
}

func (f *FakeWriteClient) Recv() (*synse.V3WriteTransaction, error) {
	if f.isComplete {
		return nil, io.EOF
	}

	f.isComplete = true
	return &synse.V3WriteTransaction{
		Id:      "123456",
		Device:  "987654",
		Timeout: "30s",
		Context: &synse.V3WriteData{
			Action: "foo",
			Data:   []byte("bar"),
		},
	}, nil
}

func (f *FakeGRPCClientV3) WriteAsync(ctx context.Context, in *synse.V3WritePayload, opts ...grpc.CallOption) (synse.V3Plugin_WriteAsyncClient, error) {
	if f.cmdErr {
		return nil, ErrFakeClient
	}
	return &FakeWriteClient{
		FakeClientStream: FakeClientStream{
			md: metadata.MD{},
		},
	}, nil
}

func (f *FakeGRPCClientV3) WriteSync(ctx context.Context, in *synse.V3WritePayload, opts ...grpc.CallOption) (synse.V3Plugin_WriteSyncClient, error) {
	if f.cmdErr {
		return nil, ErrFakeClient
	}
	return &FakeTransactionsClient{
		FakeClientStream: FakeClientStream{
			md: metadata.MD{},
		},
	}, nil
}
