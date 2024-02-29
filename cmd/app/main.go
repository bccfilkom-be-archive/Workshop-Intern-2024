package main

import (
	"github.com/Ndraaa15/workshop-bcc/internal/handler/rest"
	"github.com/Ndraaa15/workshop-bcc/internal/repository"
	"github.com/Ndraaa15/workshop-bcc/internal/service"
	"github.com/Ndraaa15/workshop-bcc/pkg/bcrypt"
	"github.com/Ndraaa15/workshop-bcc/pkg/config"
	"github.com/Ndraaa15/workshop-bcc/pkg/database/mysql"
	"github.com/Ndraaa15/workshop-bcc/pkg/jwt"
	"github.com/Ndraaa15/workshop-bcc/pkg/middleware"
)

func main() {
	config.LoadEnv()

	jwtAuth := jwt.Init()

	bcrypt := bcrypt.Init()

	db := mysql.ConnectDatabase()

	repository := repository.NewRepository(db)

	service := service.NewService(service.InitParam{Repository: repository, JwtAuth: jwtAuth, Bcrypt: bcrypt})

	middleware := middleware.Init(jwtAuth, service)

	rest := rest.NewRest(service, middleware)

	mysql.Migration(db)

	mysql.SeedData(db)

	rest.MountEndpoint()

	rest.Serve()
}
