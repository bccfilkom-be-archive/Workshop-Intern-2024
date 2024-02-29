package main

import (
	"github.com/Ndraaa15/workshop-bcc/internal/handler/rest"
	"github.com/Ndraaa15/workshop-bcc/internal/repository"
	"github.com/Ndraaa15/workshop-bcc/internal/service"
	"github.com/Ndraaa15/workshop-bcc/pkg/config"
	"github.com/Ndraaa15/workshop-bcc/pkg/database/mysql"
)

func main() {
	config.LoadEnv()

	db := mysql.ConnectDatabase()

	repository := repository.NewRepository(db)

	service := service.NewService(repository)

	rest := rest.NewRest(service)

	mysql.Migration(db)

	mysql.SeedData(db)

	rest.MountEndpoint()

	rest.Serve()
}
