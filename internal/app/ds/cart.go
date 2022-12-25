package ds

import "github.com/google/uuid"

type Cart struct {
	FilmUUID uuid.UUID `example:"976c088c-f218-422b-aff6-f9e1cf792860"`
	UserUUID uuid.UUID `example:"bb5b65bb-653a-42bf-9cf4-76b62e85a7bc"`
	Quantity uint64    `example:"3"`
}

func (Cart) TableName() string {
	return "cart"
}
