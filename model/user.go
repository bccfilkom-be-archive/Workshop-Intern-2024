package model

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type UserRegister struct {
	ID       uuid.UUID `json:"-"`
	Name     string    `json:"name" binding:"required"`
	Email    string    `json:"email" binding:"required,email"`
	Password string    `json:"password" binding:"required,min=8"`
	Nim      string    `json:"nim" binding:"required,min=15,max=15"`
	Faculty  string    `json:"faculty" binding:"required"`
	Major    string    `json:"major" binding:"required"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type UserParam struct {
	ID       uuid.UUID `json:"-"`
	Email    string    `json:"-"`
	Password string    `json:"-"`
}

type UserUploadPhoto struct {
	Photo *multipart.FileHeader `form:"photo"`
}
