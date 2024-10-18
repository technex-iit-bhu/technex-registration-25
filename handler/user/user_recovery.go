package user

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"technexRegistration/database"
	"technexRegistration/models"
	"technexRegistration/utils"
	"time"
)

func SendRecoveryEmail(c *fiber.Ctx) error {
	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	username := c.Params("username")
	var result models.Users
	err = db.Collection("users").FindOne(context.Background(), bson.D{{Key: "username", Value: username}}).Decode(&result)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "user does not exist"})
	}
	err = utils.RecoveryMail(result.Email, utils.GenerateOTPConnectionString(username))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "unable to send recovery email"})
	}
	return c.Status(200).JSON(fiber.Map{"message": "recovery email sent successfully"})
}

func UpdatePassword(c *fiber.Ctx) error {
	var body struct {
		RecoveryToken string `json:"recovery_token"`
		NewPassword   string `json:"new_password"`
	}
	c.BodyParser(&body)
	username, err := utils.DeserialiseRecovery(body.RecoveryToken)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid recovery token"})
	}
	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	var result models.Users
	err = db.Collection("users").FindOne(context.Background(), bson.D{{Key: "username", Value: username}}).Decode(&result)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "user does not exist"})
	}
	res, _ := db.Collection("users").UpdateOne(context.Background(), bson.D{{Key: "username", Value: username}}, bson.D{{Key: "$set", Value: bson.D{{Key: "password", Value: utils.HashPassword(body.NewPassword)}, {Key: "UpdatedAt", Value: time.Now()}}}})
	if res.MatchedCount == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "user does not exist"})
	} else {
		return c.Status(200).JSON(fiber.Map{"message": "password updated successfully"})
	}
}
