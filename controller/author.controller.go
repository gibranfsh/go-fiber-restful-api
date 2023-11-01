package controller

import (
	"github.com/gofiber/fiber/v2"
	"go-restapi-fiber/database"
	"go-restapi-fiber/model/entity"
	"log"
)

func AuthorControllerGetAll (c *fiber.Ctx) error {
	var authors []entity.Author

	res := database.DB.Preload("Books").Find(&authors)

	if res.Error != nil {
		log.Println("Error fetching authors")
	}

	return c.JSON(fiber.Map{
		"message": "Success getting all authors",
		"authors": authors,
	})
}

func AuthorControllerGetById (c *fiber.Ctx) error {
	id := c.Params("id")

	var author entity.Author

	res := database.DB.Find(&author, id)

	if res.Error != nil {
		log.Println("Error fetching author")
	}

	return c.JSON(fiber.Map{
		"message": "Success getting author by id",
		"author": author,
	})
}

func AuthorControllerPost (c *fiber.Ctx) error {
	var author entity.Author

	if err := c.BodyParser(&author); err != nil {
		log.Println("Error parsing author")
	}

	res := database.DB.Create(&author)

	if res.Error != nil {
		log.Println("Error creating author")
	}

	return c.JSON(fiber.Map{
		"message": "Success creating author",
		"author": author,
	})
}

func AuthorControllerPut (c *fiber.Ctx) error {
	id := c.Params("id")

	var author entity.Author

	if err := c.BodyParser(&author); err != nil {
		log.Println("Error parsing author")
	}

	res := database.DB.Model(&author).Where("id = ?", id).Updates(author)

	if res.Error != nil {
		log.Println("Error updating author")
	}

	return c.JSON(fiber.Map{
		"message": "Success updating author",
		"author": author,
	})
}

func AuthorControllerDelete (c *fiber.Ctx) error {
	id := c.Params("id")

	var author entity.Author

	res := database.DB.Delete(&author, id)

	if res.Error != nil {
		log.Println("Error deleting author")
	}

	return c.JSON(fiber.Map{
		"message": "Success deleting author",
		"author": author,
	})
}