package server

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	get "github.com/pedrobertao/go-crud/routes/GET"
	post "github.com/pedrobertao/go-crud/routes/POST"
)

var app *fiber.App

func Start(portOrHost string) error {
	app = fiber.New(fiber.Config{
		ServerHeader: "go-crud-template",
		AppName:      "go-crud-template",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 5 * time.Second,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			log.Println(err)
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
		v1.Get("/", get.GetAll)

		v1.Get("/:id", get.GetById)

		v1.Post("/", post.PostPerson)

		v1.Delete("/", func(c *fiber.Ctx) error {
			return nil
		})
		v1.Patch("/", func(c *fiber.Ctx) error {
			return nil
		})
	}

	return nil
}
