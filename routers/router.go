package Routers

import (
	"github.com/macduyhai/loadBalancingGrpcApi/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/macduyhai/loadBalancingGrpcApi/config"
	"github.com/macduyhai/loadBalancingGrpcApi/middlewares"
	"github.com/macduyhai/loadBalancingGrpcApi/services"
)

type Router struct {
	config *config.Config
	db     *gorm.DB
}

// Tao ham de chua thong tin
func NewRouter(conf *config.Config, db *gorm.DB) Router {
	return Router{config: conf, db: db}
}

// khoi tao gin Engine
func (router *Router) InitGin() (*gin.Engine, error) {
	provider := services.NewProviderService(router.config, router.db)
	controller := controllers.NewController(provider)

	engine := gin.Default()
	engine.Use(middlewares.CORSMiddleware())
	engine.GET("/ping", controller.Ping)
	engine.POST("/login", controller.Login)

	userAuthenMiddleWare := middlewares.CheckAPIkey{Apikey: router.config.APIKEY}
	{
		user := engine.Group("/api/v1/user")
		user.Use(userAuthenMiddleWare.Check)
		user.POST("/add", controller.Create)
		user.PUT("/edit", controller.Edit)
		user.DELETE("/delete", controller.Delete)
	}

	return engine, nil

}
