package main

import (
	"Rinnegan/pkg/rpc/frontend"
	pb "Rinnegan/proto-generated/protos-frontend/frontend"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

/*
*

	This is the main function where we have handled the port listening for
	the server and also the connection protocol is mentioned here
	Frontend Server will be listening to port 8080
*/
func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	// server registration
	pb.RegisterFrontendServiceServer(grpcServer, &pkg.Server{})
	// for successful registration on the port , this will be logged , else can be logged according to the next check
	fmt.Println("Server is running on port 8080...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
