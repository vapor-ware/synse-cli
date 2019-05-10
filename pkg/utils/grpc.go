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

package utils

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/vapor-ware/synse-cli/pkg/config"
	synse "github.com/vapor-ware/synse-server-grpc/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func NewSynseGrpcClient(ctx string, certFile string) (*grpc.ClientConn, synse.V3PluginClient, error) {
	var pluginContext *config.ContextRecord

	if ctx == "" {
		currentContexts := config.GetCurrentContext()
		pluginContext = currentContexts["plugin"]
		if pluginContext == nil {
			return nil, nil, fmt.Errorf("cannot create gRPC client for plugin: no current plugin context")
		}
	} else {
		pluginContext = config.GetContext(ctx)
		if pluginContext == nil {
			return nil, nil, fmt.Errorf("cannot create gRPC client for plugin: specified context does not exist")
		}
		if pluginContext.Type != "plugin" {
			return nil, nil, fmt.Errorf("cannot create gRPC client for plugin: specified context is not a plugin context")
		}
	}

	// If a cert file was provided, use it to override any cert file which may
	// already be configured for the context.
	if certFile != "" {
		pluginContext.Context.ClientCert = certFile
	}

	var dialOptions []grpc.DialOption
	if pluginContext.Context.ClientCert == "" {
		log.Debug("grpc client: with insecure")
		dialOptions = append(dialOptions, grpc.WithInsecure())
	} else {
		creds, err := credentials.NewClientTLSFromFile(pluginContext.Context.ClientCert, "")
		if err != nil {
			return nil, nil, err
		}
		log.WithField("cert", pluginContext.Context.ClientCert).Debug("grpc client: with credentials")
		dialOptions = append(dialOptions, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.Dial(pluginContext.Context.Address, dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	client := synse.NewV3PluginClient(conn)

	return conn, client, nil
}
