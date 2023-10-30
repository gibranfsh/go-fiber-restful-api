package main

import (
	"go-restapi-fiber/database"
	"go-restapi-fiber/database/migration"
	"go-restapi-fiber/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.InitDatabase()
	migration.RunMigration()
	
	app := fiber.New()

	route.SetupRoutes(app)

	app.Listen(":3000")
}
