package main

import (
	"github.com/venomuz/project4/API-GATEWAY/api"
	"github.com/venomuz/project4/API-GATEWAY/config"
	"github.com/venomuz/project4/API-GATEWAY/pkg/logger"
	"github.com/venomuz/project4/API-GATEWAY/services"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api_gateway")

	serviceManager, err := services.NewServiceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
		ServiceManager: serviceManager,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}
}
