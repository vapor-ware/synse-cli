package client

import (
	"context"
	"github.com/vapor-ware/synse-cli/pkg/config"
	"github.com/vapor-ware/synse-cli/pkg/utils"
	"github.com/vapor-ware/synse-client-go/synse"
	sgrpc "github.com/vapor-ware/synse-server-grpc/go"
	"google.golang.org/grpc"
	"io"
)

type APIClient struct {
	ctx context.Context

	pluginClient sgrpc.V3PluginClient
	httpClient   synse.Client
	connection   *grpc.ClientConn
}

func NewAPIClient(ctx context.Context, flagContext, flagTLSCert string) (*APIClient, error) {
	hc, err := utils.NewSynseHTTPClient(flagContext, flagTLSCert)
	if err != nil {
		return nil, err
	}

	conn, pc, err := utils.NewSynseGrpcClient(flagContext, flagTLSCert)
	if err != nil {
		return nil, err
	}

	return &APIClient{
		ctx: ctx,

		pluginClient: pc,
		httpClient:   hc,
		connection:   conn,
	}, nil
}

func (a *APIClient) Close() {
	a.connection.Close()
}

func (a *APIClient) ServerVersion() string {
	resp, err := a.httpClient.Version()
	if err != nil {
		return ""
	}
	return resp.Version
}

func (a *APIClient) Devices() ([]*sgrpc.V3Device, error) {
	var devices []*sgrpc.V3Device
	deviceStream, err := a.pluginClient.Devices(a.ctx, &sgrpc.V3DeviceSelector{})
	if err != nil {
		return devices, err
	}

	for {
		resp, err := deviceStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return devices, err
		}
		devices = append(devices, resp)
	}
	return devices, nil
}

func (a *APIClient) Readings() ([]*sgrpc.V3Reading, error) {
	var readings []*sgrpc.V3Reading
	devices, err := a.Devices()
	if err != nil {
		return readings, err
	}
	for _, d := range devices {
		stream, err := a.pluginClient.Read(a.ctx, &sgrpc.V3ReadRequest{
			Selector: &sgrpc.V3DeviceSelector{
				Id: d.Id,
			},
		})
		if err != nil {
			return readings, err
		}
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				return readings, nil
			}
			readings = append(readings, resp)
		}
	}
	return readings, nil
}

func (a *APIClient) Context() []config.ContextRecord {
	return config.GetContexts()
}

func (a *APIClient) PluginVersion() string {
	version, err := a.pluginClient.Version(a.ctx, &sgrpc.Empty{})
	if err != nil {
		return ""
	}

	return version.PluginVersion
}

// TODO: Implement me
func (a *APIClient) SwitchContext() {}
