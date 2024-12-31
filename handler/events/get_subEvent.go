package events

import (
	"context"
	"fmt"
	"technexRegistration/database"
	"technexRegistration/models"
	"technexRegistration/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetSubEvents(c *fiber.Ctx) error {
	// Extract the event ID from query parameters
	id := c.Query("id")
	if id == "" {
		return utils.ResponseMsg(c, 400, "Event ID is required", nil)
	}

	// Convert the ID to an ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return utils.ResponseMsg(c, 400, "Invalid ID", nil)
	}

	// Connect to the database
	var ctx = context.Background()
	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	// Find the event by ID
	var event models.Event
	err = db.Collection("events").FindOne(ctx, bson.M{"_id": objID}).Decode(&event)
	if err != nil {
		return utils.ResponseMsg(c, 404, "Event not found", nil)
	}

	// Return the subEvents of the found event
	fmt.Println("event", event)
	fmt.Println("event.SubEvents", event.SubEvents)
	return c.Status(200).JSON(fiber.Map{"subEvents": event.SubEvents})
}
