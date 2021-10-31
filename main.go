package main

import (
	"flag"
	"log"
	"net"
	"strings"

	"github.com/dimitarsi/hello-grpc/client"
	"github.com/dimitarsi/hello-grpc/service"
	"google.golang.org/grpc"
)

const (
	port = ":5000"
)

func main() {

	var clientArg string;

	flag.StringVar(&clientArg, "client", "", "Run as client")

	flag.Parse()

	if strings.Contains(clientArg, "true") {
		client.RunClient()

		return
	}

	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Unable to listen on %s", port)
	}


	s := grpc.NewServer()
	productInfoServer := &ProductInfoServer{}
	orderManagementServer := &OrderManagementServer{
		OrdersMap: map[string]service.Order {
			"1": {
				Id: "1", 
				Items: []string {"Iphone", "macBook", "toaster", "Foobar"},
			},
			"2": {
				Id: "2",
				Items: []string {"Pot", "flower", "chocolates", "Movie ticket"},
			},
		},
	}

	service.RegisterProductInfoServer(s, productInfoServer)
	service.RegisterOrderManagementServer(s, orderManagementServer)


	log.Printf("Listening on port %s", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Unable to start server listening on port %s", port)
	}

}