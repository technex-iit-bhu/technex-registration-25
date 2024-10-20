package events

import (
	"context"
	"technexRegistration/database"
	"technexRegistration/models"
	"technexRegistration/utils"
	"github.com/gofiber/fiber/v2"
)

func BulkInsertEvents(c *fiber.Ctx) error {
	ctx := context.Background()

	var events []models.Event
	if err := c.BodyParser(&events); err != nil {
		return utils.ResponseMsg(c, 400, "Error parsing body", nil)
	}

	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	var docs []interface{}
	for _, event := range events {
		docs = append(docs, event)
	}

	res, err := db.Collection("events").InsertMany(ctx, docs)
	if err != nil {
		return utils.ResponseMsg(c, 400, "Failed to insert events", nil)
	}

	return c.Status(200).JSON(fiber.Map{"message": "Events inserted", "insertedIDs": res.InsertedIDs})
}
