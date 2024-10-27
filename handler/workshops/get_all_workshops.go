package workshops

import (
	"context"
	"technexRegistration/database"
	"technexRegistration/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllWorkshops(c *fiber.Ctx) error {
	ctx := context.Background()
	token := c.Get("Authorization")[7:]
	if token == "" {
		return c.Status(401).JSON(fiber.Map{"message": "Unauthorized"})
	}

	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	
	cursor, err := db.Collection("workshops").Find(ctx, bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	defer cursor.Close(ctx)

	var workshops []models.Workshop

	for cursor.Next(ctx) {
		var workshop models.Workshop
		if err := cursor.Decode(&workshop); err != nil {
			return c.Status(500).JSON(fiber.Map{"message": err.Error()})
		}
		workshops = append(workshops, workshop)
	}

	if err := cursor.Err(); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"workshops": workshops})
}
