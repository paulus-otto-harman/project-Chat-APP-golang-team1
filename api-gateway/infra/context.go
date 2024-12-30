package infra

import (
	"go.uber.org/zap"
	"project/api-gateway/config"
	"project/api-gateway/database"
	"project/api-gateway/log"
	"project/api-gateway/repository"
	"project/api-gateway/service"
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

	rdb := database.NewCacher(appConfig, 60*60)

	// instance repository
	repo := repository.NewRepository(rdb, appConfig, logger)

	// instance service
	services := service.NewService(repo, appConfig, logger)

	//return &ServiceContext{Cacher: rdb, Cfg: appConfig, Svc: &services, Log: logger}, nil
	return &ServiceContext{Cfg: appConfig, Svc: &services, Log: logger}, nil
}
