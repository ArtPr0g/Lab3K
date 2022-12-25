package repository

import (
	"awesomeProject/internal/app/ds"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *Repository) GetCart(userUUID uuid.UUID) ([]ds.Cart, error) {
	var cart []ds.Cart
	err := r.db.Order("film_uuid").Find(&cart, "user_uuid = ?", userUUID).Error
	return cart, err
}

func (r *Repository) GetCart1(filmUUID uuid.UUID, userUUID uuid.UUID) (ds.Cart, error) {
	var cart ds.Cart
	err := r.db.First(&cart, "film_uuid = ? and user_uuid = ?", filmUUID, userUUID).Error
	return cart, err
}

func (r *Repository) DeleteCart(filmUUID uuid.UUID, userUUID uuid.UUID) (int, error) {
	var cart ds.Cart
	err := r.db.First(&cart, "film_uuid = ? and user_uuid = ?", filmUUID, userUUID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 404, err
		}
		return 500, err
	}

	err = r.db.Delete(&cart, "film_uuid = ? and user_uuid = ?", filmUUID, userUUID).Error
	if err != nil {
		return 500, err
	}
	return 0, nil
}

func (r *Repository) IncreaseQuantity(filmUUID uuid.UUID, userUUID uuid.UUID) (uint64, error) {
	var film ds.Film
	err := r.db.First(&film, filmUUID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 404, err
		}
		return 500, err
	}

	var cart ds.Cart
	err = r.db.First(&cart, "film_uuid = ? and user_uuid = ?", filmUUID, userUUID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			cart.FilmUUID = filmUUID
			cart.UserUUID = userUUID
			cart.Quantity = 0
			err = r.db.Create(cart).Error
			if err != nil {
				return 0, err
			}
		} else {
			return 0, err
		}
	}

	if cart.Quantity == 0 {
		err = r.db.Model(&cart).Where("film_uuid = ? and user_uuid = ?", filmUUID, userUUID).Update("Quantity", cart.Quantity+1).Error
		if err != nil {
			return 0, err
		}
	}

	return cart.Quantity, nil
}

func (r *Repository) DeleteByUser(userUUID uuid.UUID) error {
	var cart ds.Cart
	err := r.db.Where("user_uuid = ?", userUUID).Delete(&cart).Error
	if err != nil {
		return err
	}
	return nil
}
