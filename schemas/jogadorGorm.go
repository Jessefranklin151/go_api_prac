package schemas

import (
	"gorm.io/gorm"
)

type PlayerGormModel struct {
	gorm.Model
	Nome   string `json:"nome"`
	Rating int8   `json:"rating"`
}

type Tabler interface {
	TableName() string
}

func (PlayerGormModel) TableName() string {
	return "playeres"
}
