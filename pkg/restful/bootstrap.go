package restful

import (
	"context"
	"net/http"
	"sync"

	"github.com/evilmonkeyinc/account-manager/gen/openapi"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

// Run will start the RESTful server on the specified server endpoint linked to the GRPC server on the GRPC server endpoint.
func Run(ctx context.Context, waitGroup *sync.WaitGroup, grpcServerEndpoint, restfulSeverEndpoint string) error {
	defer waitGroup.Done()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := openapi.RegisterOpenapiHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(restfulSeverEndpoint, mux)
}
