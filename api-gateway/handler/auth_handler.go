package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
	"project/api-gateway/database"
	"project/api-gateway/service"
	pbAuth "project/auth-service/proto"
)

type AuthController struct {
	service service.Service
	logger  *zap.Logger
	cacher  database.Cacher
}

func NewAuthController(service service.Service, logger *zap.Logger, cacher database.Cacher) *AuthController {
	return &AuthController{service, logger, cacher}
}

func (ctrl *AuthController) Register(c *gin.Context) {
	var req pbAuth.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := ctrl.service.Auth.Register(context.Background(), &req)
	if err != nil {
		BadResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	emailData := EmailData{ID: uuid.New(), OTP: "123456"}
	_, err = ctrl.service.Email.Send("tes@mailinator.com", "Chateo OTP", "otp", emailData)
	if err != nil {
		ctrl.logger.Error("failed to send email", zap.Error(err))
		BadResponse(c, "failed to send email", http.StatusInternalServerError)
		return
	}

	GoodResponseWithData(c, "registration success. otp sent", http.StatusOK, res)
}

type EmailData struct {
	ID  uuid.UUID
	OTP string
}
