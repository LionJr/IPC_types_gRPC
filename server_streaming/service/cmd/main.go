package main

import (
	"IPC_types_gRPC/server_streaming/service"
	pb "IPC_types_gRPC/server_streaming/service/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const port = ":50052"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterOrderManagementServiceServer(grpcServer, &service.OrderManagementServerStreaming{})

	reflection.Register(grpcServer)

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
