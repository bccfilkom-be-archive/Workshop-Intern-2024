package handler

import (
	"github.com/Ndraaa15/workshop-bcc/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	us service.IUserService
}

func NewUserHandler(us service.IUserService) *UserHandler {
	return &UserHandler{
		us: us,
	}
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {

}
