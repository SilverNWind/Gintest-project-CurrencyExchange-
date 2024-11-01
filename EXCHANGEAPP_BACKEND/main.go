package main

import (
	"exchangeapp/config"
	"exchangeapp/router"
)

func main() {
	config.InitConfig()

	r := router.SetupRouter()

	port := config.Appconfig.App.Port

	if port == "" {
		port = ":8080"
	}

	r.Run(port)
}