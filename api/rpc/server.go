package rpc

import (
	"context"
	"fmt"
	"net"

	"github.com/9ssi7/webstack/pkg/server"
	"google.golang.org/grpc"
)

type srv struct {
	server *grpc.Server
}

func NewServer() server.Listener {
	server := grpc.NewServer()

	return &srv{
		server: server,
	}
}

func (s *srv) Listen() error {
	// Register your gRPC services here
	return nil
}

func (s *srv) Shutdown(ctx context.Context) error {
	s.server.GracefulStop()
	return nil
}

func (s *srv) Start(port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	return s.server.Serve(lis)
}
