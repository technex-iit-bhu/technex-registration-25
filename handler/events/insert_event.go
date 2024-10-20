package events

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"technexRegistration/database"
	"technexRegistration/models"
	"technexRegistration/utils"
)

func InsertEvent(c *fiber.Ctx) error {
	event := new(models.Event)
	var ctx = context.Background()
	token := c.Get("Authorization")[7:]
	if token == "" {
		return c.Status(401).JSON(fiber.Map{"message": "Unauthorized"})
	}

	db, err := database.Connect()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	if err := c.BodyParser(event); err != nil {
		return utils.ResponseMsg(c, 400, err.Error(), nil)
	} else {
		if r, err := db.Collection("events").InsertOne(ctx, event); err != nil {
			return c.Status(500).JSON(fiber.Map{"message": err.Error()})
		} else {
			return c.Status(201).JSON(fiber.Map{"id": r.InsertedID})
		}
	}
}
