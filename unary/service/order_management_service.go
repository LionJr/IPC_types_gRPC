package service

import (
	pb "IPC_types_gRPC/unary/service/pkg/ecommerce/api"
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
)

type OrderManagementServer struct {
	pb.UnimplementedOrderManagementServer
	orderMap map[string]*pb.Order
}

func (o *OrderManagementServer) GetOrder(ctx context.Context, orderId *wrappers.StringValue) (*pb.Order, error) {
	order := o.orderMap[orderId.Value]
	return order, nil
}
