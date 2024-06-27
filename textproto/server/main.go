package main

import (
    "context"
    "fmt"
    "log"
    "net"

    pb "Nishikant/yourproject/proto" // Import generated package

    "google.golang.org/grpc"
)

// server is used to implement proto.TextServiceServer.
type server struct {
    pb.UnimplementedTextServiceServer
}

// ProcessText implements proto.TextServiceServer
func (s *server) ProcessText(ctx context.Context, req *pb.TextRequest) (*pb.TextResponse, error) {
    // Here you can add any text processing logic
    processedText := fmt.Sprintf("Processed: %s", req.GetText())
    return &pb.TextResponse{ProcessedText: processedText}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterTextServiceServer(s, &server{})
    fmt.Println("Server is running on port :50051")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}