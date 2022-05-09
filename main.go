package main

import (
	"mini-project-go/config"
	"mini-project-go/database"
	"mini-project-go/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	config := config.InitConfiguration()
	// database.SeedDB(database.InitDB(config), config)
	e := echo.New()
	routes.RegisterAuthAPI(e, database.InitDB(config), config)
	routes.RegisterAdminAPI(e, database.InitDB(config), config)
	e.Logger.Fatal(e.Start(config.SERVER_ADDRESS))
}
