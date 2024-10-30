package events

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"technexRegistration/database"
	"technexRegistration/models"
)

func GetAllEvents(c *fiber.Ctx) error {
	var ctx = context.Background()

	token := c.Get("Authorization")[7:]
	if token == "" {
		return c.Status(401).JSON(fiber.Map{"message": "Unauthorized"})
	}

	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	cursor, err := db.Collection("events").Find(ctx, bson.D{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	defer cursor.Close(ctx)

	var events []models.Event

	for cursor.Next(ctx) {
		var event models.Event
		if err := cursor.Decode(&event); err != nil {
			return c.Status(500).JSON(fiber.Map{"message": err.Error()})
		}
		events = append(events, event)
	}

	if err := cursor.Err(); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"events": events})
}
