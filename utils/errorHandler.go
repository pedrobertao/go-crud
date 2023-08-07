package utils

import (
	"reflect"

	"github.com/gofiber/fiber/v2"
)

const VALIDATOR_ERROR string = "validator.ValidationErrors"

func HandleFiberError(c *fiber.Ctx, err error) error {
	t := reflect.TypeOf(err)
	if t.String() == VALIDATOR_ERROR {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	} else {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
}
