package repository

import (
	"awesomeProject/internal/app/ds"
	"awesomeProject/internal/app/dsn"
	"context"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(ctx context.Context) (*Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn.FromEnv()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) GetFilmList() ([]ds.Kino, error) {

	var kinos []ds.Kino
	result := r.db.Find(&kinos)
	if result.Error != nil {
		return kinos, result.Error
	}
	return kinos, nil

}

func (r *Repository) AddFilm(kino ds.Kino) error {
	err := r.db.Create(&kino).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetFilmPrice(uuid string) (uint64, error) {
	var kino ds.Kino
	result := r.db.First(&kino, "uuid = ?", uuid)
	if result.Error != nil {
		return 0, result.Error
	}
	return kino.Price, nil
}

func (r *Repository) ChangePrice(uuid uuid.UUID, price uint64) error {
	var kino ds.Kino
	kino.UUID = uuid
	result := r.db.Model(&kino).Update("Price", price)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Repository) DeleteFilm(uuid string) error {
	var kino ds.Kino
	result := r.db.Delete(&kino, "uuid = ?", uuid)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
