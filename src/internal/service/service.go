package service

import "github.com/Ndraaa15/workshop-bcc/src/internal/repository"

type Service struct {
	userRepository repository.IUserRepository
}

func NewService(repository *repository.Repository) *Service {
	userRepository := NewUserService(repository.UserRepository)

	return &Service{
		userRepository: userRepository,
	}
}
