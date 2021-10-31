package client

import (
	"io"
	"log"
	"time"

	"github.com/dimitarsi/hello-grpc/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func RunClient() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Cannot connet to localhost:5000")
	}

	defer conn.Close()

	orderManagementClient := service.NewOrderManagementClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), time.Second * 3)

	stream, err := orderManagementClient.SearchOrder(ctx, &wrapperspb.StringValue{Value:"ticket"} )

	if err != nil {
		log.Fatalf("Error seraching for an order %v", err)
	}

	for {
		order, err := stream.Recv()

		if err == io.EOF {
			break;
		}

		log.Printf("Order: %v", order)
	}

}