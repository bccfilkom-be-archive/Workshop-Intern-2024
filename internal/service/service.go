package service

import "github.com/Ndraaa15/workshop-bcc/internal/repository"

type Service struct {
	UserService IUserService
	BookService IBookService
}

func NewService(repository *repository.Repository) *Service {
	userService := NewUserService(repository.UserRepository)
	bookService := NewBookService(repository.BookRepository)

	return &Service{
		UserService: userService,
		BookService: bookService,
	}
}
