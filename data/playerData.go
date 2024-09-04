package data

import "github.com/Jessefranklin151/go_api_prac/go_api_prac/models"

type IPlayerData interface {
	FindPlayeres() ([]models.Player, error)
	CreatePlayer(player models.Player) (models.Player, error)
}
