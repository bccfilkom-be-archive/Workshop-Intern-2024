package rest

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Ndraaa15/workshop-bcc/entity"
	"github.com/Ndraaa15/workshop-bcc/internal/service"
	"github.com/Ndraaa15/workshop-bcc/pkg/middleware"
	"github.com/Ndraaa15/workshop-bcc/pkg/response"
	"github.com/gin-gonic/gin"
)

type Rest struct {
	router     *gin.Engine
	service    *service.Service
	middleware middleware.Interface
}

func NewRest(service *service.Service, middleware middleware.Interface) *Rest {
	return &Rest{
		router:     gin.Default(),
		service:    service,
		middleware: middleware,
	}
}

func (r *Rest) MountEndpoint() {
	r.router.Use(r.middleware.Timeout())

	routerGroup := r.router.Group("/api/v1")

	routerGroup.GET("/health-check", healthCheck)

	routerGroup.GET("/time-out", testTimeout)

	routerGroup.GET("/login-user", r.middleware.AuthenticateUser, getLoginUser)

	routerGroup.POST("/register", r.Register)
	routerGroup.POST("/login", r.Login)

	user := routerGroup.Group("/user")
	user.GET("/rent", r.middleware.AuthenticateUser, r.GetUserRentBook)

	book := routerGroup.Group("/book")
	book.POST("/", r.CreateBook)
	book.GET("/:id", r.GetBookByID)
	book.DELETE("/:id", r.DeleteBook)
	book.PATCH("/:id", r.UpdateBook)
	book.GET("/", r.GetAllBook)

	rent := routerGroup.Group("/rent", r.middleware.AuthenticateUser)
	rent.POST("/", r.RentBook)
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

func testTimeout(ctx *gin.Context) {
	time.Sleep(3 * time.Second)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func getLoginUser(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		response.Error(ctx, http.StatusInternalServerError, "failed get login user", errors.New(""))
		return
	}

	response.Success(ctx, http.StatusOK, "get login user", user.(entity.User))
}
