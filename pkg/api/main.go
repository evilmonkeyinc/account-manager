package main

import (
	"context"
	"flag"
	"net/http"

	"github.com/evilmonkeyinc/account-manager/gen/openapi"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

const grpcServerEndpoint string = "localhost:9090"
const httpServerEndpoint string = ":8081"

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := openapi.RegisterOpenapiHandlerServer(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(httpServerEndpoint, mux)
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		panic(err)
	}
}
