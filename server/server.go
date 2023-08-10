package server

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pedrobertao/go-crud/logging"
	auth "github.com/pedrobertao/go-crud/middlewares/authorization"
	delete "github.com/pedrobertao/go-crud/routes/DELETE"
	get "github.com/pedrobertao/go-crud/routes/GET"
	patch "github.com/pedrobertao/go-crud/routes/PATCH"
	post "github.com/pedrobertao/go-crud/routes/POST"
	"github.com/pedrobertao/go-crud/utils"
)

var app *fiber.App

func Start(host string) error {
	app = fiber.New(fiber.Config{
		ServerHeader: "go-crud-template",
		AppName:      "go-crud-template",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 5 * time.Second,
		ErrorHandler: utils.HandleFiberError,
	})

	createRoutes()

	logging.L.Infof("Starting server on: %s", host)
	if err := app.Listen(host); err != nil {
		return err
	}
	return Close()
}

func Close() error {
	logging.L.Infof("Shutting down the server")
	return app.Shutdown()
}

func createRoutes() {
	if app == nil {
		logging.L.Fatal("App not initialized")
	}

	v1 := app.Group("/v1/api")
	{

		v1.Get("/protected", auth.Authorize, func(c *fiber.Ctx) error {
			return c.SendString("Protected")
		})

		v1.Get("/", get.GetAll)

		v1.Get("/:id", get.GetById)

		v1.Post("/", post.PostPerson)

		v1.Delete("/", delete.DeleteByID)

		v1.Patch("/", patch.PatchById)

	}
}
