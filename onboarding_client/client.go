package main

import (
    "context"
    "log"
    "time"

    pb "tproto/proto" // Adjust this import path as per your actual setup

    "google.golang.org/grpc"
)

func main() {
    // Establish a connection to the gRPC server
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Failed to dial server: %v", err)
    }
    defer conn.Close()

    // Create a gRPC client
    client := pb.NewOnboardingServiceClient(conn)

    // Example RPC call: Replace with your actual RPC method
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Replace with your actual RPC method
    // Example: Call the GiveConsent RPC method
    resp, err := client.GiveConsent(ctx, &pb.ConsentRequest{
        UserId:           "user123",
        TermsOfService:   "Agreed to terms",
        ConsentGiven:     true,
    })
    if err != nil {
        log.Fatalf("Error calling GiveConsent RPC: %v", err)
    }

    log.Printf("GiveConsent response: %v", resp)
}
