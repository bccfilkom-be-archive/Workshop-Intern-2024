package mysql

import (
	"log"

	"github.com/Ndraaa15/workshop-bcc/sdk/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.LoadDatabaseConfig()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Error while connecting to database: %v", err)
		return nil
	}

	return db
}