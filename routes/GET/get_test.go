package get

import (
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/pedrobertao/go-crud/database"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

var app *fiber.App
var fiberContext *fiber.Ctx

func TestMain(m *testing.M) {
	database.Start()
	app = fiber.New()
	fiberContext = app.AcquireCtx(&fasthttp.RequestCtx{})
	code := m.Run()
	os.Exit(code)
}

func TestGetAll(t *testing.T) {
	res := GetAll(fiberContext)
	assert.Nil(t, res)
}
