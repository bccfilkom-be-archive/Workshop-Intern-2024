package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Nim       string    `json:"nim"`
	Faculty   string    `json:"faculty"`
	Major     string    `json:"major"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeleteAt  time.Time `json:"delete_at"`
	Book      []Book    `gorm:"many2many:rents;"`
}
