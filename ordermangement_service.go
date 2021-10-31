package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"strings"

	pb "github.com/dimitarsi/hello-grpc/service"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

type OrderManagementServer struct {
	pb.UnimplementedOrderManagementServer

	OrdersMap map[string]pb.Order
}

func (s *OrderManagementServer) GetOrder(_ context.Context, in *wrapperspb.StringValue) (*pb.Order, error) {
	return &pb.Order{}, status.New(codes.OK, "").Err()
}

func (s *OrderManagementServer) SearchOrder(in *wrapperspb.StringValue, server pb.OrderManagement_SearchOrderServer) error {

	for _, order := range s.OrdersMap {
		if strings.Contains(order.Description, in.Value) ||
			findItemInOrder(order.Items, in.Value) {
			err := server.Send(&order)

			if err != nil {
				return fmt.Errorf("error sending to stream %w", err)
			}

			log.Printf("Matching order found - %s", order.Id)
		}
	}

	return nil
}

func (s *OrderManagementServer) UpdateOrders(stream pb.OrderManagement_UpdateOrdersServer) error {
	counter := 0
	for {
		order, err := stream.Recv()

		if err == io.EOF {
			msg := wrapperspb.StringValue{
				Value: fmt.Sprintf("Total items received: %d", counter),
			}

			return stream.SendAndClose(&msg)
		}

		counter += 1

		id := fmt.Sprintf("%d", uuid.New().ID())
		order.Id = id

		s.OrdersMap[id] = *order
	}
}

func findItemInOrder(items []string, needle string) bool {
	for _, item := range items {
		if strings.ContainsAny(item, needle) {
			return true
		}
	}

	return false
}
