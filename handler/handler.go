package handler

import (
	"github.com/gin-gonic/gin"
)

func RegisterPlayerHandlers(router *gin.RouterGroup, handler *PlayerHandler) {
	router.GET("/player", handler.getPlayer)
	router.POST("/player", handler.createPlayer)
}
