package service

import (
	"context"
	"project/auth-service/model"
	pb "project/auth-service/proto"
	"project/auth-service/repository"
)

type AuthService struct {
	repo repository.Repository
	pb.UnimplementedAuthServiceServer
}

func NewAuthService(repo repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	if err := s.repo.Auth.Create(&model.User{Email: req.Email}); err != nil {
		return nil, err
	}

	s.repo.Otp.Create()

	return &pb.RegisterResponse{Otp: "4321"}, nil
}

func (s *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	return &pb.LoginResponse{Token: "abcd"}, nil
}

func (s *AuthService) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {

	return &pb.ValidateTokenResponse{Email: "test@mail.com"}, nil
}
