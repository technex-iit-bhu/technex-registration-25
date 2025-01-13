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

func UpdateSubEvents(c *fiber.Ctx) error {
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

	var subEvents []models.SubEvent
	if err := c.BodyParser(&subEvents); err != nil {
		return utils.ResponseMsg(c, 400, "Error parsing body", nil)
	}

	// Validate required fields for each subevent
	for _, subEvent := range subEvents {
		if subEvent.Name == "" || subEvent.Description == "" || subEvent.SubDescription == "" ||
			subEvent.Start_Date.IsZero() || subEvent.End_Date.IsZero() || subEvent.Github == "" ||
			subEvent.DriveLink == "" || subEvent.UnstopLink == "" || subEvent.PrizeMoney == 0 {
			return utils.ResponseMsg(c, 400, "Missing required fields in subevent", nil)
		}
	}

	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	// Find the event
	var event models.Event
	err = db.Collection("events").FindOne(ctx, bson.M{"_id": objID}).Decode(&event)
	if err != nil {
		return utils.ResponseMsg(c, 404, "Event not found", nil)
	}

	// Check if any subevent already exists
	for _, newSubEvent := range subEvents {
		for _, existingSubEvent := range event.SubEvents {
			if existingSubEvent.Name == newSubEvent.Name {
				return utils.ResponseMsg(c, 400, "SubEvent "+newSubEvent.Name+" already exists", nil)
			}
		}
	}

	// Append new subevents
	update := bson.M{
		"$push": bson.M{
			"subEvents": bson.M{
				"$each": subEvents,
			},
		},
	}

	_, err = db.Collection("events").UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return utils.ResponseMsg(c, 500, "Failed to update event", nil)
	}

	return c.Status(200).JSON(fiber.Map{"message": "SubEvents added successfully"})
}
