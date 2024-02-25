package rest

import (
	"fmt"
	"log"
	"os"

	"github.com/Ndraaa15/workshop-bcc/sdk/database/mysql"
	"github.com/Ndraaa15/workshop-bcc/src/internal/handler"
	"github.com/Ndraaa15/workshop-bcc/src/internal/repository"
	"github.com/Ndraaa15/workshop-bcc/src/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Rest struct {
	router *gin.Engine
	db     *gorm.DB
}

func NewRest() (*Rest, error) {
	db, err := mysql.ConnectDatabase()
	if err != nil {
		return nil, err
	}

	return &Rest{
		router: gin.Default(),
		db:     db,
	}, nil
}

func (r *Rest) MountEndpoint() {
	repository := repository.NewRepository(r.db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	r.router.POST("/users", handler.UserHandler.CreateUser)

}

func (r *Rest) Serve() error {
	addr := os.Getenv("APP_ADDRESS")
	port := os.Getenv("APP_PORT")

	err := r.router.Run(fmt.Sprintf("%s:%s", addr, port))
	if err != nil {
		log.Printf("Error while serving: %v", err)
		return err
	}

	return nil
}
