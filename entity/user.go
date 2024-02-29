package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null;"`
	Email     string    `json:"email" gorm:"type:varchar(255);not null;unique"`
	Password  string    `json:"password" gorm:"type:varchar(255);not null;"`
	Nim       string    `json:"nim" gorm:"type:varchar(255);not null;unique"`
	Faculty   string    `json:"faculty" gorm:"type:varchar(255);not null;"`
	Major     string    `json:"major" gorm:"type:varchar(255);not null;"`
	Role      string    `json:"role" gorm:"type:varchar(255);not null;"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Rent      []Rent    `json:"rents"`
}
