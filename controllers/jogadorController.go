package controllers

import (
	"github.com/Jessefranklin151/go_api_prac/go_api_prac/models"
	"github.com/Jessefranklin151/go_api_prac/go_api_prac/repositories"
)

type PlayerController struct {
	repo *repositories.PlayerRepository
}

func CratePlayerController(playerRepository *repositories.PlayerRepository) *PlayerController {
	return &PlayerController{
		repo: playerRepository,
	}
}

func (c *PlayerController) GetPlayers() ([]models.Player, error) {
	return c.repo.GetPlayers()
}

func (c *PlayerController) CreatePlayer(player models.Player) (models.Player, error) {
	return c.repo.CreatePlayer(player)
}
