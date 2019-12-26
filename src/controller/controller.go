package controller

import (
	"github.com/gin-gonic/gin"
	"go-learn/components/logger"
	"go-learn/config"
	"go-learn/openapi"
)

func Start() {

	webconfig := config.WebMetaConfig

	// 初始化引擎
	engine := gin.New()

	engine.Use(logger.GinMiddleLogger(), gin.Recovery())

	/*openapi分组*/
	openapiRouter := engine.Group(webconfig.Web.Context + "/openapi")

	// 注册一个路由和处理函数
	openapi.Prediction(openapiRouter)

	openapi.Features(openapiRouter)

	// 绑定端口，然后启动应用
	engine.Run(":" + webconfig.Web.Server.Port)
}
