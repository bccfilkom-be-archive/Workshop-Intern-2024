package main

import (
	"github.com/Ndraaa15/workshop-bcc/rest"
	"github.com/Ndraaa15/workshop-bcc/sdk/config"
)

func main() {
	config.LoadEnv()

	rest := rest.NewRest()

	rest.MountEndpoint()

	rest.Serve()
}
