package main

import (
	"github.com/naomihc/echo-rest/config"
	"github.com/naomihc/echo-rest/database"
	"github.com/naomihc/echo-rest/routes"
)

func main() {
	database.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":" + config.GetApplicationPort()))
}
