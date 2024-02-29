package rest

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Ndraaa15/workshop-bcc/entity"
	"github.com/Ndraaa15/workshop-bcc/internal/handler"
	"github.com/Ndraaa15/workshop-bcc/internal/repository"
	"github.com/Ndraaa15/workshop-bcc/internal/service"
	"github.com/Ndraaa15/workshop-bcc/pkg/database/mysql"
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

	routerGroup := r.router.Group("/api/v1")

	routerGroup.GET("/health-check", healthCheck)

	book := routerGroup.Group("/book")
	book.POST("/", handler.BookHandler.CreateBook)
	book.GET("/:id", handler.BookHandler.GetBookByID)
	book.DELETE("/:id", handler.BookHandler.DeleteBook)
	book.PATCH("/:id", handler.BookHandler.UpdateBook)
	book.GET("/", handler.BookHandler.GetAllBook)

}

func (r *Rest) Serve() {
	mysql.Migration(r.db)

	var totalBook int64
	if err := r.db.Model(&entity.Book{}).Count(&totalBook).Error; err != nil {
		log.Fatalf("Error while counting book: %v", err)
		return
	}

	if totalBook == 0 {
		if err := mysql.GenerateBook(r.db); err != nil {
			log.Fatalf("Error while generating book: %v", err)
			return
		}
	}

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
