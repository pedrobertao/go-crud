package get

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pedrobertao/go-crud/database"
	"github.com/pedrobertao/go-crud/logging"
)

func GetById(c *fiber.Ctx) error {
	id := c.Params("id")
	res, err := database.FindById(id)
	if err != nil {
		logging.L.Error(err)
		return err
	}

	if res.ID == "" {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(res)
}

func GetAll(c *fiber.Ctx) error {
	res, err := database.FindAll()
	if err != nil {
		return err
	}
	return c.JSON(res)
}
