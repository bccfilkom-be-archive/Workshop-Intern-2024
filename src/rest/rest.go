package rest

import (
	"fmt"
	"log"
	"net/http"
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

func NewRest() *Rest {
	db := mysql.ConnectDatabase()

	// global middleware goes here

	return &Rest{
		router: gin.Default(),
		db:     db,
	}
}

func (r *Rest) MountEndpoint() {
	repository := repository.NewRepository(r.db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	r.router.GET("/health-check", healthCheck)

	r.router.POST("/users", handler.UserHandler.CreateUser)

}

func (r *Rest) Serve() {
	mysql.Migration(r.db)

	addr := os.Getenv("APP_ADDRESS")
	port := os.Getenv("APP_PORT")

	err := r.router.Run(fmt.Sprintf("%s:%s", addr, port))
	if err != nil {
		log.Fatalf("Error while serving: %v", err)
	}

}

func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
