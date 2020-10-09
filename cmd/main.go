package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/evilmonkeyinc/account-manager/gen/Openapi"
	"github.com/evilmonkeyinc/account-manager/pkg/service"
	"github.com/go-chi/chi"
)

const (
	grpcServerEndpoint string = "localhost:9090"
	httpServerEndpoint string = ":8081"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// gnostic-go-generator
	// openapi.Initialize(service.New())
	// err := openapi.ServeHTTP(httpServerEndpoint)
	// if err != nil {
	// 	panic(err)
	// }

	// Chi and oapi
	r := chi.NewRouter()
	r.Mount("/api", Openapi.Handler(service.New()))
	http.ListenAndServe(httpServerEndpoint, r)

	fmt.Println("service is stopping")
}
