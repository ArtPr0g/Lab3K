package repository

import (
	"awesomeProject/internal/app/ds"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *Repository) GetFilms() ([]ds.Film, error) {

	var films []ds.Film
	err := r.db.Order("uuid").Find(&films).Error
	return films, err
}

func (r *Repository) GetFilm(uuid uuid.UUID) (ds.Film, error) {
	var film ds.Film
	err := r.db.First(&film, uuid).Error
	return film, err
}

func (r *Repository) GetFilmName(uuid uuid.UUID) (string, error) {
	var film ds.Film
	err := r.db.Select("name").First(&film, "uuid = ?", uuid).Error
	return film.Name, err
}

func (r *Repository) CreateFilm(film ds.Film) error {
	err := r.db.Create(&film).Error
	return err
}

func (r *Repository) ChangeFilm(uuid uuid.UUID, film ds.Film) (int, error) {
	film.UUID = uuid

	err := r.db.Model(&film).Updates(ds.Film{Name: film.Name, Release: film.Release, Grade: film.Grade, Genre: film.Genre, Price: film.Price, WhatchTime: film.WhatchTime, Summary: film.Summary, Video: film.Video, Image: film.Image}).Error
	if err != nil {
		return 500, err
	}

	return 0, nil
}

func (r *Repository) DeleteFilm(uuid uuid.UUID) (int, error) {
	var film ds.Film
	err := r.db.First(&film, uuid).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 404, err
		}
		return 500, err
	}
	err = r.db.Delete(&film, uuid).Error
	if err != nil {
		return 500, err
	}
	return 0, nil
}

func (r *Repository) GetVideo(quantity uint64, filmUUID uuid.UUID, userUUID uuid.UUID) (int, string, error) {
	var film ds.Film
	err := r.db.First(&film, filmUUID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 404, "", err
		}
		return 500, "", err
	}

	var cart ds.Cart
	err = r.db.First(&cart, filmUUID, userUUID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 404, "", err
		}
		return 500, "", err
	}
	err = r.db.Delete(&cart, filmUUID, userUUID).Error
	if err != nil {
		return 500, "", err
	}

	err = r.AddOrder(userUUID, filmUUID, quantity)
	if err != nil {
		return 500, "", err
	}

	PromoString := film.Video
	return 0, PromoString, nil
}
