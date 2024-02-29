package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository IUserRepository
	BookRepository IBookRepository
}

func NewRepository(db *gorm.DB) *Repository {
	userRepository := NewUserRepository(db)
	bookRepository := NewBookRepository(db)

	return &Repository{
		UserRepository: userRepository,
		BookRepository: bookRepository,
	}
}
