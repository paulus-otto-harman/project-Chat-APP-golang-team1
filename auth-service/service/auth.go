package service

import (
	"context"
	"log"
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
	x := a.Repo.Auth.Db.Error
	log.Println(x)
	return &pb.RegisterResponse{Otp: "1234"}, nil
}

func (a *AuthService) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {

	return &pb.ValidateTokenResponse{Username: "test@mail.com"}, nil
}
