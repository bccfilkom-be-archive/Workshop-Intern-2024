package repository

import (
	"github.com/Ndraaa15/workshop-bcc/entity"
	"gorm.io/gorm"
)

type IRentRepository interface {
	Create(rent entity.Rent) error
}

type RentRepository struct {
	db *gorm.DB
}

func NewRentRepository(db *gorm.DB) IRentRepository {
	return &RentRepository{
		db: db,
	}
}

func (r *RentRepository) Create(rent entity.Rent) error {
	err := r.db.Debug().Create(&rent).Error
	if err != nil {
		return err
	}

	return nil
}
