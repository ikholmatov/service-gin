package api

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	v1 "github.com/venomuz/project4/API-GATEWAY/api/handlers/v1"
	"github.com/venomuz/project4/API-GATEWAY/config"
	"github.com/venomuz/project4/API-GATEWAY/pkg/logger"
	"github.com/venomuz/project4/API-GATEWAY/services"
)

type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
}

func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
	})

	api := router.Group("/v1")
	api.POST("/users", handlerV1.CreateUser)
	//api.GET("/users/:id", handlerV1.GetUser)
	// api.GET("/users", handlerV1.ListUsers)
	// api.PUT("/users/:id", handlerV1.UpdateUser)
	// api.DELETE("/users/:id", handlerV1.DeleteUser)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")

	return router
}
