package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository IUserRepository
	BookRepository IBookRepository
	RentRepository IRentRepository
}

func NewRepository(db *gorm.DB) *Repository {
	userRepository := NewUserRepository(db)
	bookRepository := NewBookRepository(db)
	rentRepository := NewRentRepository(db)

	return &Repository{
		UserRepository: userRepository,
		BookRepository: bookRepository,
		RentRepository: rentRepository,
	}
}
