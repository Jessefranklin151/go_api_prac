package repositories

import (
	"github.com/Jessefranklin151/go_api_prac/go_api_prac/data"
	"github.com/Jessefranklin151/go_api_prac/go_api_prac/models"
)

type PlayerRepository struct {
	dataLayer data.IPlayerData
}

func CreatePlayerRepository(dataLayer data.IPlayerData) *PlayerRepository {
	return &PlayerRepository{
		dataLayer: dataLayer,
	}
}

func (repo *PlayerRepository) GetPlayers() ([]models.Player, error) {
	return repo.dataLayer.FindPlayeres()

}

func (repo *PlayerRepository) CreatePlayer(player models.Player) (models.Player, error) {
	return repo.dataLayer.CreatePlayer(player)
}
