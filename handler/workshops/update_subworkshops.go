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

func UpdateSubWorkshops(c *fiber.Ctx) error {
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

	var subWorkshops []models.SubWorkshop
	if err := c.BodyParser(&subWorkshops); err != nil {
		return utils.ResponseMsg(c, 400, "Error parsing body", nil)
	}

	// Validate required fields for each subworkshop
	for _, subWorkshop := range subWorkshops {
		if subWorkshop.Name == "" || subWorkshop.Description == "" || subWorkshop.SubDescription == "" || 
		   subWorkshop.Start_Date.IsZero() || subWorkshop.End_Date.IsZero() || subWorkshop.Github == "" {
			return utils.ResponseMsg(c, 400, "Missing required fields in subworkshop", nil)
		}
	}

	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	// Find the workshop
	var workshop models.Workshop
	err = db.Collection("workshops").FindOne(ctx, bson.M{"_id": objID}).Decode(&workshop)
	if err != nil {
		return utils.ResponseMsg(c, 404, "Workshop not found", nil)
	}

	// Check if any subworkshop already exists
	for _, newSubWorkshop := range subWorkshops {
		for _, existingSubWorkshop := range workshop.SubWorkshops {
			if existingSubWorkshop.Name == newSubWorkshop.Name {
				return utils.ResponseMsg(c, 400, "SubWorkshop "+newSubWorkshop.Name+" already exists", nil)
			}
		}
	}

	// Append new subworkshops
	update := bson.M{
		"$push": bson.M{
			"subWorkshops": bson.M{
				"$each": subWorkshops,
			},
		},
	}

	_, err = db.Collection("workshops").UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return utils.ResponseMsg(c, 500, "Failed to update workshop", nil)
	}

	return c.Status(200).JSON(fiber.Map{"message": "SubWorkshops added successfully"})
}
