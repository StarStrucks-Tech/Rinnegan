package main

import (
	pb "checkerfolder/firstgen"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type bankServer struct {
	pb.UnimplementedBankServer
	accounts map[int32]float32
	nextID   int32
}

func (s *bankServer) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	s.nextID++
	s.accounts[s.nextID] = 0
	return &pb.CreateAccountResponse{
		AccountId: s.nextID,
		Message:   "Account created successfully",
	}, nil
}

func (s *bankServer) Deposit(ctx context.Context, req *pb.DepositRequest) (*pb.DepositResponse, error) {
	if balance, ok := s.accounts[req.AccountId]; ok {
		s.accounts[req.AccountId] = balance + req.Amount
		return &pb.DepositResponse{
			NewBalance: s.accounts[req.AccountId],
			Message:    "Deposit successful",
		}, nil
	}
	return &pb.DepositResponse{Message: "This Account does not found"}, nil
}

func (s *bankServer) Withdraw(ctx context.Context, req *pb.WithdrawRequest) (*pb.WithdrawResponse, error) {
	if balance, ok := s.accounts[req.AccountId]; ok {
		if balance >= req.Amount {
			s.accounts[req.AccountId] = balance - req.Amount
			return &pb.WithdrawResponse{
				NewBalance: s.accounts[req.AccountId],
				Message:    "Withdrawal successful",
			}, nil
		}
		return &pb.WithdrawResponse{Message: "Insufficient funds"}, nil
	}
	return &pb.WithdrawResponse{Message: "Account not found"}, nil
}

func (s *bankServer) GetBalance(ctx context.Context, req *pb.GetBalanceRequest) (*pb.GetBalanceResponse, error) {
	if balance, ok := s.accounts[req.AccountId]; ok {
		return &pb.GetBalanceResponse{
			Balance: balance,
			Message: "Balance retrieved successfully",
		}, nil
	}
	return &pb.GetBalanceResponse{Message: "Account not found"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8008")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBankServer(s, &bankServer{
		accounts: make(map[int32]float32),
		nextID:   0,
	})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
