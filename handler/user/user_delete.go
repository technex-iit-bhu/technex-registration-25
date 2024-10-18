package user

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"technexRegistration/database"
	"technexRegistration/utils"
)

func DeleteUser(c *fiber.Ctx) error {
	token := c.Get("Authorization")[7:]
	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	username, err := utils.DeserialiseUser(token)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "invalid token"})
	}
	result, _ := db.Collection("users").DeleteOne(context.Background(), bson.D{{Key: "username", Value: username}})
	if result.DeletedCount == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "user does not exist"})
	}
	return c.Status(200).JSON(fiber.Map{"message": "user yeeted successfully"})
}
