package service

import (
	"context"
	"go.uber.org/zap"
	"project/api-gateway/helper"
	"project/api-gateway/model"
	pbAuth "project/auth-service/proto"
)

type AuthService interface {
	GetOtp(user model.User) (*pbAuth.CreateOtpResponse, error)
	ValidateOtp(otp model.Otp) (*string, error)
}
type authService struct {
	serviceUrl string
	log        *zap.Logger
}

func NewAuthService(serviceUrl string, log *zap.Logger) AuthService {
	return &authService{serviceUrl, log}
}

func (s *authService) GetOtp(user model.User) (*pbAuth.CreateOtpResponse, error) {
	authConn := helper.MustConnect(s.serviceUrl)
	defer authConn.Close()

	authClient := pbAuth.NewAuthServiceClient(authConn)

	req := &pbAuth.CreateOtpRequest{Email: user.Email}
	res, err := authClient.CreateOtp(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *authService) ValidateOtp(otp model.Otp) (*string, error) {
	authConn := helper.MustConnect(s.serviceUrl)
	defer authConn.Close()

	authClient := pbAuth.NewAuthServiceClient(authConn)
	req := &pbAuth.ValidateOtpRequest{Id: otp.ID.String(), Otp: otp.Otp}
	res, err := authClient.ValidateOtp(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return &res.Token, nil
}
