package main

import (
	pb "IPC_types_gRPC/bidirectional_streaming/client/pkg/api"
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"time"
)

const address = ":50054"

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

	c := pb.NewBidirectionalStreamingServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	streamProcOrder, err := c.ProcessOrders(ctx)
	if err != nil {
		log.Fatalf("could not process order: %v", err)
	}

	if err = streamProcOrder.Send(&wrappers.StringValue{Value: "102"}); err != nil {
		log.Fatalf("%v.Send(%v) = %v", c, "102", err)
	}

	if err = streamProcOrder.Send(&wrappers.StringValue{Value: "103"}); err != nil {
		log.Fatalf("%v.Send(%v) = %v", c, "103", err)
	}

	if err = streamProcOrder.Send(&wrappers.StringValue{Value: "104"}); err != nil {
		log.Fatalf("%v.Send(%v) = %v", c, "104", err)
	}

	channel := make(chan struct{})
	go asyncClientBidirectionalRPC(streamProcOrder, channel)

	if err = streamProcOrder.Send(&wrappers.StringValue{Value: "101"}); err != nil {
		log.Fatalf("%v.Send(%v) = %v", c, "101", err)
	}

	if err = streamProcOrder.CloseSend(); err != nil {
		log.Fatal(err)
	}

	<-channel
}

func asyncClientBidirectionalRPC(streamProcOrder pb.BidirectionalStreamingService_ProcessOrdersClient, c chan struct{}) {
	for {
		combinedDelivery, err := streamProcOrder.Recv()
		if err == io.EOF {
			break
		}

		if combinedDelivery != nil {
			log.Printf("Combined shipment : %v", combinedDelivery.OrdersList)
		}
	}
	c <- struct{}{}
}
