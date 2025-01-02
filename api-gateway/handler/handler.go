package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"project/api-gateway/database"
	"project/api-gateway/model"
	"project/api-gateway/service"
)

type Handler struct {
	AuthHandler    AuthController
	ContactHandler ContactController
	OtpHandler     OtpController
	UserHandler    UserController
}

func NewHandler(service service.Service, logger *zap.Logger, rdb database.Cacher) *Handler {
	return &Handler{
		AuthHandler:    *NewAuthController(service, logger, rdb),
		ContactHandler: *NewContactController(service, logger),
		OtpHandler:     *NewOtpController(service, logger),
		UserHandler:    *NewUserController(service, logger),
	}
}

func BadResponse(c *gin.Context, message string, statusCode int) {
	c.JSON(statusCode, model.Response{
		Status:  false,
		Message: message,
	})
}

func GoodResponseWithData(c *gin.Context, message string, statusCode int, data interface{}) {
	c.JSON(statusCode, model.Response{
		Status:  true,
		Message: message,
		Data:    data,
	})
}

func GoodResponseWithPage(c *gin.Context, message string, statusCode, total, totalPages, page, Limit int, data interface{}) {
	c.JSON(statusCode, model.DataPage{
		Status:      true,
		Message:     message,
		Total:       int64(total),
		Pages:       totalPages,
		CurrentPage: uint(page),
		Limit:       uint(Limit),
		Data:        data,
	})
}
