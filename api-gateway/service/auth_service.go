package service

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pbAuth "project/auth-service/proto"
)

type AuthService interface {
	Register(ctx context.Context, req *pbAuth.RegisterRequest) (*pbAuth.RegisterResponse, error)
}
type authService struct {
	log *zap.Logger
}

func NewAuthService(log *zap.Logger) AuthService {
	return &authService{log: log}
}

func (s *authService) Register(ctx context.Context, req *pbAuth.RegisterRequest) (*pbAuth.RegisterResponse, error) {
	authClient := makeAuthClient()
	return authClient.Register(ctx, req)
}

func makeAuthClient() pbAuth.AuthServiceClient {
	authConn, err := grpc.NewClient("localhost:51151", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("can't init auth grpc client %w", err)
	}
	//defer authConn.Close()
	return pbAuth.NewAuthServiceClient(authConn)
}
