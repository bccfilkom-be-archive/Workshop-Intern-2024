package rest

import (
	"github.com/Ndraaa15/workshop-bcc/internal/service"
)

type UserHandler struct {
	us service.IUserService
}

func NewUserHandler(us service.IUserService) *UserHandler {
	return &UserHandler{
		us: us,
	}
}
