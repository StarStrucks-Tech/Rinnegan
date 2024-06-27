package frontend

import (
	pb "Rinnegan/proto-generated/protos-frontend/frontend"
)

// TODO : Use this client in the main.go for frontend
type frontendServer struct {
	pb.UnimplementedFrontendServiceServer
}
