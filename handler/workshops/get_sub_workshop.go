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

func GetSubWorkshops(c *fiber.Ctx) error {
	// Extract the workshop ID from query parameters
	id := c.Query("id")
	if id == "" {
		return utils.ResponseMsg(c, 400, "Workshop ID is required", nil)
	}

	// Convert the ID to an ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return utils.ResponseMsg(c, 400, "Invalid ID", nil)
	}

	// Connect to the database
	var ctx = context.Background()
	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	// Find the workshop by ID
	var workshop models.Workshop
	err = db.Collection("workshops").FindOne(ctx, bson.M{"_id": objID}).Decode(&workshop)
	if err != nil {
		return utils.ResponseMsg(c, 404, "Workshop not found", nil)
	}

	// Return the subWorkshops of the found workshop
	return c.Status(200).JSON(fiber.Map{"subWorkshops": workshop.SubWorkshops})
}
