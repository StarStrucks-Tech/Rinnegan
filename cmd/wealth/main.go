package main

import (
	pb "Rinnegan/proto-generated/wealth"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type wealthServer struct {
	pb.UnimplementedWealthServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":8085")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	// server registration
	pb.RegisterWealthServiceServer(grpcServer, &wealthServer{})
	// for successful registration on the port , this will be logged , else can be logged according to the next check
	fmt.Println("Server is running on port 8085...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
