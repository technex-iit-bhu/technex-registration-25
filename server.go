package main

import (
	"log"
	"os"
	"technexRegistration/config"
	"technexRegistration/database"
	"technexRegistration/router"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	// app.Use(cors.New())
	router.Route(app)

	// Validate required environment variables
	requiredEnvVars := []string{"JWT_SECRET", "GMAIL_SECRET", "RECOVERY_SECRET", "QR_SECRET", "REFRESH_SECRET"}
	for _, envVar := range requiredEnvVars {
		if config.Config(envVar) == "" {
			log.Fatalf("Required environment variable %s is not set", envVar)
		}
	}

	if err := database.Init(); err != nil {
		log.Fatal("unable to connect to client")
	}
	defer database.Disconnect()
	port := os.Getenv("PORT")
	if port == "" {
		port = "6969"
	}
	log.Fatal(app.Listen(":" + port))

	log.Println("Server started on port 6969")
}
