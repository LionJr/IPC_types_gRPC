package service

import (
	pb "IPC_types_gRPC/bidirectional_streaming/service/pkg/api"
	"io"
	"log"
)

var (
	OrderMap       = make(map[string]*pb.Order)
	orderBatchSize = 3
)

type BidirectionalStreamingServer struct {
	pb.UnimplementedBidirectionalStreamingServiceServer
}

func (b *BidirectionalStreamingServer) ProcessOrders(stream pb.BidirectionalStreamingService_ProcessOrdersServer) error {

	batchMarker := 1
	combinedDeliveryMap := make(map[string]*pb.CombinedDelivery)
	for {
		orderId, err := stream.Recv()
		if err == io.EOF {
			log.Printf("EOF : %s", orderId)
			for _, del := range combinedDeliveryMap {
				if err = stream.Send(del); err != nil {
					return err
				}
			}
			return nil
		}
		if err != nil {
			log.Println(err)
			return err
		}

		ord := OrderMap[orderId.GetValue()]
		if ord != nil {
			destination := ord.Destination
			delivery, found := combinedDeliveryMap[destination]
			if !found {
				delivery = &pb.CombinedDelivery{
					Id:     "cmd - " + OrderMap[orderId.GetValue()].Destination,
					Status: "Processed!",
				}
			}

			delivery.OrdersList = append(delivery.OrdersList, ord)
			combinedDeliveryMap[destination] = delivery

			if batchMarker == orderBatchSize {
				for _, del := range combinedDeliveryMap {
					if err = stream.Send(del); err != nil {
						return err
					}
				}

				batchMarker = 0
				combinedDeliveryMap = make(map[string]*pb.CombinedDelivery)
			} else {
				batchMarker++
			}
		}
	}
}
