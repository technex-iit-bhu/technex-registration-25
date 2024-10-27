package workshops

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"technexRegistration/database"
	"technexRegistration/models"
	"technexRegistration/utils"
)

func DeleteWorkshop(c *fiber.Ctx) error {
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

	workshop := new(models.Workshop)
	if err := c.BodyParser(workshop); err != nil {
		return utils.ResponseMsg(c, 400, "Error parsing body", nil)
	}

	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	if err := c.BodyParser(workshop); err != nil {
		return utils.ResponseMsg(c, 400, "Error parsing body", nil)
	} else {
		err = db.Collection("workshops").FindOneAndDelete(ctx, bson.D{{Key: "_id", Value: objID}}).Decode(&workshop)
		if err != nil {
			return utils.ResponseMsg(c, 404, "Workshop not found", nil)
		}
	}

	return c.Status(200).JSON(fiber.Map{"deleted workshop": workshop})
}
