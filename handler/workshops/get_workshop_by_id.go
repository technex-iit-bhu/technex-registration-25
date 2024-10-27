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

func GetWorkshopByID(c *fiber.Ctx) error {
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

	workshop := new(models.Workshop)

	if err := c.BodyParser(workshop); err != nil {
		return utils.ResponseMsg(c, 400, "Error parsing body", nil)
	} else {
		err = db.Collection("workshops").FindOne(ctx, bson.D{{Key: "_id", Value: objID}}).Decode(&workshop)
		if err != nil {
			return utils.ResponseMsg(c, 404, "Event not found", nil)
		}
	}

	return c.Status(200).JSON(fiber.Map{"workshop": workshop})
}
