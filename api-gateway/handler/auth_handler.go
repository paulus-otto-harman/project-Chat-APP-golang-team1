package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"project/api-gateway/database"
	"project/api-gateway/service"
	pbAuth "project/auth-service/proto"
)

type AuthController struct {
	service service.AuthService
	logger  *zap.Logger
	cacher  database.Cacher
}

func NewAuthController(service service.AuthService, logger *zap.Logger, cacher database.Cacher) *AuthController {
	return &AuthController{service, logger, cacher}
}

func (ctrl *AuthController) Register(c *gin.Context) {
	var req pbAuth.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ctrl.service.Register(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
