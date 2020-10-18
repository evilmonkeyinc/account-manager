package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/evilmonkeyinc/account-manager/pkg/middleware"
	"github.com/evilmonkeyinc/account-manager/pkg/service"
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
)

const (
	grpcServerEndpoint string = "localhost:9090"
	httpServerEndpoint string = ":8081"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	router := chi.NewRouter()
	router.Use(chimiddleware.RequestID, chimiddleware.Logger, middleware.ErrorWrapper)
	router.Mount("/api/v1/", service.New())
	http.ListenAndServe(httpServerEndpoint, router)

	fmt.Println("service is stopping")
}
