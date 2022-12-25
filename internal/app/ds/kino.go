package ds

import (
	"github.com/google/uuid"
)

type Film struct {
	UUID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name       string
	Release    uint64
	Grade      float64
	Genre      string
	Price      uint64
	WhatchTime uint64
	Summary    string
	Image      string
	Video      string
}

func (Film) TableName() string {
	return "film"
}

type FilmPrice struct {
	Price uint64
}

type QuantityFilms struct {
	Quantity uint64 `example:"10"`
}
