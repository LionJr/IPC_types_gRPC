package service

import (
	pb "IPC_types_gRPC/client_streaming/service/pkg/api"
	"github.com/golang/protobuf/ptypes/wrappers"
	"io"
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
			Description: "test-description",
			Price:       float32(i),
			Destination: "Mary",
		}
	}
}

type UpdateOrderServiceServer struct {
	pb.UnimplementedUpdateOrderServiceServer
}

func (u *UpdateOrderServiceServer) UpdateOrders(stream pb.UpdateOrderService_UpdateOrdersServer) error {
	orderStr := "Updated Order IDs : "
	for {
		order, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return stream.SendAndClose(&wrappers.StringValue{
					Value: "Orders processed " + orderStr,
				})
			}
		}

		ordersMap[order.Id] = order

		log.Println("Order ID ", order.Id, " Updated")
		orderStr += order.GetId() + " "
	}
}
