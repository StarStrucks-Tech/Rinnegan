package main

import (
	pb "Rinnegan/proto-generated/simulator"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type simulatorServer struct {
	pb.UnimplementedSimulatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":8084")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	// server registration
	pb.RegisterSimulatorServiceServer(grpcServer, &simulatorServer{})
	// for successful registration on the port , this will be logged , else can be logged according to the next check
	fmt.Println("Server is running on port 8084...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
