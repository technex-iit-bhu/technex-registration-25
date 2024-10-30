package events

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"technexRegistration/database"
	"technexRegistration/models"
	"technexRegistration/utils"
)

func DeleteEvent(c *fiber.Ctx) error {
	ctx := context.Background()
	token := c.Get("Authorization")[7:]
	if token == "" {
		return c.Status(401).JSON(fiber.Map{"message": "Unauthorized"})
	}

	event := new(models.Event)
	if err := c.BodyParser(event); err != nil {
		return utils.ResponseMsg(c, 400, "Error parsing body", nil)
	}

	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	if r, err := db.Collection("events").DeleteOne(ctx, bson.D{{Key: "name", Value: event.Name}}); err != nil {
		return utils.ResponseMsg(c, 400, "Delete failed", nil)
	} else {
		return c.Status(200).JSON(fiber.Map{
			"message": "Event deleted",
			"deleted": r.DeletedCount,
			"ID":      r.DeletedCount,
		})
	}
}
