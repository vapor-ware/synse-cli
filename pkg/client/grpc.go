package client

import (
	"fmt"
	"time"
	"net"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-server-grpc/go"
	"google.golang.org/grpc"
)

// MakeGrpcClient creates an instance of the Synse InternalApiClient for gRPC
// communication with plugins.
func MakeGrpcClient(c *cli.Context) (synse.InternalApiClient, error) {
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

	client := synse.NewInternalApiClient(grpcConn)
	return client, nil
}