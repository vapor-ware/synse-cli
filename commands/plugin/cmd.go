package plugin

import (
	"fmt"
	"net"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/flags"
	"github.com/vapor-ware/synse-server-grpc/go"
	"google.golang.org/grpc"
)

// PluginCommand is the CLI command for interacting with Synse plugins.
var PluginCommand = cli.Command{
	Name:  "plugin",
	Usage: "Interact with Synse Plugins",
	Flags: []cli.Flag{
		flags.HostFlag,
		flags.UnixFlag,
	},
	Subcommands: []cli.Command{
		pluginMetainfoCommand,
		pluginReadCommand,
		pluginTransactionCommand,
		pluginWriteCommand,
	},
}

func makeGrpcClient(c *cli.Context) (synse.InternalApiClient, error) {
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
