package entity

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID          uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;"`
	Name        string    `json:"name"`
	Writter     string    `json:"writter"`
	Year        uint      `json:"year"`
	Genre       string    `json:"genre"`
	Description string    `json:"description"`
	Stock       uint      `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeleteAt    time.Time `json:"delete_at"`
	User        []User    `gorm:"many2many:rents;"`
}
