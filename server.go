package main

import (
	"fmt"
	_ "fmt"
	"log"
	"technexRegistration/database"
	"technexRegistration/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	fmt.Println("Using CORS")
	app.Use(cors.New())
	fmt.Println("Initialising Routes")
	router.Route(app)
	fmt.Println("Init DB connections")
	if err := database.Init(); err != nil {
		log.Fatal("unable to connect to client")
	}
	defer database.Disconnect()
	app.Listen(":6969")
	log.Println("Server started on port 6969")
}
