package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"project/api-gateway/service"
)

type UserController struct {
	service service.Service
	logger  *zap.Logger
}

func NewUserController(service service.Service, logger *zap.Logger) *UserController {
	return &UserController{service, logger}
}

func (ctrl *UserController) Update(c *gin.Context) {
	username := c.MustGet("username")

	GoodResponseWithData(c, "profile updated", http.StatusOK, gin.H{"user_id": username})
}
