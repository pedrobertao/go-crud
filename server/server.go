package server

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

var app *fiber.App

func Start(portOrHost string) error {
	app = fiber.New(fiber.Config{
		ServerHeader: "go-crud-template",
		AppName:      "go-crud-template",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 5 * time.Second,
		ErrorHandler: func(c *fiber.Ctx, err error) error {

			return err
		},
	})

	createMiddlewares()

	createRoutes()

	err := app.Listen(portOrHost)
	if err != nil {
		return err
	}
	return nil
}

func createMiddlewares() {

}

func createRoutes() fiber.Router {
	if app == nil {
		return nil
	}

	v1 := app.Group("/v1/api")
	{
		v1.Get("/", func(c *fiber.Ctx) error {
			return nil
		})
		v1.Post("/", func(c *fiber.Ctx) error {
			return nil
		})
		v1.Delete("/", func(c *fiber.Ctx) error {
			return nil
		})
		v1.Put("/", func(c *fiber.Ctx) error {
			return nil
		})
	}

	return nil
}
