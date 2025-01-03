package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"project/auth-service/config"
	"project/auth-service/model"
	pb "project/auth-service/proto"
	"project/auth-service/repository"
	"strings"
)

type AuthService struct {
	repo    repository.Repository
	log     *zap.Logger
	rsaKeys config.RSAKeys
	pb.UnimplementedAuthServiceServer
}

func NewAuthService(repo repository.Repository, log *zap.Logger, rsaKeys config.RSAKeys) *AuthService {
	return &AuthService{repo: repo, log: log, rsaKeys: rsaKeys}
}

func (s *AuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user := model.User{Email: req.Email}
	if err := s.repo.Auth.Create(&user); err != nil {
		return nil, err
	}

	otp := model.Otp{UserID: user.ID}
	s.repo.Otp.Create(&otp)

	return &pb.RegisterResponse{Otp: otp.Otp}, nil
}

func (s *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := s.repo.Auth.Get(model.User{Email: req.Email})
	if err != nil {
		return nil, err
	}

	otp := model.Otp{UserID: user.ID}
	s.repo.Otp.Create(&otp)

	return &pb.LoginResponse{Otp: otp.Otp}, nil
}

func (s *AuthService) ValidateOtp(ctx context.Context, req *pb.ValidateOtpRequest) (*pb.ValidateOtpResponse, error) {
	return &pb.ValidateOtpResponse{Token: "jwt-token-here"}, nil
}

func (s *AuthService) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(s.rsaKeys.PublicKey))
	if err != nil {
		return nil, errors.New("failed to get authentication key")
	}

	if len(req.Token) == 0 {
		return nil, errors.New("authorization token not found")
	}

	claims := &customClaims{}
	tkn, err := jwt.ParseWithClaims(strings.Split(req.Token, "Bearer ")[1], claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New(fmt.Sprintf("unexpected method: %s", token.Header["alg"]))
		}
		return key, nil
	})

	if err != nil {
		return nil, errors.New("fail to validate signature or session expired")
	}

	if !tkn.Valid {
		return nil, errors.New("invalid token")
	}

	return &pb.ValidateTokenResponse{Email: claims.Email}, nil
}

type customClaims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	IP    string `json:"ip"`
	jwt.StandardClaims
}
