package main

import (
	"context"

	pb "github.com/dimitarsi/hello-grpc/service"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


type ProductInfoServer struct {
	pb.UnimplementedProductInfoServer
	
	productMap map[string]*pb.Product
}

func (s *ProductInfoServer) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {
	out, err := uuid.NewUUID()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to generate UUID", err)
	}

	in.Id = out.String()

	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}

	s.productMap[in.Id] = in;

	return &pb.ProductID{ Value: in.Id}, status.New(codes.OK, "").Err()
}

func (s *ProductInfoServer) GetProduct(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {

	if s.productMap == nil || s.productMap[in.Value] == nil {
		return nil, status.Errorf(codes.NotFound, "Product not found", in)
	}

	return s.productMap[in.Value], status.New(codes.OK, "").Err()
}