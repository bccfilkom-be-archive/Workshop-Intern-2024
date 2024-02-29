package model

import "github.com/google/uuid"

type UserRegister struct {
	ID       uuid.UUID `json:"-"`
	Name     string    `json:"name" binding:"required"`
	Email    string    `json:"email" binding:"required,email"`
	Password string    `json:"password" binding:"required,min=8"`
	Nim      string    `json:"nim" binding:"required,min=15,max=15"`
	Faculty  string    `json:"faculty" binding:"required"`
	Major    string    `json:"major" binding:"required"`
}

type UserRegisterResponse struct {
	Token string `json:"token"`
}
