package router

import (
	"github.com/Jessefranklin151/go_api_prac/go_api_prac/di"
	"github.com/Jessefranklin151/go_api_prac/go_api_prac/handler"
	"github.com/gin-gonic/gin"
)

func registerRoutes(router *gin.Engine) {
	v1 := router.Group("api/v1")

	playerHandler := di.Dependencies.PlayerHandler
	handler.RegisterPlayerHandlers(v1, playerHandler)

}
