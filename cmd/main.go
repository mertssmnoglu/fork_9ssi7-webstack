package main

import (
	"log"
	"sync"

	"github.com/9ssi7/webstack/api/rest"
	"github.com/9ssi7/webstack/api/rpc"
	"github.com/9ssi7/webstack/api/web"
	"github.com/9ssi7/webstack/pkg/server"
)

func main() {
	http := server.NewXHTTP(server.XHTTPConfig{
		Host: "0.0.0.0",
		Port: 8080,
	}, rest.New(), web.New())

	// Create RPC server
	rpcServer := rpc.NewServer()

	// Start both servers concurrently
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := http.Listen(); err != nil {
			log.Fatalf("Failed to start API server: %v", err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := rpcServer.Listen(); err != nil {
			log.Fatalf("Failed to start RPC server: %v", err)
		}
	}()

	wg.Wait()
}
