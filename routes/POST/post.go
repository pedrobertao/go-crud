package post

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pedrobertao/go-crud/database"
	"github.com/pedrobertao/go-crud/models"
)

func PostPerson(c *fiber.Ctx) error {
	var person models.Person

	if err := c.BodyParser(&person); err != nil {
		return err
	}

	id, err := database.Post(person)
	if err != nil {
		return err
	}

	return c.JSON(id)

}
