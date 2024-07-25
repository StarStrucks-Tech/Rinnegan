package pkg

import (
	"Rinnegan/proto-generated/protos-frontend/enums"
	pb "Rinnegan/proto-generated/protos-frontend/frontend"
	"Rinnegan/proto-generated/protos-frontend/generic"
	"Rinnegan/proto-generated/protos-frontend/onboarding"
	"context"
	"fmt"
	"math/rand"
	"time"
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

// OTP storage (in-memory for demonstration purposes)
var otpStore = map[string]string{}

func (s *Server) VerifyAadhar(ctx context.Context, req *onboarding.AadharVerificationRequest) (*onboarding.AadharVerificationResponse, error) {
	if req.Otp == "" {
		// Generate OTP
		otp := fmt.Sprintf("%06d", rand.Intn(1000000))
		otpTransactionID := fmt.Sprintf("%d", time.Now().UnixNano())
		otpStore[otpTransactionID] = otp

		// Simulate sending OTP (in reality, integrate with an SMS gateway)
		fmt.Printf("Generated OTP: %s for transaction ID: %s\n", otp, otpTransactionID)

		return &onboarding.AadharVerificationResponse{
			Success:          true,
			Message:          "OTP generated successfully",
			OtpTransactionId: otpTransactionID,
			GeneratedOtp:     otp, // Include the generated OTP in the response
		}, nil
	} else {
		// Verify OTP
		storedOtp, exists := otpStore[req.Otp]
		if !exists || storedOtp != req.Otp {
			return &onboarding.AadharVerificationResponse{
				Success: false,
				Message: "OTP verification failed",
			}, nil
		}

		// OTP verified successfully
		delete(otpStore, req.Otp) // Clean up the stored OTP

		return &onboarding.AadharVerificationResponse{
			Success: true,
			Message: "OTP verification successful",
		}, nil
	}
}
