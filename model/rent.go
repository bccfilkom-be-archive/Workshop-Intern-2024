package model

import "github.com/google/uuid"

type RentBook struct {
	UserID uuid.UUID `json:"userId" binding:"required,uuid"`
	BookID uuid.UUID `json:"bookId" binding:"required,uuid"`
	Total  int       `json:"total" binding:"required,numeric"`
}
