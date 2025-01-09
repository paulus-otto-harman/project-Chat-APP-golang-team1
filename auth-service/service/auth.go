package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"math/rand"
	"project/auth-service/config"
	"project/auth-service/model"
	pb "project/auth-service/proto"
	"project/auth-service/repository"
	"strings"
	"time"
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

func (s *AuthService) CreateOtp(ctx context.Context, req *pb.CreateOtpRequest) (*pb.CreateOtpResponse, error) {
	seed := time.Now().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	oneTimePassword := fmt.Sprintf("%04d", rng.Intn(10000)) // Generate 4 digit OTP

	otp := model.Otp{UserEmail: req.Email, Otp: oneTimePassword}
	if err := s.repo.Otp.Create(&otp); err != nil {
		return nil, err
	}
	return &pb.CreateOtpResponse{Otp: otp.Otp, Id: otp.ID.String()}, nil
}

func (s *AuthService) ValidateOtp(ctx context.Context, req *pb.ValidateOtpRequest) (*pb.ValidateOtpResponse, error) {
	OtpID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	otp, err := s.repo.Otp.Update(model.Otp{ID: OtpID, Otp: req.Otp})
	if err != nil {
		return nil, err
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(s.rsaKeys.PrivateKey))
	if err != nil {
		s.log.Error("Failed to parse RSA key", zap.Error(err))
		return nil, err
	}

	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &customClaims{
		Email:          otp.UserEmail,
		IP:             "",
		StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTime.Unix()},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return nil, err
	}

	return &pb.ValidateOtpResponse{Token: token}, nil
}

func (s *AuthService) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(s.rsaKeys.PublicKey))
	if err != nil {
		s.log.Error("failed to parse public key", zap.Error(err))
		return nil, err
	}

	if len(req.Token) == 0 {
		s.log.Error("authorization token is empty")
		return nil, errors.New("authorization token not found")
	}

	claims := &customClaims{}
	jsonWebToken, err := jwt.ParseWithClaims(strings.ReplaceAll(req.Token, "Bearer ", ""), claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			s.log.Error("unexpected signing method")
			return nil, errors.New(fmt.Sprintf("unexpected method: %s", token.Header["alg"]))
		}
		return key, nil
	})

	if err != nil {
		s.log.Error("failed to parse token", zap.Error(err))
		return nil, err
	}

	if !jsonWebToken.Valid {
		return nil, errors.New("invalid token")
	}

	return &pb.ValidateTokenResponse{Email: claims.Email}, nil
}

type customClaims struct {
	Email string `json:"email"`
	IP    string `json:"ip"`
	jwt.StandardClaims
}
