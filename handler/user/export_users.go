package user

import (
	"context"
	"log"
	"technexRegistration/config"
	"technexRegistration/database"
	"technexRegistration/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func ExportUsers(c *fiber.Ctx) error {
	// Get API key from header
	apiKey := c.Get("api-key")

	// Get admin key from environment variable
	adminKey := config.Config("admin_key")

	// Authenticate with API key
	if apiKey == "" {
		return c.Status(401).JSON(fiber.Map{"message": "api-key header is required"})
	}

	if apiKey != adminKey {
		return c.Status(401).JSON(fiber.Map{"message": "invalid api-key"})
	}

	// Connect to database
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	// Fetch all users from database
	cursor, err := db.Collection("users").Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	defer cursor.Close(context.Background())

	var users []models.Users
	if err = cursor.All(context.Background(), &users); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	if users == nil {
		users = []models.Users{}
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "users exported successfully",
		"count":   len(users),
		"data":    users,
	})
}
