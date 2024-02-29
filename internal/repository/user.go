package repository

import (
	"github.com/Ndraaa15/workshop-bcc/entity"
	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(user entity.User) (entity.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) CreateUser(user entity.User) (entity.User, error) {
	err := u.db.Debug().Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
