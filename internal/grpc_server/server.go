package grpc_server

import (
	"context"
	pb "go-grpc-demo/api/proto/order"
	"go-grpc-demo/internal/service"
)

type OrderGRPCServer struct {
	pb.UnimplementedOrderServiceServer
	orderService *service.OrderService
}

func NewOrderGRPCServer(orderService *service.OrderService) *OrderGRPCServer {
	return &OrderGRPCServer{orderService: orderService}
}

func (s *OrderGRPCServer) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	order, err := s.orderService.CreateOrder(req.UserId, req.ProductId, req.Amount)
	if err != nil {
		return &pb.CreateOrderResponse{Status: "FAIL", Message: err.Error()}, nil
	}
	return &pb.CreateOrderResponse{
		OrderId: order.ID,
		Status:  "SUCCESS",
		Message: "Order created successfully",
	}, nil
}

func (s *OrderGRPCServer) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	order, err := s.orderService.GetOrderByID(req.OrderId)
	if err != nil {
		return &pb.GetOrderResponse{Status: "FAIL"}, nil
	}
	return &pb.GetOrderResponse{
		OrderId:        order.ID,
		TotalAmountFen: order.TotalAmountFen,
		Status:         "SUCCESS",
	}, nil
}
