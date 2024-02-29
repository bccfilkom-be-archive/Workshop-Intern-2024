package rest

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Ndraaa15/workshop-bcc/internal/service"
	"github.com/gin-gonic/gin"
)

type Rest struct {
	router  *gin.Engine
	service *service.Service
}

func NewRest(service *service.Service) *Rest {
	return &Rest{
		router:  gin.Default(),
		service: service,
	}
}

func (r *Rest) MountEndpoint() {
	routerGroup := r.router.Group("/api/v1")

	routerGroup.GET("/health-check", healthCheck)

	book := routerGroup.Group("/book")
	book.POST("/", r.CreateBook)
	book.GET("/:id", r.GetBookByID)
	book.DELETE("/:id", r.DeleteBook)
	book.PATCH("/:id", r.UpdateBook)
	book.GET("/", r.GetAllBook)
}

func (r *Rest) Serve() {
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
