package utils

import (
	"reflect"

	"github.com/gofiber/fiber/v2"
)

const VALIDATOR_ERROR string = "validator.ValidationErrors"

func HandleFiberError(c *fiber.Ctx, err error) error {
	t := reflect.TypeOf(err)
	if t.String() == VALIDATOR_ERROR {
		c.Status(fiber.StatusBadRequest)
		return c.SendString(err.Error())
	} else {
		c.Status(fiber.StatusInternalServerError)
		return c.SendString(err.Error())
	}
}
