package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository IUserRepository
}

func NewRepository(db *gorm.DB) *Repository {
	userRepository := NewUserRepository(db)

	return &Repository{
		UserRepository: userRepository,
	}
}
