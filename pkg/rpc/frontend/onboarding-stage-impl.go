package pkg

import (
	"Rinnegan/proto-generated/protos-frontend/enums"
	pb "Rinnegan/proto-generated/protos-frontend/frontend"
	"Rinnegan/proto-generated/protos-frontend/generic"
	"Rinnegan/proto-generated/protos-frontend/onboarding"
	"context"
)

// Server is the struct that implements the FrontendServiceServer interface.
type Server struct {
	pb.UnimplementedFrontendServiceServer
}

// GetCurrentOnboardingStage handles the request to get the current onboarding stage.
// It takes a context and an empty request as parameters, and returns a response containing the current onboarding stage.
func (s *Server) GetCurrentOnboardingStage(
	ctx context.Context,
	_ *generic.EmptyRequest) (*onboarding.GetCurrentOnboardingStageResponse, error) {
	// Return a response with the current onboarding stage set to EMAIL_VERIFICATION.
	return &onboarding.GetCurrentOnboardingStageResponse{
		CurrentStage: enums.OnboardingStage_ONBOARDING_STAGE_EMAIL_VERIFICATION,
	}, nil
}
