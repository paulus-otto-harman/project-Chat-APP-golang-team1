package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"project/api-gateway/service"
)

type ChatController struct {
	service service.Service
	logger  *zap.Logger
}

func NewChatController(service service.Service, logger *zap.Logger) *ChatController {
	return &ChatController{service, logger}
}

func (ctrl *ChatController) Websocket(c *gin.Context) {

}
