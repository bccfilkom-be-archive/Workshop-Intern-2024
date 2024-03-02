package mysql

import (
	"log"

	"github.com/Ndraaa15/workshop-bcc/entity"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	// db.Migrator().DropTable(
	// 	&entity.User{},
	// 	&entity.Book{},
	// 	&entity.Rent{},
	// 	&entity.Role{},
	// )

	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Book{},
		&entity.Rent{},
		&entity.Role{},
	); err != nil {
		log.Fatalf("failed migration db: %v", err)
	}
}
