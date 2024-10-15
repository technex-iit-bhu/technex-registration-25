package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"technexRegistration/handler"
)

func Route(app *fiber.App) {

	api := app.Group("/api", logger.New())
	app.Use(cors.New())
	api.Get("/", handler.Hello)
}
