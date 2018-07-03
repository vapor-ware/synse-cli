package client

import (
	"fmt"
	"io"
	"net"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-server-grpc/go"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Grpc is an instance of the grpcClient that is used by the CLI to make requests
// against a plugin via the Synse gRPC API.
var Grpc = grpcClient{}

// grpcClient is a client for making requests against the Synse Server gRPC API.
type grpcClient struct {
	apiClient synse.PluginClient
}

// Reset resets the grpcClient state. This is used primarily for testing.
func (client *grpcClient) Reset() {
	client.apiClient = nil
}

// newGrpcClient creates an instance of the Synse PluginClient for gRPC
// communication with plugins.
func (client *grpcClient) newGrpcClient(c *cli.Context) (synse.PluginClient, error) {
	t := c.GlobalString("tcp")
	s := c.GlobalString("unix")

	if (t == "" && s == "") || (t != "" && s != "") {
		return nil, fmt.Errorf("one of 'tcp' or 'unix' flag must be set to specify plugin")
	}

	var grpcConn *grpc.ClientConn
	var err error

	if t != "" {
		log.Debugf("dial with: %v\n", t)
		grpcConn, err = grpc.Dial(
			t,
			grpc.WithInsecure(),
		)
	}
	if s != "" {
		log.Debugf("dial with %v\n", s)
		grpcConn, err = grpc.Dial(
			s,
			grpc.WithInsecure(),
			grpc.WithDialer(func(addr string, timeout time.Duration) (net.Conn, error) {
				return net.DialTimeout("unix", addr, timeout)
			}),
		)
	}

	if err != nil {
		return nil, fmt.Errorf("unable to connect to plugin: %v", err)
	}
	return synse.NewPluginClient(grpcConn), nil
}

// Capabilities issues a "capabilities" request to a plugin via the gRPC API.
func (client *grpcClient) Capabilities(c *cli.Context) (out []*synse.DeviceCapability, err error) {
	if client.apiClient == nil {
		client.apiClient, err = client.newGrpcClient(c)
		if err != nil {
			return nil, err
		}
	}

	stream, err := client.apiClient.Capabilities(
		context.Background(),
		&synse.Empty{},
	)
	if err != nil {
		return nil, err
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		out = append(out, resp)
	}
	return out, nil
}

// Devices issues a "device" request to a plugin via the gRPC API.
func (client *grpcClient) Devices(c *cli.Context) (out []*synse.Device, err error) {
	if client.apiClient == nil {
		client.apiClient, err = client.newGrpcClient(c)
		if err != nil {
			return nil, err
		}
	}

	stream, err := client.apiClient.Devices(
		context.Background(),
		&synse.DeviceFilter{},
	)
	if err != nil {
		return nil, err
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		out = append(out, resp)
	}
	return out, nil
}

// Read issues a "read" request to a plugin via the gRPC API.
func (client *grpcClient) Read(c *cli.Context, rack, board, device string) (out []*synse.Reading, err error) {
	if client.apiClient == nil {
		client.apiClient, err = client.newGrpcClient(c)
		if err != nil {
			return nil, err
		}
	}

	stream, err := client.apiClient.Read(context.Background(), &synse.DeviceFilter{
		Rack:   rack,
		Board:  board,
		Device: device,
	})

	if err != nil {
		return nil, err
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		out = append(out, resp)
	}
	return out, nil
}

// Write issues a "write" request to a plugin via the gRPC API.
func (client *grpcClient) Write(c *cli.Context, rack, board, device string, data *synse.WriteData) (out *synse.Transactions, err error) {
	if client.apiClient == nil {
		client.apiClient, err = client.newGrpcClient(c)
		if err != nil {
			return nil, err
		}
	}

	transactions, err := client.apiClient.Write(context.Background(), &synse.WriteInfo{
		DeviceFilter: &synse.DeviceFilter{
			Rack:   rack,
			Board:  board,
			Device: device,
		},
		Data: []*synse.WriteData{data},
	})
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

// Transaction issues a "transaction" request to a plugin via the gRPC API.
func (client *grpcClient) Transaction(c *cli.Context, transactionID string) (out []*synse.WriteResponse, err error) {
	if client.apiClient == nil {
		client.apiClient, err = client.newGrpcClient(c)
		if err != nil {
			return nil, err
		}
	}

	stream, err := client.apiClient.Transaction(context.Background(), &synse.TransactionFilter{
		Id: transactionID,
	})
	if err != nil {
		return nil, err
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		out = append(out, resp)
	}
	return out, nil
}
