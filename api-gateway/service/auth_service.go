package service

import (
	"context"
	"go.uber.org/zap"
	"project/api-gateway/helper"
	"project/api-gateway/model"
	pbAuth "project/auth-service/proto"
)

type AuthService interface {
	Register(user model.User) (*pbAuth.RegisterResponse, error)
	Login(ctx context.Context, req *pbAuth.LoginRequest) (*pbAuth.LoginResponse, error)
	CreateOtp()
	ValidateOtp() (*string, error)
}
type authService struct {
	serviceUrl string
	log        *zap.Logger
}

func NewAuthService(serviceUrl string, log *zap.Logger) AuthService {
	return &authService{serviceUrl, log}
}

func (s *authService) Register(user model.User) (*pbAuth.RegisterResponse, error) {
	authConn := helper.NewConnection(s.serviceUrl)
	defer authConn.Close()

	authClient := pbAuth.NewAuthServiceClient(authConn)

	req := &pbAuth.RegisterRequest{Username: user}
	res, err := authClient.Register(context.Background(), req)
	if err != nil {
		return nil, nil
	}

	return res, nil
}

func (s *authService) Login(ctx context.Context, req *pbAuth.LoginRequest) (*pbAuth.LoginResponse, error) {
	//TODO implement me
	return nil, nil
}

func (s *authService) CreateOtp() {
	//TODO implement me
	return
}

func (s *authService) ValidateOtp() (*string, error) {
	//TODO implement me
	return nil, nil
}
