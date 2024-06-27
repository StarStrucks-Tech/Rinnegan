package main

import (
	"context"
	"log"
	"net"

	pb "Nishikant/uiproject/proto" // Adjust this path based on your actual directory structure and package name

	"google.golang.org/grpc"
)

// server is used to implement UiServiceServer.
type server struct {
	pb.UnimplementedUiServiceServer
}

// GetUiResponse implements UiServiceServer.GetUiResponse
func (s *server) GetUiResponse(ctx context.Context, req *pb.UiRequest) (*pb.UiResponse, error) {
	return &pb.UiResponse{
		ResponseId: "123",
		Message:    "User profile received successfully",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUiServiceServer(s, &server{})
	log.Println("Server is running on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
