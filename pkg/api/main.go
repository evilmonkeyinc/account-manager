package main

import (
	"context"
	"net/http"

	"github.com/evilmonkeyinc/account-manager/gen/openapi"
	"github.com/evilmonkeyinc/account-manager/pkg/server"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

const (
	grpcServerEndpoint string = "localhost:9090"
	httpServerEndpoint string = ":8081"
)

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go server.RunServer(grpcServerEndpoint)

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := openapi.RegisterOpenapiHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(httpServerEndpoint, mux)
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
