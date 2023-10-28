package main

import (
	"github.com/gofiber/fiber/v2"
	"go-restapi-fiber/database"
	"go-restapi-fiber/route"
)

func main() {
	database.InitDatabase()
	app := fiber.New()

	route.SetupRoutes(app)

	app.Listen(":3000")
}
