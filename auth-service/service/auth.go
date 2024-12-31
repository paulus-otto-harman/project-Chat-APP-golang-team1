package service

import (
	"context"
	"project/auth-service/model"
	pb "project/auth-service/proto"
	"project/auth-service/repository"
)

type AuthService struct {
	Repo repository.Repository
	pb.UnimplementedAuthServiceServer
}

func (a *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	return &pb.LoginResponse{Token: "abcd"}, nil
}

func (a *AuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	if err := a.Repo.Auth.Create(&model.User{Username: req.Username, Password: req.Password}); err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{Otp: "1234"}, nil
}

func (a *AuthService) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {

	return &pb.ValidateTokenResponse{Username: "test@mail.com"}, nil
}
