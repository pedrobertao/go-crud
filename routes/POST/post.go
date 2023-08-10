package post

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pedrobertao/go-crud/database"
	"github.com/pedrobertao/go-crud/logging"
	"github.com/pedrobertao/go-crud/models"
)

func PostPerson(c *fiber.Ctx) error {
	var req models.PostRequest

	if err := c.BodyParser(&req); err != nil {
		logging.L.Error(err)
		return err
	}

	if err := req.Validate(); err != nil {
		logging.L.Error(err)
		return err
	}

	id, err := database.Post(req)
	if err != nil {
		logging.L.Error(err)
		return err
	}

	return c.JSON(id)

}
