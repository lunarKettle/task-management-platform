package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lunarKettle/task-management-platform/api-gateway/internal/models/dto"
	pb "github.com/lunarKettle/task-management-platform/api-gateway/proto"
)

func (s *HTTPServer) registerUser(w http.ResponseWriter, r *http.Request) error {
	var regUserReq dto.RegisterUserRequestDTO
	err := json.NewDecoder(r.Body).Decode(&regUserReq)
	if err != nil {
		return fmt.Errorf("error while decoding request body: %w", err)
	}
	defer r.Body.Close()

	grpcRequest := &pb.RegisterRequest{
		Email:    regUserReq.Email,
		Password: regUserReq.Password,
		Username: regUserReq.Username,
		Role:     regUserReq.Role,
	}

	grpcResponse, err := s.grpcClient.Register(grpcRequest)
	if err != nil {
		return fmt.Errorf("failed to register user: %w", err)

	}

	reqUserResp := dto.RegisterUserResponseDTO{
		AccessToken: grpcResponse.Token,
	}
	if err := json.NewEncoder(w).Encode(reqUserResp); err != nil {
		return fmt.Errorf("failed to encode response to JSON: %w", err)
	}
	return err
}

func (s *HTTPServer) authenticate(w http.ResponseWriter, r *http.Request) error {
	authUserReq, err := extractBasicAuth(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return err
	}

	grpcRequest := &pb.AuthRequest{
		Username: authUserReq.Username,
		Password: authUserReq.Password,
	}

	grpcResponse, err := s.grpcClient.Authenticate(grpcRequest)
	if err != nil {
		return fmt.Errorf("failed to authenticate user: %w", err)

	}

	reqUserResp := dto.LoginUserResponseDTO{
		AccessToken: grpcResponse.Token,
	}
	if err := json.NewEncoder(w).Encode(reqUserResp); err != nil {
		return fmt.Errorf("failed to encode response to JSON: %w", err)
	}
	return err
}
