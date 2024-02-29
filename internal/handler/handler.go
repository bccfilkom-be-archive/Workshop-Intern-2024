package handler

import "github.com/Ndraaa15/workshop-bcc/internal/service"

type Handler struct {
	UserHandler *UserHandler
	BookHandler *BookHandler
}

func NewHandler(service *service.Service) *Handler {
	userHandler := NewUserHandler(service.UserService)
	bookHandler := NewBookHandler(service.BookService)

	return &Handler{
		UserHandler: userHandler,
		BookHandler: bookHandler,
	}
}
