package infra

import (
	"api_gateway/config"
	"api_gateway/database"
	"api_gateway/handler"
	"api_gateway/helper"
	"api_gateway/log"
	pbAuth "api_gateway/proto/auth_proto"
	pbUser "api_gateway/proto/user_proto"
	"api_gateway/service"
	"fmt"

	"go.uber.org/zap"
)

type ServiceContext struct {
	Cacher database.Cacher
	Cfg    config.Config
	Ctl    handler.Handler
	Log    *zap.Logger
	Svc    *service.Service
}

func NewServiceContext() (*ServiceContext, error) {

	handlerError := func(err error) (*ServiceContext, error) {
		return nil, err
	}

	// instance config
	appConfig, err := config.LoadConfig()
	if err != nil {
		return handlerError(err)
	}

	// instance logger
	logger, err := log.InitZapLogger(appConfig)
	if err != nil {
		return handlerError(err)
	}

	rdb := database.NewCacher(appConfig, 60*60)

	// instance service
	services := service.NewService(appConfig, logger)

	// instance gRPC Connection
	authConn := helper.NewConnection(fmt.Sprintf("%s:%s", appConfig.AuthServiceIp, appConfig.AuthServicePort))
	pbAuth := pbAuth.NewAuthServiceClient(authConn)
	userConn := helper.NewConnection(fmt.Sprintf("%s:%s", appConfig.UserServiceIp, appConfig.UserServicePort))
	pbUser := pbUser.NewUserServiceClient(userConn)
	// instance controller
	Ctl := handler.NewHandler(logger, rdb, pbAuth, pbUser)
	//return &ServiceContext{Cacher: rdb, Cfg: appConfig, Svc: &services, Log: logger}, nil
	return &ServiceContext{Cfg: appConfig, Ctl: *Ctl, Svc: &services, Log: logger}, nil
}
