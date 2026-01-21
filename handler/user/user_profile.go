package user

import (
	"context"
	"technexRegistration/database"
	"technexRegistration/models"
	"technexRegistration/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUserFromToken(c *fiber.Ctx) error {
	token := c.Get("Authorization")[7:]
	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	username, err := utils.DeserialiseUser(token)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "invalid token"})
	}
	// Check cache first
	if cached, ok := utils.GetUserProfile(username); ok {
		qrToken, _ := utils.SerialiseQR(cached.Username)
		return c.Status(200).JSON(fiber.Map{"data": cached, "qrToken": qrToken})
	}

	var result models.Users
	err = db.Collection("users").FindOne(context.Background(), bson.D{{Key: "username", Value: username}}).Decode(&result)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "user does not exist"})
	}

	// Store in cache
	utils.SetUserProfile(username, result)
	qrToken, _ := utils.SerialiseQR(result.Username)
	return c.Status(200).JSON(fiber.Map{
		"data":    result,
		"qrToken": qrToken,
	})
}
