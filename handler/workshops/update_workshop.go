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

func UpdateWorkshop(c *fiber.Ctx) error {
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

	updatedWorkshop := bson.D{}
	if workshop.Name != "" {
		updatedWorkshop = append(updatedWorkshop, bson.E{Key: "name", Value: workshop.Name})
	}
	if workshop.Description != "" {
		updatedWorkshop = append(updatedWorkshop, bson.E{Key: "description", Value: workshop.Description})
	}

	if len(updatedWorkshop) == 0 {
		return utils.ResponseMsg(c, 400, "No fields to update", nil)
	}

	if err := c.BodyParser(workshop); err != nil {
		return utils.ResponseMsg(c, 400, "Error parsing body", nil)
	} else {
		if _, err := db.Collection("workshops").UpdateByID(ctx, objID, bson.D{{Key: "$set", Value: updatedWorkshop}}); err != nil {
			return utils.ResponseMsg(c, 400, "Update failed", nil)
		} else {
			return c.Status(200).JSON(fiber.Map{"message": "workshop updated successfully"})
		}
	}
}
