package main

import (
	"github.com/Ndraaa15/workshop-bcc/pkg/config"
	"github.com/Ndraaa15/workshop-bcc/rest"
)

func main() {
	config.LoadEnv()

	rest := rest.NewRest()

	rest.MountEndpoint()

	rest.Serve()
}
