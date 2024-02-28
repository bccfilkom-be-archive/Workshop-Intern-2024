package mysql

import (
	"log"

	"github.com/Ndraaa15/workshop-bcc/src/entity"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	// db.Migrator().DropTable(
	// 	&entity.User{},
	// 	&entity.Book{},
	// 	&entity.Rent{},
	// )

	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Book{},
		&entity.Rent{},
	); err != nil {
		log.Fatalf("failed migration db: %v", err)
	}
}
