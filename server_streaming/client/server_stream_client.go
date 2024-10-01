package main

import (
	pb "IPC_types_gRPC/server_streaming/client/pkg/api"
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"time"
)

const address = ":50052"

func main() {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("client connection error: %v", err)
	}

	defer func(conn *grpc.ClientConn) {
		err = conn.Close()
		if err != nil {
			log.Fatalf("client connection error: %v", err)
		}
	}(conn)

	c := pb.NewOrderManagementServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	searchStream, _ := c.SearchOrders(ctx, &wrappers.StringValue{Value: "1"})

	for {
		searchOrder, receiveErr := searchStream.Recv()
		if receiveErr != nil {
			if receiveErr == io.EOF {
				break
			}
			log.Fatal("error receiving order: ", receiveErr)
		}

		log.Printf("search order: %v", searchOrder.String())
	}
}
