package main

import (
	pb "Rinnegan/proto-generated/home"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type homeServer struct {
	pb.UnimplementedHomeServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":8086")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	// server registration
	pb.RegisterHomeServiceServer(grpcServer, &homeServer{})
	// for successful registration on the port , this will be logged , else can be logged according to the next check
	fmt.Println("Server is running on port 8086...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
