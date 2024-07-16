package main

import (
    "context"
    "log"
    "net"
    "net/http"
    "sync"

    pb "tproto/proto"

    "google.golang.org/grpc"
)

type server struct {
    pb.UnimplementedOnboardingServiceServer
    consents map[string]bool
    mu       sync.Mutex
}

func (s *server) GiveConsent(ctx context.Context, req *pb.ConsentRequest) (*pb.ConsentResponse, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    if req.ConsentGiven {
        s.consents[req.UserId] = true
        return &pb.ConsentResponse{
            Success: true,
            Message: "Consent given",
        }, nil
    }

    return &pb.ConsentResponse{
        Success: false,
        Message: "Consent not given",
    }, nil
}

func (s *server) livelinessCheck(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Liveliness check passed"))
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    srv := &server{
        consents: make(map[string]bool),
    }
    pb.RegisterOnboardingServiceServer(s, srv)

    http.HandleFunc("/liveliness", srv.livelinessCheck)
    go func() {
        if err := http.ListenAndServe(":8080", nil); err != nil {
            log.Fatalf("failed to start liveliness check server: %v", err)
        }
    }()

    log.Println("Starting gRPC server on port 50051...")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
