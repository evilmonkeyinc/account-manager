package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	grpcServerEndpoint string = "localhost:9090"
	httpServerEndpoint string = ":8081"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	r := mux.NewRouter()
	r.HandleFunc("/ping", pingHandler)
	http.Handle("/", r)

	fmt.Println("service is stopping")
}

func pingHandler(http.ResponseWriter, *http.Request) {

}
