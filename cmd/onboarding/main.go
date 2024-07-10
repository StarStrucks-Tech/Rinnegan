package main

import (
	pb "Rinnegan/proto-generated/onboarding"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type onboardingServer struct {
	pb.UnimplementedOnboardingServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	// server registration
	pb.RegisterOnboardingServiceServer(grpcServer, &onboardingServer{})
	// for successful registration on the port , this will be logged , else can be logged according to the next check
	fmt.Println("Server is running on port 8082...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
