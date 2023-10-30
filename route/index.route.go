package route

import (
	"go-restapi-fiber/controller"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes will setup all routes
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	v1 := api.Group("/v1")

	books := v1.Group("/books")
	books.Get("/", controller.BookControllerGetAll)
	books.Get("/:id", controller.BookControllerGetById)
	books.Post("/", controller.BookControllerPost)

	authors := v1.Group("/authors")
	authors.Get("/", controller.AuthorControllerGetAll)
	authors.Get("/:id", controller.AuthorControllerGetById)
	authors.Post("/", controller.AuthorControllerPost)
}