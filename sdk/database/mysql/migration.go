package mysql

import (
	"github.com/Ndraaa15/workshop-bcc/src/entity"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Book{},
	); err != nil {
		return err
	}

	return nil
}
