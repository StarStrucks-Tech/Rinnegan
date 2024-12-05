package main

import (
	"context"
	"log"
	"time"

	pb "checkerfolder/firstgen" // Replace with the actual path to your generated pb.go files

	"google.golang.org/grpc"
)

// Define client helper functions
func createAccount(client pb.BankClient) {
	req := &pb.CreateAccountRequest{Name: "John Doe"}
	res, err := client.CreateAccount(context.Background(), req)
	if err != nil {
		log.Fatalf("could not create account: %v", err)
	}
	log.Printf("Created account with ID: %d", res.AccountId)
}

func deposit(client pb.BankClient, accountId int32, amount float32) {
	req := &pb.DepositRequest{AccountId: accountId, Amount: amount}
	res, err := client.Deposit(context.Background(), req)
	if err != nil {
		log.Fatalf("could not deposit: %v", err)
	}
	log.Printf("New balance: %f", res.NewBalance)
}

func withdraw(client pb.BankClient, accountId int32, amount float32) {
	req := &pb.WithdrawRequest{AccountId: accountId, Amount: amount}
	res, err := client.Withdraw(context.Background(), req)
	if err != nil {
		log.Fatalf("could not withdraw: %v", err)
	}
	log.Printf("New balance: %f", res.NewBalance)
}

func getBalance(client pb.BankClient, accountId int32) {
	req := &pb.GetBalanceRequest{AccountId: accountId}
	res, err := client.GetBalance(context.Background(), req)
	if err != nil {
		log.Fatalf("could not get balance: %v", err)
	}
	log.Printf("Balance: %f", res.Balance)
}

// Main function to setup client and perform operations
func main() {
	// Wait a bit for the server to start
	// In a real-world scenario, use better synchronization
	<-time.After(time.Second * 1)

	// Setup the client connection
	conn, err := grpc.Dial("localhost:8008", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewBankClient(conn)

	// Perform some client operations
	createAccount(client)
	deposit(client, 1, 100.0)
	withdraw(client, 1, 50.0)
	getBalance(client, 1)
}
