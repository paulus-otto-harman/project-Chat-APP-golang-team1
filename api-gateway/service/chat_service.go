package service

import "go.uber.org/zap"

type ChatService interface {
}

type chatService struct {
	serviceUrl string
	log        *zap.Logger
}

func NewChatService(serviceUrl string, log *zap.Logger) ChatService {
	return &chatService{serviceUrl, log}
}
