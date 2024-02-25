package main

import (
	"github.com/Ndraaa15/workshop-bcc/sdk/config"
	"github.com/Ndraaa15/workshop-bcc/src/rest"
)

func main() {
	config.LoadEnv()

	rest := rest.NewRest()

	rest .MountEndpoint()

	rest.Serve()
}
