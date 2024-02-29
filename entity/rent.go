package entity

import "github.com/google/uuid"

type Rent struct {
	ID        uint      `json:"id" gorm:"primary_key;autoIncrement"`
	UserID    uuid.UUID `json:"userId" gorm:"type:varchar(36);primary_key;foreignkey:ID;references:users;onUpdate:CASCADE;onDelete:CASCADE"`
	BookID    uuid.UUID `json:"bookId" gorm:"type:varchar(36);primary_key;foreignkey:ID;references:books;onUpdate:CASCADE;onDelete:CASCADE"`
	Total     int       `json:"total"`
	IsReturn  bool      `json:"isReturn"`
	ReturnAt  string    `json:"returnAt"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeleteAt  string    `json:"deleteAt"`
}
