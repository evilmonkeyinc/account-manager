package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/evilmonkeyinc/account-manager/pkg/grpc"
	"github.com/evilmonkeyinc/account-manager/pkg/restful"
)

const (
	grpcServerEndpoint string = "localhost:9090"
	httpServerEndpoint string = ":8081"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	channel := make(chan int)
	defer close(channel)

	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(2)

	fmt.Println("run grpc")
	go grpc.Run(ctx, waitGroup, grpcServerEndpoint)

	fmt.Println("run restful")
	go restful.Run(ctx, waitGroup, grpcServerEndpoint, httpServerEndpoint)

	fmt.Println("service is running")

	waitGroup.Wait()

	fmt.Println("service is stopping")
}
