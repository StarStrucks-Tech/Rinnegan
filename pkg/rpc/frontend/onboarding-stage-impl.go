package pkg

import (
	"Rinnegan/proto-generated/protos-frontend/enums"
	pb "Rinnegan/proto-generated/protos-frontend/frontend"
	"Rinnegan/proto-generated/protos-frontend/generic"
	"Rinnegan/proto-generated/protos-frontend/onboarding"
	"context"
)

type Server struct {
	pb.UnimplementedFrontendServiceServer
}

func (s *Server) GetCurrentOnboardingStage(
	ctx context.Context,
	_ *generic.EmptyRequest) (*onboarding.GetCurrentOnboardingStageResponse, error) {
	return &onboarding.GetCurrentOnboardingStageResponse{
		CurrentStage: enums.OnboardingStage_ONBOARDING_STAGE_EMAIL_VERIFICATION,
	}, nil
}
