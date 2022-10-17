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

/*func (r *Repository) FilmPrice(id int) (*ds.Kino, error) {

	product := &ds.Kino{}
	err := r.db.First(product, id).Error // find product with code D42
	if err != nil {
		return nil, err
	}

	return product, nil
}*/

func (r *Repository) GetFilmPrice(uuid string) (string, int, error) {
	var kino ds.Kino
	result := r.db.First(&kino, "uuid = ?", uuid)
	if result.Error != nil {
		return "no film found with this uuid", 0, result.Error
	}
	return kino.Name, kino.Price, nil
}

func (r *Repository) ChangePrice(uuid uuid.UUID, price uint64) (error, string) {
	var product ds.Kino
	product.UUID = uuid
	err := r.db.First(&product, "uuid = ?", uuid).Error
	if err != nil {
		return err, "record not found"
	}
	err = r.db.Model(&product).Update("price", price).Error
	if err != nil {
		return err, "record not update"
	}
	return nil, ""
}

func (r *Repository) DeleteFilm(uuid string) (string, error) {
	var product ds.Kino
	result := r.db.Delete(&product, "uuid = ?", uuid)
	if result.Error != nil {
		return "no product found with this uuid", result.Error
	}
	return uuid, nil
}

func (r *Repository) AddFilm(book ds.Kino) error {
	return r.db.Create(&book).Error

}
