package service

import (
	"github.com/Ndraaa15/workshop-bcc/internal/repository"
	"github.com/Ndraaa15/workshop-bcc/pkg/bcrypt"
	"github.com/Ndraaa15/workshop-bcc/pkg/jwt"
)

type Service struct {
	UserService IUserService
	BookService IBookService
	RentService IRentService
}

type InitParam struct {
	Repository *repository.Repository
	JwtAuth    jwt.Interface
	Bcrypt     bcrypt.Interface
}

func NewService(param InitParam) *Service {
	userService := NewUserService(param.Repository.UserRepository, param.Bcrypt, param.JwtAuth)
	bookService := NewBookService(param.Repository.BookRepository)
	rentService := NewRentService(param.Repository.RentRepository, param.Repository.UserRepository, param.Repository.BookRepository)

	return &Service{
		UserService: userService,
		BookService: bookService,
		RentService: rentService,
	}
}
