package main

import (
	"mini-project-go/config"
	"mini-project-go/constants"
	"mini-project-go/database"
	"mini-project-go/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	config := config.InitConfiguration()
	// database.SeedDB(database.InitDB(config), config)
	e := echo.New()
	e.Static(constants.STATIC_FILE_FOTO_ABSEN, constants.DIR_FILE_FOTO_ABSEN)
	routes.RegisterAuthAPI(e, database.InitDB(config), config)
	routes.RegisterAdminAPI(e, database.InitDB(config), config)
	routes.RegisterPegawaiAPI(e, database.InitDB(config), config)
	e.Logger.Fatal(e.Start(config.SERVER_ADDRESS))
}
