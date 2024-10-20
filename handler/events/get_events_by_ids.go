package events

import (
	"context"
	"technexRegistration/database"
	"technexRegistration/models"
	"technexRegistration/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetEventsByID(c *fiber.Ctx) error {
	ctx := context.Background()

	var requestData struct {
		IDs []string `json:"ids"`
	}

	if err := c.BodyParser(&requestData); err != nil {
		return utils.ResponseMsg(c, 400, "Error parsing body", nil)
	}

	var objIDs []primitive.ObjectID
	for _, id := range requestData.IDs {
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return utils.ResponseMsg(c, 400, "Invalid ID format", nil)
		}
		objIDs = append(objIDs, objID)
	}

	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	cursor, err := db.Collection("events").Find(ctx, bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: objIDs}}}})
	if err != nil {
		return utils.ResponseMsg(c, 500, "Error fetching events", nil)
	}
	defer cursor.Close(ctx)

	var events []models.Event
	if err := cursor.All(ctx, &events); err != nil {
		return utils.ResponseMsg(c, 500, "Error decoding events", nil)
	}

	if len(events) == 0 {
		return utils.ResponseMsg(c, 404, "No events found", nil)
	}

	return c.Status(200).JSON(fiber.Map{"events": events})
}
