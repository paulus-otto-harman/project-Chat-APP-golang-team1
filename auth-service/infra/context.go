package infra

import (
	"project/auth-service/config"
	"project/auth-service/database"
	"project/auth-service/log"
	"project/auth-service/repository"
	"project/auth-service/service"

	"go.uber.org/zap"
)

type ServiceContext struct {
	Cfg config.Config
	Log *zap.Logger
	Svc *service.Service
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
	services := service.NewService(repo, appConfig, logger)

	return &ServiceContext{Cfg: appConfig, Svc: &services, Log: logger}, nil
}
