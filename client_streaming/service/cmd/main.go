package main

import (
	"IPC_types_gRPC/client_streaming/service"
	pb "IPC_types_gRPC/client_streaming/service/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const port = ":50053"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pb.RegisterUpdateOrderServiceServer(grpcServer, &service.UpdateOrderServiceServer{})

	log.Printf("server listening at %v", lis.Addr())
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
