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

	"github.com/vapor-ware/synse-cli/pkg/config"
	synse "github.com/vapor-ware/synse-server-grpc/go"
	"google.golang.org/grpc"
)

func NewSynseGrpcClient() (*grpc.ClientConn, synse.V3PluginClient, error) {
	// TODO: secure transport

	currentContexts := config.GetCurrentContext()
	pluginCtx := currentContexts["plugin"]
	if pluginCtx == nil {
		return nil, nil, fmt.Errorf("cannot create gRPC client for plugin: no current plugin context")
	}

	conn, err := grpc.Dial(pluginCtx.Context.Address, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}
	client := synse.NewV3PluginClient(conn)

	return conn, client, nil
}
