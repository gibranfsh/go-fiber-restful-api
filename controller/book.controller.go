package controller

import (
	"github.com/gofiber/fiber/v2"
	"go-restapi-fiber/database"
	"go-restapi-fiber/model/entity"
	"log"
)

func BookControllerGetAll(c *fiber.Ctx) error {
	var books []entity.Book

	res := database.DB.Preload("Author").Find(&books) // Preload("Author") untuk mengambil data author dari tabel author

	if res.Error != nil {
		log.Println("Error fetching books")
	}

	return c.JSON(fiber.Map{
		"message": "Success getting all books",
		"books":   books,
	})
}

func BookControllerGetById(c *fiber.Ctx) error {
	id := c.Params("id")

	var book entity.Book

	res := database.DB.Preload("Author").Find(&book, id) // Preload("Author") untuk mengambil data author dari tabel author

	if res.Error != nil {
		log.Println("Error fetching book")
	}

	return c.JSON(fiber.Map{
		"message": "Success getting book by id",
		"book":    book,
	})
}

func BookControllerPost(c *fiber.Ctx) error {
	var book entity.Book

	if err := c.BodyParser(&book); err != nil {
		log.Println("Error parsing book")
	}

	res := database.DB.Create(&book)

	if res.Error != nil {
		log.Println("Error creating book")
	}

	return c.JSON(fiber.Map{
		"message": "Success creating book",
		"book":    book,
	})
}

func BookControllerPut(c *fiber.Ctx) error {
	id := c.Params("id")

	var book entity.Book

	res := database.DB.Find(&book, id)

	if res.Error != nil {
		log.Println("Error fetching book")
	}

	if err := c.BodyParser(&book); err != nil {
		log.Println("Error parsing book")
	}

	res = database.DB.Save(&book)

	if res.Error != nil {
		log.Println("Error updating book")
	}

	return c.JSON(fiber.Map{
		"message": "Success updating book",
		"book":    book,
	})
}

func BookControllerDelete(c *fiber.Ctx) error {
	id := c.Params("id")

	var book entity.Book

	res := database.DB.Find(&book, id)

	if res.Error != nil {
		log.Println("Error fetching book")
	}

	res = database.DB.Delete(&book)

	if res.Error != nil {
		log.Println("Error deleting book")
	}

	return c.JSON(fiber.Map{
		"message": "Success deleting book",
		"book":    book,
	})
}

// gibran CR, iyal UD
