package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"project/api-gateway/service"
)

type OtpController struct {
	service service.Service
	logger  *zap.Logger
}

func NewOtpController(service service.Service, logger *zap.Logger) *OtpController {
	return &OtpController{service, logger}
}

func (ctrl *OtpController) Validate(c *gin.Context) {}
