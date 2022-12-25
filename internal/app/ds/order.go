package ds

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	UUID     uuid.UUID `db:"uuid" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Film     string    `db:"film"`
	Quantity uint64    `db:"quantity"`
	UserUUID uuid.UUID `db:"user_uuid"`
	Date     time.Time `db:"date" gorm:"type:timestamp"`
	Status   string    `db:"status"`
}
