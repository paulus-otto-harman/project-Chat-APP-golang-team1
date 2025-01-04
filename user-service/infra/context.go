package infra

import (
	"user_service/config"
	"user_service/database"
	"user_service/log"
	"user_service/repository"
	"user_service/service"

	"go.uber.org/zap"
)

type ServiceContext struct {
	Cacher database.Cacher
	Cfg    config.Config
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

	// instance database
	db, err := database.ConnectDB(appConfig)
	if err != nil {
		return handlerError(err)
	}

	rdb := database.NewCacher(appConfig, 60*60)

	// instance repository
	repo := repository.NewRepository(db)

	// instance service
	services := service.NewService(repo, appConfig, logger)

	return &ServiceContext{Cacher: rdb, Cfg: appConfig, Svc: &services, Log: logger}, nil
}
