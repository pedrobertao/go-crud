package delete

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pedrobertao/go-crud/database"
	"github.com/pedrobertao/go-crud/logging"
	"github.com/pedrobertao/go-crud/models"
)

func DeleteByID(c *fiber.Ctx) error {
	id := c.Query("id", "")

	req := models.DeleteRequest{
		ID: id,
	}

	if err := req.Validate(); err != nil {
		logging.L.Error(err)
		return err
	}

	count := database.DeleteByID(id)
	if count > 0 {
		return c.JSON(id)
	}
	return c.SendStatus(fiber.StatusNotFound)
}
