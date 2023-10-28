package route

import (
	"go-restapi-fiber/controller"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes will setup all routes
func SetupRoutes(app *fiber.App) {
	app.Get("/user", controller.UserControllerRead)
}