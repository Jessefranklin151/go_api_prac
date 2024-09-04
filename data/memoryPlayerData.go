package data

import "github.com/Jessefranklin151/go_api_prac/go_api_prac/models"

type MemoryPlayerData struct {
}

var memoruPlayeres []models.Player = []models.Player{
	{Id: 1, Nome: "Player1", Rating: 89},
	{Id: 2, Nome: "Player2", Rating: 85},
}

func (d *MemoryPlayerData) FindPlayeres() ([]models.Player, error) {
	return []models.Player{
		{Id: 1, Nome: "Player1", Rating: 89},
		{Id: 2, Nome: "Player2", Rating: 85},
	}, nil
}

func (d *MemoryPlayerData) CreatePlayer(player models.Player) (models.Player, error) {
	memoruPlayeres = append(memoruPlayeres, player)
	return player, nil
}
