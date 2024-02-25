package service

import (
	"github.com/Ndraaa15/workshop-bcc/src/internal/repository"
)

type IUserService interface {
}

type UserService struct {
	ur repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &UserService{
		ur: userRepository,
	}
}
