package patch

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pedrobertao/go-crud/database"
	"github.com/pedrobertao/go-crud/models"
)

func PatchById(c *fiber.Ctx) error {
	id := c.Query("id", "")
	var patchReq models.PatchRequest

	if err := c.BodyParser(&patchReq); err != nil {
		return err
	}

	person, err := database.PatchById(id, patchReq)
	if err != nil {
		return err
	}
	return c.JSON(person)
}
