package main

import (
	pb "IPC_types_gRPC/unary/client/pkg/ecommerce/api"
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const address = ":50051"

func main() {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("client connection error: %+v", err)
	}

	defer func(conn *grpc.ClientConn) {
		err = conn.Close()
		if err != nil {
			log.Fatalf("close connection error: %+v", err)
		}
	}(conn)

	c := pb.NewOrderManagementClient(conn)

	orderId := "first"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.GetOrder(ctx, &wrappers.StringValue{Value: orderId})
	if err != nil {
		log.Fatalf("get order error: %+v", err)
	}

	log.Println("Success!!!")
	log.Printf("get order response: %+v", res.String())
}
