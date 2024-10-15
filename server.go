package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"technexRegistration/router"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	router.Route(app)
	app.Listen(":6969")
	log.Println("Server started on port 6969")
}
