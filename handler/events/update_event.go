package events

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"technexRegistration/database"
	"technexRegistration/models"
	"technexRegistration/utils"
)

func UpdateEvent(c *fiber.Ctx) error {
	ctx := context.Background()
	token := c.Get("Authorization")[7:]
	if token == "" {
		return c.Status(401).JSON(fiber.Map{"message": "Unauthorized"})
	}

	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return utils.ResponseMsg(c, 400, "Invalid ID", nil)
	}

	event := new(models.Event)
	if err := c.BodyParser(event); err != nil {
		return utils.ResponseMsg(c, 400, "Error parsing body", nil)
	}

	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	updatedEvent := bson.D{}
	if event.Name != "" {
		updatedEvent = append(updatedEvent, bson.E{Key: "name", Value: event.Name})
	}
	if event.Description != "" {
		updatedEvent = append(updatedEvent, bson.E{Key: "description", Value: event.Description})
	}

	if len(updatedEvent) == 0 {
		return utils.ResponseMsg(c, 400, "No fields to update", nil)
	}

	if err := c.BodyParser(event); err != nil {
		return utils.ResponseMsg(c, 400, "Error parsing body", nil)
	} else {
		if _, err := db.Collection("events").UpdateByID(ctx, objID, bson.D{{Key: "$set", Value: updatedEvent}}); err != nil {
			return utils.ResponseMsg(c, 400, "Update failed", nil)
		} else {
			return c.Status(200).JSON(fiber.Map{"message": "user updated successfully"})
		}
	}
}
