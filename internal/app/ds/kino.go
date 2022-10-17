package ds

import "github.com/google/uuid"

type Kino struct {
	UUID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key;"`
	Name       string
	Release    uint64
	Grade      float64
	Genre      string
	Price      int
	WhatchTime uint64
	Summary    string
}
