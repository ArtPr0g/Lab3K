package repository

import (
	"awesomeProject/internal/app/ds"
	"awesomeProject/internal/app/dsn"
	"context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"math/rand"
	"time"
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

func (r *Repository) GetKinoByID(id uint) (*ds.Kino, error) {
	promo := &ds.Kino{}

	err := r.db.First(promo, id).Error
	if err != nil {
		return nil, err
	}

	return promo, nil
}

func (r *Repository) NewRandRecords() error {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(900000) + 100000
	release := rand.Intn(9990) + 10
	grade := rand.Intn(10) + 1
	storeList := []string{"Терминатор", "Бэтмен", "Челове паук", "Летнее время", "Горько", "Токийский гуль"}
	storeRandom := rand.Intn(len(storeList))
	store := storeList[storeRandom]
	new := ds.Kino{
		Code:    uint(code), // код от 100000 до 999999
		Release: release,    // цена от 10 до 999
		Grade:   grade,      //промо даёт в 2 раза больше цены промо
		Name:    store,
	}
	err := r.db.Create(&new).Error
	if err != nil {
		return err
	}
	return nil
}
