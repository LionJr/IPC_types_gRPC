package service

import (
	pb "IPC_types_gRPC/server_streaming/service/pkg/api"
	"errors"
	"github.com/golang/protobuf/ptypes/wrappers"
	"log"
	"strconv"
)

var ordersMap map[string]*pb.Order

func init() {
	ordersMap = make(map[string]*pb.Order)
	for i := range 10 {
		key := strconv.Itoa(i)
		ordersMap[key] = &pb.Order{
			Id: key,
			Items: []string{
				key,
			},
			Description: "Test-description",
			Price:       float32(i),
			Destination: "Ashgabat",
		}
	}
}

type OrderManagementServerStreaming struct {
	pb.UnimplementedOrderManagementServiceServer
}

func (s *OrderManagementServerStreaming) SearchOrders(searchQuery *wrappers.StringValue, stream pb.OrderManagementService_SearchOrdersServer) error {

	for key, order := range ordersMap {
		log.Println(key, order)
		for _, orderItem := range order.Items {
			item, _ := strconv.Atoi(orderItem)
			if item%2 == 0 {
				err := stream.Send(order)
				if err != nil {
					return errors.New("error sending message to stream: " + err.Error())
				}

				log.Print("Matching Order Found : " + key)
				break
			}
		}
	}

	return nil
}
