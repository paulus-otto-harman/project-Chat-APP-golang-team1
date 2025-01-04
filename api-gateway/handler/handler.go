package handler

import (
	"api_gateway/database"
	"api_gateway/model"
	pbAuth "api_gateway/proto/auth_proto"
	pbUser "api_gateway/proto/user_proto"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	//"project/infra/jwt"
)

type Handler struct {
	logger *zap.Logger
	rdb    database.Cacher
	pbAuth pbAuth.AuthServiceClient
	pbUser pbUser.UserServiceClient
}

func NewHandler(logger *zap.Logger, rdb database.Cacher, pbAuth pbAuth.AuthServiceClient, pbUser pbUser.UserServiceClient) *Handler {
	return &Handler{logger: logger, rdb: rdb, pbAuth: pbAuth, pbUser: pbUser}
}
func (h *Handler) Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	_, err := h.pbUser.CreateUser(context.Background(), &pbUser.CreateUserRequest{Email: user.Email})
	if err != nil {
		BadResponse(c, "Email Already Registered", http.StatusBadRequest)
		return
	}
	authGrpc, err := h.pbAuth.Register(context.Background(), &pbAuth.RegisterRequest{Email: user.Email})
	if err != nil {
		BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	GoodResponseWithData(c, "Register success. otp sent", http.StatusCreated, model.OTP{OTP: authGrpc.Otp})
}
func (h *Handler) Login(c *gin.Context) {
	var email pbAuth.LoginRequest
	if err := c.ShouldBindJSON(&email); err != nil {
		BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	resGrpc, err := h.pbAuth.Login(context.Background(), &email)
	if err != nil {
		BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	GoodResponseWithData(c, "login success. otp sent", http.StatusOK, model.OTP{OTP: resGrpc.Otp})
}
func (h *Handler) ValidateOTP(c *gin.Context) {
	var otp pbAuth.ValidateOtpRequest
	if err := c.ShouldBindJSON(&otp); err != nil {
		BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	resGrpc, err := h.pbAuth.ValidateOtp(context.Background(), &otp)
	if err != nil {
		BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	GoodResponseWithData(c, "login success. otp sent", http.StatusOK, model.Token{Token: resGrpc.Token})
}
func (h *Handler) ValidateToken(c *gin.Context) {
	token := c.GetHeader("token")
	if token == "" {
		BadResponse(c, "Unauthentication", http.StatusBadRequest)
		c.Abort()
		return
	}
	resGrpc, err := h.pbAuth.ValidateToken(context.Background(), &pbAuth.ValidateTokenRequest{Token: token})
	if err != nil {
		BadResponse(c, "Invalid Token", http.StatusBadRequest)
		c.Abort()
		return
	}
	c.Set("email", resGrpc.Email)
	c.Next()
}
func (h *Handler) GetAllUsers(c *gin.Context) {
	query := c.Query("filter")
	md := metadata.Pairs(
		"filter", query,
	)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	resGrpc, err := h.pbUser.GetAllUsers(ctx, &pbUser.Empty{})
	if err != nil {
		log.Println(err)
		BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	GoodResponseWithData(c, "Get All Users Success", http.StatusOK, resGrpc.Users)
}
func (h *Handler) UpdateProfile(c *gin.Context) {
	var profileReq pbUser.UpdateUserRequest
	if err := c.ShouldBindJSON(&profileReq); err != nil {
		BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	email, _ := c.Get("email")
	md := metadata.Pairs(
		"email", email.(string),
	)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	resGrpc, err := h.pbUser.UpdateUser(ctx, &profileReq)
	if err != nil {
		log.Println(err)
		BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	GoodResponseWithData(c, resGrpc.Message, http.StatusOK, nil)
}

func BadResponse(c *gin.Context, message string, statusCode int) {
	c.JSON(statusCode, model.Response{
		StatusCode: statusCode,
		Message:    message,
	})
}

func GoodResponseWithData(c *gin.Context, message string, statusCode int, data interface{}) {
	c.JSON(statusCode, model.Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
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
