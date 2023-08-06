package delete

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pedrobertao/go-crud/database"
)

func DeleteByID(c *fiber.Ctx) error {
	id := c.Query("id", "")
	count := database.DeleteByID(id)
	if count > 0 {
		return c.JSON(id)
	}
	return c.SendStatus(fiber.StatusNotFound)
}
