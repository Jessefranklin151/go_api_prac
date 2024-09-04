package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() *gin.Engine {

	router := gin.Default()

	createMiddlewares(router)

	setupProxies(router)

	registerRoutes(router)

	return router
}
