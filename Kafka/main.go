package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}