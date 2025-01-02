package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
	"project/api-gateway/database"
	"project/api-gateway/model"
	"project/api-gateway/service"
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
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := ctrl.service.Auth.Register(user)
	if err != nil {
		BadResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	emailData := EmailData{ID: uuid.New(), OTP: res.Otp}
	_, err = ctrl.service.Email.Send(user.Email, "Chateo OTP", "otp", emailData)
	if err != nil {
		ctrl.logger.Error("failed to send email", zap.Error(err))
		BadResponse(c, "failed to send email", http.StatusInternalServerError)
		return
	}

	GoodResponseWithData(c, "registration success. otp sent", http.StatusOK, nil)
}

type EmailData struct {
	ID  uuid.UUID
	OTP string
}
