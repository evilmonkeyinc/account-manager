package grpc

import (
	"context"
	"fmt"
	"net"
	"sync"

	"github.com/evilmonkeyinc/account-manager/gen/openapi"
	"google.golang.org/grpc"
)

// Run will start the GRPC server on the specified server endpoint
func Run(ctx context.Context, waitGroup *sync.WaitGroup, grpcServerEndpoint string) error {
	defer waitGroup.Done()

	listener, err := net.Listen("tcp", grpcServerEndpoint)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	openapi.RegisterOpenapiServer(server, &impl{})

	return server.Serve(listener)
}

type impl struct {
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

func (server *impl) FetchPlugin(ctx context.Context, params *openapi.FetchPluginParameters) (*openapi.Plugin, error) {

	if params.Name == "default" {
		return &openapi.Plugin{
				Name: fmt.Sprintf("plugin-%s", params.Name),
			},
			nil
	}

	return nil, fmt.Errorf("failed to find plugin %s", params.Name)
}
