package main

import (
	pb "IPC_types_gRPC/client_streaming/client/pkg/api"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const address = ":50053"

func main() {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer func(conn *grpc.ClientConn) {
		err = conn.Close()
		if err != nil {
			log.Fatalf("could not close connection: %v", err)
		}
	}(conn)

	c := pb.NewUpdateOrderServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	updateStream, err := c.UpdateOrders(ctx)
	if err != nil {
		log.Fatalf("%v.UpdateOrders(_) = _, %v", c, err)
	}

	if err = updateStream.Send(&pb.Order{
		Id:    "1",
		Price: 3.6,
	}); err != nil {
		log.Fatalf("%v.Send(%v) = %v", updateStream, &pb.Order{
			Id:    "1",
			Price: 3.6,
		}, err)
	}

	if err = updateStream.Send(&pb.Order{
		Id:    "2",
		Price: 3.6,
	}); err != nil {
		log.Fatalf("%v.Send(%v) = %v", updateStream, &pb.Order{
			Id:    "2",
			Price: 3.6,
		}, err)
	}

	if err = updateStream.Send(&pb.Order{
		Id:    "3",
		Price: 3.6,
	}); err != nil {
		log.Fatalf("%v.Send(%v) = %v", updateStream, &pb.Order{
			Id:    "3",
			Price: 3.6,
		}, err)
	}

	if err = updateStream.Send(&pb.Order{
		Id:    "4",
		Price: 3.6,
	}); err != nil {
		log.Fatalf("%v.Send(%v) = %v", updateStream, &pb.Order{
			Id:    "4",
			Price: 3.6,
		}, err)
	}

	updateRes, err := updateStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv(): %v", updateStream, err)
	}

	log.Println(updateRes)
}
