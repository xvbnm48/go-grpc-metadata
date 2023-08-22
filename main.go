package main

import (
	"fmt"
	"google.golang.org/grpc"
	pb "grpc-header/proto"
	"grpc-header/server"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	fmt.Println("Server running on port :8080")
	if err != nil {
		panic(err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register the RegisterData server
	pb.RegisterRegisterDataServer(s, &server.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
