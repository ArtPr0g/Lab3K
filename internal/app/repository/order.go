package repository

import (
	"awesomeProject/internal/app/ds"
	"errors"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/url"
	"time"
)

func (r *Repository) AddOrder(userUUID uuid.UUID, filmUUID uuid.UUID, quantity uint64) error {
	var err error
	var order ds.Order

	order.Film, err = r.GetFilmName(filmUUID)
	if err != nil {
		return err
	}

	order.Quantity = quantity

	order.UserUUID = userUUID

	order.Date = time.Now()

	order.Status = "Заказан"

	err = r.db.Create(&order).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetOrders(stDate, endDate, status string) ([]ds.Order, error) {
	var orders []ds.Order
	var err error
	st, _ := url.QueryUnescape(status)
	log.Println(st)
	if st == "Любой" {
		if stDate == "" && endDate == "" {
			err = r.db.Order("date").Find(&orders).Error
			return orders, err
		} else if stDate != "" && endDate == "" {
			err = r.db.Order("date").Where("date > ?", stDate).Find(&orders).Error
			return orders, err
		} else if stDate == "" && endDate != "" {
			err = r.db.Order("date").Where("date < ?", endDate).Find(&orders).Error
			return orders, err
		} else if stDate != "" && endDate != "" {
			err = r.db.Order("date").Where("date > ? and date < ?", stDate, endDate).Find(&orders).Error
			return orders, err
		}
	} else {
		if stDate == "" && endDate == "" {
			err = r.db.Order("date").Where("status = ?", st).Find(&orders).Error
			return orders, err
		} else if stDate != "" && endDate == "" {
			err = r.db.Order("date").Where("date > ? and status = ?", stDate, st).Find(&orders).Error
			return orders, err
		} else if stDate == "" && endDate != "" {
			err = r.db.Order("date").Where("date < ? and status = ?", endDate, st).Find(&orders).Error
			return orders, err
		} else if stDate != "" && endDate != "" {
			err = r.db.Order("date").Where("date > ? and date < ? and status = ?", stDate, endDate, st).Find(&orders).Error
			return orders, err
		}
	}

	return orders, nil
}

func (r *Repository) ChangeStatus(uuid uuid.UUID, status string) (int, error) {
	var order ds.Order
	err := r.db.First(&order, "uuid = ?", uuid).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 404, err
		}
		return 500, err
	}
	err = r.db.Model(&order).Update("Status", status).Error
	//if errors.Is(err, gorm.ErrRecordNotFound)
	if err != nil {
		return 500, err
	}
	return 0, nil
}
