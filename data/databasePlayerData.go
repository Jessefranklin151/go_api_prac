package data

import (
	"fmt"

	"github.com/Jessefranklin151/go_api_prac/go_api_prac/models"

	"github.com/Jessefranklin151/go_api_prac/go_api_prac/schemas"
	"gorm.io/gorm"
)

type DatabasePlayerData struct {
	DB *gorm.DB
}

func (d *DatabasePlayerData) FindPlayeres() ([]models.Player, error) {
	var playeres []schemas.PlayerGormModel

	if err := d.DB.Find(&playeres).Error; err != nil {
		return []models.Player{}, err
	}

	var result []models.Player

	for _, j := range playeres {
		result = append(result, models.Player{
			Id:     int32(j.ID),
			Nome:   j.Nome,
			Rating: j.Rating,
		})
	}

	return result, nil
}

func (d *DatabasePlayerData) CreatePlayer(player models.Player) (models.Player, error) {

	gormPlayer := schemas.PlayerGormModel{
		Nome:   player.Nome,
		Rating: player.Rating,
	}

	result := d.DB.Create(&gormPlayer)

	if result.Error != nil {
		return models.Player{}, result.Error
	}

	fmt.Printf("Player criado com sucesso: %+v\n", gormPlayer)

	created := models.Player{
		Id:     int32(gormPlayer.ID),
		Nome:   gormPlayer.Nome,
		Rating: gormPlayer.Rating,
	}

	return created, nil
}
