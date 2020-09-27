package server

import (
	"context"
	"fmt"
	"net"

	"github.com/evilmonkeyinc/account-manager/gen/openapi"
	"google.golang.org/grpc"
)

func RunServer(serverEndpoint string) error {
	listener, err := net.Listen("tcp", serverEndpoint)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	openapi.RegisterOpenapiServer(server, &impl{})

	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}

type impl struct {
}

func (server *impl) GetPlugin(ctx context.Context, params *openapi.GetPluginParameters) (*openapi.Plugin, error) {

	if params.Name == "default" {
		return &openapi.Plugin{
				Name: fmt.Sprintf("plugin-%s", params.Name),
			},
			nil
	}

	return nil, fmt.Errorf("failed to find plugin %s", params.Name)
}

func (server *impl) ListPlugins(ctx context.Context, request *openapi.ListPluginsParameters) (*openapi.ListPluginsOK, error) {

	response := &openapi.ListPluginsOK{
		Count: 0,
		Limit: request.Limit,
		Page:  request.Page,
		Total: 0,
	}

	return response, nil
}
