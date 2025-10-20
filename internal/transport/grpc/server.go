package grpc

import (
	"fmt"
	"net"

	userpb "github.com/Wendiboy/project-protos/proto/user"
	"github.com/Wendiboy/users-service/internal/user"
	"google.golang.org/grpc"
)

func RunGRPC(svc user.Service) error {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	userpb.RegisterUserServiceServer(grpcServer, NewHandler(svc))

	if err := grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("gRPC server failed: %v", err)
	}

	return nil
}
