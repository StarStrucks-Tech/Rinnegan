package main

import (
	"context"
	"log"
	"time"

	pb "Nishikant/uiproject/proto" // Adjust this path based on your actual directory structure and package name

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUiServiceClient(conn)

	// Prepare a request
	req := &pb.UiRequest{
		RequestId: "req-123",
		UserProfile: &pb.UserProfile{
			User: &pb.User{
				Id:     "user-1",
				Name:   "John Doe",
				Age:    30,
				Emails: []string{"john.doe@example.com", "john@example.com"},
			},
			Address: &pb.Address{
				Street:  "123 Main St",
				City:    "Anytown",
				Country: "USA",
			},
		},
	}

	// Set a timeout for the request
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Make the request using the correct method name from the .proto file
	res, err := client.GetUiResponse(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			log.Fatalf("gRPC error: %v - %v", st.Code(), st.Message())
		} else {
			log.Fatalf("Unknown error: %v", err)
		}
	}

	// Handle the response
	log.Printf("Response ID: %s", res.GetResponseId())
	log.Printf("Message: %s", res.GetMessage())
}
