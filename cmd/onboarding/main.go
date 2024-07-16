package main

import (
	pb "Rinnegan/proto-generated/onboarding"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type onboardingServer struct {
	pb.UnimplementedOnboardingServiceServer
}

func (s *onboardingServer) SendVerificationCode(ctx context.Context, req *pb.EmailVerificationRequest) (*pb.EmailVerificationResponse, error) {
	// TODO: Implement email sending logic here.
	log.Printf("Sending verification code to email: %s", req.Email)
	return &pb.EmailVerificationResponse{Success: true, Message: "Verification code sent"}, nil
}

// VerifyCode verifies the code sent to the user's email.
func (s *onboardingServer) VerifyCode(ctx context.Context, req *pb.VerifyCodeRequest) (*pb.VerifyCodeResponse, error) {
	// TODO: Implement code verification logic here.
	log.Printf("Verifying code for email: %s", req.Email)
	return &pb.VerifyCodeResponse{Success: true, Message: "Code verified"}, nil
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
