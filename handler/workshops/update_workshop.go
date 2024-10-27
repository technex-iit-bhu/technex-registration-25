package workshops

import (
	"context"
	"technexRegistration/database"
	"technexRegistration/models"
	"technexRegistration/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
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

	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	updatedWorkshop := bson.M{}

	if workshop.Name != "" {
		updatedWorkshop["name"] = workshop.Name
	} else if workshop.Description != "" {
		updatedWorkshop["description"] = workshop.Description
	} else if workshop.Start_Date.IsZero() {
		updatedWorkshop["startDate"] = workshop.Start_Date
	} else if workshop.End_Date.IsZero() {
		updatedWorkshop["endDate"] = workshop.End_Date
	} else if workshop.SubDescription != "" {
		updatedWorkshop["sub_description"] = workshop.SubDescription
	} else if workshop.Github != "" {
		updatedWorkshop["github"] = workshop.Github
	}

	if len(updatedWorkshop) == 0 {
		return utils.ResponseMsg(c, 400, "No fields to update", nil)
	}

	if err := c.BodyParser(workshop); err != nil {
		return utils.ResponseMsg(c, 400, err.Error(), nil)
	} else {
		if err := db.Collection("workshops").FindOneAndUpdate(ctx, workshop, updatedWorkshop); err != nil {
			return utils.ResponseMsg(c, 500, "Failed to Update", nil)
		} else {
			return utils.ResponseMsg(c, 200, "Workshop updated successfully", nil)
		}
	}
}
