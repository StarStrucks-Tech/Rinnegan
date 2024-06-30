package main

import (
	pb "Rinnegan/proto-generated/payments"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type paymentsServer struct {
	pb.UnimplementedPaymentsServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	// server registration
	pb.RegisterPaymentsServiceServer(grpcServer, &paymentsServer{})
	// for successful registration on the port , this will be logged , else can be logged according to the next check
	fmt.Println("Server is running on port 8083...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
