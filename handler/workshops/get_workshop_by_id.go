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

	reqBody := new(struct {
		ID string `json:"id"`
	})
	if err := c.BodyParser(reqBody); err != nil {
		return utils.ResponseMsg(c, 400, "Error parsing body", nil)
	}

	objID, err := primitive.ObjectIDFromHex(reqBody.ID)
	if err != nil {
		return utils.ResponseMsg(c, 400, "Invalid ID", nil)
	}

	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	workshop := new(models.Workshop)
	err = db.Collection("workshops").FindOne(ctx, bson.D{{Key: "_id", Value: objID}}).Decode(&workshop)
	if err != nil {
		return utils.ResponseMsg(c, 404, "Workshop not found", nil)
	}

	return c.Status(200).JSON(fiber.Map{"workshop": workshop})
}
