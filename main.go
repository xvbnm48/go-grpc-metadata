package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	pb "grpc-header/proto"
	"log"
	"net"
)

type server struct {
}

func (s *server) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata from context")
	}

	response := &pb.RegisterResponse{}
	if phoneNumbers := md["nomor-hp"]; len(phoneNumbers) > 0 {
		response.Data = phoneNumbers[0]
	}
	if NIP := md["nip"]; len(NIP) > 0 {
		response.Data = NIP[0]
	}
	fmt.Println(response.Data)
	response.Message = "Success"

	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	fmt.Println("Server running on port :8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterRegisterDataServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
