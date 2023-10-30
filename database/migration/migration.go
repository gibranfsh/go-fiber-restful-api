package migration

import (
	"fmt"
	"go-restapi-fiber/database"
	"go-restapi-fiber/model/entity"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.Book{}, &entity.Author{})

	if err != nil {
		log.Println("Error running migration")
	}

	fmt.Println("Migration run successfully")
}