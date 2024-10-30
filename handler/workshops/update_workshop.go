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

	workshop := new(models.Workshop)
	if err := c.BodyParser(workshop); err != nil {
		return utils.ResponseMsg(c, 400, "Error parsing body", nil)
	}

	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return utils.ResponseMsg(c, 400, "Invalid ID", nil)
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
	if workshop.Start_Date.IsZero() {
		updatedWorkshop = append(updatedWorkshop, bson.E{Key: "startDate", Value: workshop.Start_Date})
	}
	if workshop.End_Date.IsZero() {
		updatedWorkshop = append(updatedWorkshop, bson.E{Key: "endDate", Value: workshop.End_Date})
	}
	if workshop.SubDescription != "" {
		updatedWorkshop = append(updatedWorkshop, bson.E{Key: "sub_description", Value: workshop.SubDescription})
	}
	if workshop.Github != "" {
		updatedWorkshop = append(updatedWorkshop, bson.E{Key: "github", Value: workshop.Github})
	}

	if len(updatedWorkshop) == 0 {
		return utils.ResponseMsg(c, 400, "No fields to update", nil)
	}

	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.D{{Key: "$set", Value: updatedWorkshop}}

	result := db.Collection("workshops").FindOneAndUpdate(ctx, filter, update)
	if result.Err() != nil {
		return utils.ResponseMsg(c, 500, "Failed to update", nil)
	}

	return utils.ResponseMsg(c, 200, "Workshop updated successfully", nil)
}
