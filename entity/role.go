package entity

type Role struct {
	ID   uint   `json:"id" gorm:"primary_key;autoIncrement"`
	Role string `json:"role" gorm:"type:varchar(255);not null;"`
}
