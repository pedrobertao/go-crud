package server

import (
	"time"

	"github.com/gofiber/fiber/v2"
	delete "github.com/pedrobertao/go-crud/routes/DELETE"
	get "github.com/pedrobertao/go-crud/routes/GET"
	patch "github.com/pedrobertao/go-crud/routes/PATCH"
	post "github.com/pedrobertao/go-crud/routes/POST"
	"github.com/pedrobertao/go-crud/utils"
)

var app *fiber.App

func Start(portOrHost string) error {
	app = fiber.New(fiber.Config{
		ServerHeader: "go-crud-template",
		AppName:      "go-crud-template",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 5 * time.Second,
		ErrorHandler: utils.HandleFiberError,
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

		v1.Delete("/", delete.DeleteByID)

		v1.Patch("/", patch.PatchById)
	}

	return nil
}
