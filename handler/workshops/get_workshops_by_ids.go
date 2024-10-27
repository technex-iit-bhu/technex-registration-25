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

func GetWorkshopsByID(c *fiber.Ctx) error {
	ctx := context.Background()

	var requestData struct {
		IDs []string `json:"ids"`
	}

	if err := c.BodyParser(&requestData); err != nil {
		return utils.ResponseMsg(c, 400, "Error parsing body", nil)
	}

	var objIDs []primitive.ObjectID
	for _, id := range requestData.IDs {
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return utils.ResponseMsg(c, 400, "Invalid ID format", nil)
		}
		objIDs = append(objIDs, objID)
	}

	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	cursor, err := db.Collection("workshops").Find(ctx, bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: objIDs}}}})
	if err != nil {
		return utils.ResponseMsg(c, 500, "Error fetching workshops", nil)
	}
	defer cursor.Close(ctx)

	var workshops []models.Workshop
	if err := cursor.All(ctx, &workshops); err != nil {
		return utils.ResponseMsg(c, 500, "Error decoding workshops", nil)
	}

	if len(workshops) == 0 {
		return utils.ResponseMsg(c, 404, "No workshops found", nil)
	}

	return c.Status(200).JSON(fiber.Map{"workshops": workshops})
}
