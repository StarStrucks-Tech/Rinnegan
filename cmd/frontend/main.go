package main

import (
	pb "Rinnegan/proto-generated/protos-frontend/frontend"
	generic "Rinnegan/proto-generated/protos-frontend/generic"
	frontendOnboarding "Rinnegan/proto-generated/protos-frontend/onboarding"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

// TODO : move this to the pkg package and use accordingly
type frontendServer struct {
	pb.UnimplementedFrontendServiceServer
}

func (s *frontendServer) GetEmailVerificationScreen(
	ctx context.Context,
	req *generic.EmptyRequest) (*frontendOnboarding.GetEmailVerificationScreenResponse, error) {
	fmt.Println("The rpc for email verification screen is used")
	// TODO implement this function for proper handling of the servers
	return &frontendOnboarding.GetEmailVerificationScreenResponse{
		Title:    "",
		SubTitle: "",
	}, nil
}

func (s *frontendServer) GetPhoneVerificationScreen(ctx context.Context, req *generic.EmptyRequest) (*frontendOnboarding.GetPhoneVerificationScreenResponse, error) {
	// TODO implement this function for proper handling of the servers
	return &frontendOnboarding.GetPhoneVerificationScreenResponse{
		Title: "",
	}, nil
}

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
	pb.RegisterFrontendServiceServer(grpcServer, &frontendServer{})
	// for successful registration on the port , this will be logged , else can be logged according to the next check
	fmt.Println("Server is running on port 8080...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
