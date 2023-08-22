package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	pb "grpc-header/proto"
)

func (s *Server) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
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
