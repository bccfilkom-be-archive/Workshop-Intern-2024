package entity

import "github.com/google/uuid"

type Rent struct {
	ID        uint      `json:"id" gorm:"primary_key;autoIncrement"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:varchar(36);primary_key;foreignkey:ID;references:users;onUpdate:CASCADE;onDelete:CASCADE"`
	BookID    uuid.UUID `json:"book_id" gorm:"type:varchar(36);primary_key;foreignkey:ID;references:books;onUpdate:CASCADE;onDelete:CASCADE"`
	Total     int       `json:"total"`
	IsReturn  bool      `json:"is_return"`
	ReturnAt  string    `json:"return_at"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	DeleteAt  string    `json:"delete_at"`
}
