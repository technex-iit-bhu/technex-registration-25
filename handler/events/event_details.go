package events

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"technexRegistration/database"
	"technexRegistration/models"
	"technexRegistration/utils"
)

func GetEventDetails(c *fiber.Ctx) error {
	event := new(models.Event)
	var ctx = context.Background()

	db, err := database.Connect()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	if err := c.BodyParser(event); err != nil {
		return utils.ResponseMsg(c, 400, err.Error(), nil)
	}

	filter := bson.D{{Key: "name", Value: event.Name}}
	var foundEvent models.Event
	err = db.Collection("events").FindOne(ctx, filter).Decode(&foundEvent)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(200).JSON(foundEvent)
}
