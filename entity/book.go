package entity

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID          uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;"`
	Title       string    `json:"title" gorm:"type:varchar(255);not null;unique"`
	Writter     string    `json:"writter" gorm:"type:varchar(255);not null;"`
	Year        uint      `json:"year" gorm:"type:int;not null;"`
	Genre       string    `json:"genre" gorm:"type:varchar(255);not null;"`
	Description string    `json:"description" gorm:"type:text;not null;"`
	Stock       uint      `json:"stock" gorm:"type:int unsigned;not null;"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Rent        []Rent    `json:"rents"`
}
