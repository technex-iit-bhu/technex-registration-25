package events

import (
	"context"
	"technexRegistration/database"
	"technexRegistration/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"technexRegistration/models"
)

func GetEventByID(c *fiber.Ctx) error {
	ctx := context.Background()
	
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return utils.ResponseMsg(c, 400, "Invalid ID", nil)
	}

	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	event := new(models.Event)
	
	if err := c.BodyParser(event); err != nil {
		return utils.ResponseMsg(c, 400, "Error parsing body", nil)
	} else {
		err = db.Collection("events").FindOne(ctx, bson.D{{Key: "_id", Value: objID}}).Decode(&event)
		if err != nil {
			return utils.ResponseMsg(c, 404, "Event not found", nil)
		}
	}

	return c.Status(200).JSON(fiber.Map{"event": event})
}
