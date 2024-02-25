package handler

import "github.com/Ndraaa15/workshop-bcc/src/internal/service"

type Handler struct {
	UserHandler *UserHandler
}

func NewHandler(userService service.IUserService) *Handler {
	return &Handler{
		UserHandler: NewUserHandler(userService),
	}

}
