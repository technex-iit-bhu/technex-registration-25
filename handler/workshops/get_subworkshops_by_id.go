package workshops

import (
	"context"
	"technexRegistration/database"
	"technexRegistration/models"
	"technexRegistration/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetSubWorkshopsByID(c *fiber.Ctx) error {
	ctx := context.Background()

	id := c.Query("id")
	if id == "" {
		return utils.ResponseMsg(c, 400, "ID is required", nil)
	}

	objID, err := primitive.ObjectIDFromHex(id)
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

	return c.Status(200).JSON(fiber.Map{"subworkshops": workshop.SubWorkshops})
}
