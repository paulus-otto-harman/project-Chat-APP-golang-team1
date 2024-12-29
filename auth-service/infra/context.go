package infra

import (
	"project/auth-service/config"
	"project/auth-service/database"
	"project/auth-service/handler"
	"project/auth-service/log"
	"project/auth-service/repository"
	"project/auth-service/service"

	"go.uber.org/zap"
)

type ServiceContext struct {
	Cacher database.Cacher
	Cfg    config.Config
	Log    *zap.Logger
	Ctl    handler.Handler
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
	repo := repository.NewRepository(db, rdb, appConfig, logger)

	// instance service
	services := service.NewService(repo, appConfig, logger)

	// instance controller
	Ctl := handler.NewHandler(services, logger, rdb)

	return &ServiceContext{Cacher: rdb, Cfg: appConfig, Ctl: *Ctl, Log: logger}, nil
}
