package main

import (
	"log"
	"net"

	"github.com/dimitarsi/hello-grpc/service"
	"google.golang.org/grpc"
)

const (
	port = ":5000"
)

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Unable to listen on %s", port)
	}

	s := grpc.NewServer()
	server := &Server{}

	service.RegisterProductInfoServer(s, server)

	log.Printf("Listening on port %s", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Unable to start server listening on port %s", port)
	}

}