package main

import (
	"context"

	pb "github.com/dimitarsi/hello-grpc/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

type OrderManagementServer struct {
	pb.UnimplementedOrderManagementServer
}

func (s *OrderManagementServer) GetOrder(_ context.Context, in *wrapperspb.StringValue) (*pb.Order, error) {
	return &pb.Order{}, status.New(codes.OK, "").Err()
}