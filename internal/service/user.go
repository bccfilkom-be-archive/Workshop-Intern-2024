package service

import (
	"github.com/Ndraaa15/workshop-bcc/entity"
	"github.com/Ndraaa15/workshop-bcc/internal/repository"
	"github.com/Ndraaa15/workshop-bcc/model"
	"github.com/Ndraaa15/workshop-bcc/pkg/bcrypt"
	"github.com/Ndraaa15/workshop-bcc/pkg/jwt"
	"github.com/google/uuid"
)

type IUserService interface {
	Register(param model.UserRegister) (model.UserRegisterResponse, error)
}

type UserService struct {
	ur      repository.IUserRepository
	bcrypt  bcrypt.Interface
	jwtAuth jwt.Interface
}

func NewUserService(userRepository repository.IUserRepository, bcrypt bcrypt.Interface, jwtAuth jwt.Interface) IUserService {
	return &UserService{
		ur:      userRepository,
		bcrypt:  bcrypt,
		jwtAuth: jwtAuth,
	}
}

func (u *UserService) Register(param model.UserRegister) (model.UserRegisterResponse, error) {
	var result model.UserRegisterResponse

	hashPassword, err := u.bcrypt.GenerateFromPassword(param.Password)
	if err != nil {
		return result, err
	}

	param.ID = uuid.New()
	param.Password = hashPassword

	user := entity.User{
		ID:       param.ID,
		Name:     param.Name,
		Email:    param.Email,
		Password: param.Password,
		Nim:      param.Nim,
		Faculty:  param.Faculty,
		Major:    param.Major,
		Role:     2,
	}

	user, err = u.ur.CreateUser(user)
	if err != nil {
		return result, err
	}

	token, err := u.jwtAuth.CreateJWTToken(user.ID.String())
	if err != nil {
		return result, err
	}

	result.Token = token

	return result, nil
}
