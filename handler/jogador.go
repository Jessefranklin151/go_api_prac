package handler

import (
	"net/http"

	"github.com/Jessefranklin151/go_api_prac/go_api_prac/controllers"
	"github.com/Jessefranklin151/go_api_prac/go_api_prac/models"
	"github.com/gin-gonic/gin"
)

type PlayerHandler struct {
	controller *controllers.PlayerController
}

func InitPlayerHandler(playerController *controllers.PlayerController) *PlayerHandler {
	return &PlayerHandler{
		controller: playerController,
	}
}

func (handler *PlayerHandler) getPlayer(ctx *gin.Context) {
	playeres, err := handler.controller.GetPlayers()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, playeres)
}

func (handler *PlayerHandler) createPlayer(ctx *gin.Context) {
	var j models.Player

	err := ctx.ShouldBindJSON(&j)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	player, err := handler.controller.CreatePlayer(j)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, player)
}
