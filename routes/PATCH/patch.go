package patch

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pedrobertao/go-crud/database"
	"github.com/pedrobertao/go-crud/logging"
	"github.com/pedrobertao/go-crud/models"
)

func PatchById(c *fiber.Ctx) error {
	id := c.Query("id", "")
	var req models.PatchRequest

	if err := c.BodyParser(&req); err != nil {
		logging.L.Error(err)
		return err
	}

	req.ID = id

	if err := req.Validate(); err != nil {
		logging.L.Error(err)
		return err
	}

	person, err := database.PatchById(id, req)
	if err != nil {
		logging.L.Error(err)
		return err
	}
	return c.JSON(person)
}
