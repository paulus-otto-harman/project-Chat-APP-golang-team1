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

func (s *AuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	if err := s.repo.Auth.Create(&model.User{Username: req.Username, Password: req.Password}); err != nil {
		return nil, err
	}

	s.repo.Otp.Create()

	return &pb.RegisterResponse{Otp: "1234"}, nil
}

func (s *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	return &pb.LoginResponse{Token: "abcd"}, nil
}

func (s *AuthService) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {

	return &pb.ValidateTokenResponse{Username: "test@mail.com"}, nil
}
