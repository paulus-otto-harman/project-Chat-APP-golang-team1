package infra

import (
	"project/chat-service/config"
	"project/chat-service/database"
	"project/chat-service/handler"
	"project/chat-service/log"
	"project/chat-service/repository"
	"project/chat-service/service"

	"go.uber.org/zap"
)

type ServiceContext struct {
	Cfg config.Config
	Log *zap.Logger
	Ctl handler.Handler
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

	// instance database
	db, err := database.ConnectDB(appConfig)
	if err != nil {
		return handlerError(err)
	}

	// instance repository
	repo := repository.NewRepository(db, logger)

	// instance service
	services := service.NewService(repo, logger)

	// instance controller
	Ctl := handler.NewHandler(services, logger)

	return &ServiceContext{Cfg: appConfig, Ctl: *Ctl, Log: logger}, nil
}
