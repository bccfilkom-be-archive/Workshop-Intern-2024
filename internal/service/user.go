package service

import (
	"github.com/Ndraaa15/workshop-bcc/entity"
	"github.com/Ndraaa15/workshop-bcc/internal/repository"
	"github.com/Ndraaa15/workshop-bcc/model"
	"github.com/Ndraaa15/workshop-bcc/pkg/bcrypt"
	"github.com/Ndraaa15/workshop-bcc/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IUserService interface {
	Register(param model.UserRegister) error
	GetUser(param model.UserParam) (entity.User, error)
	Login(param model.UserLogin) (model.UserLoginResponse, error)
	GetUserRentBook(ctx *gin.Context) (entity.User, error)
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

func (u *UserService) Register(param model.UserRegister) error {
	hashPassword, err := u.bcrypt.GenerateFromPassword(param.Password)
	if err != nil {
		return err
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

	_, err = u.ur.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) Login(param model.UserLogin) (model.UserLoginResponse, error) {
	result := model.UserLoginResponse{}

	user, err := u.ur.GetUser(model.UserParam{
		Email: param.Email,
	})
	if err != nil {
		return result, err
	}

	err = u.bcrypt.CompareAndHashPassword(user.Password, param.Password)
	if err != nil {
		return result, err
	}

	token, err := u.jwtAuth.CreateJWTToken(user.ID)
	if err != nil {
		return result, err
	}

	result.Token = token

	return result, nil
}

func (u *UserService) GetUser(param model.UserParam) (entity.User, error) {
	return u.ur.GetUser(param)
}

func (u *UserService) GetUserRentBook(ctx *gin.Context) (entity.User, error) {
	user, err := u.jwtAuth.GetLoginUser(ctx)
	if err != nil {
		return user, err
	}

	param := model.UserParam{
		ID: user.ID,
	}

	return u.ur.GetUserWithRent(param)
}
