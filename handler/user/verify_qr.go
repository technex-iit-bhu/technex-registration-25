package user

import (
	"context"
	"technexRegistration/database"
	"technexRegistration/models"
	"technexRegistration/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func VerifyQR(c *fiber.Ctx) error {
	var body struct {
		QRToken   string `json:"qr_token" bson:"qr_token"`
		EventName string `json:"event_name" bson:"event_name"`
	}

	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	username, err := utils.DeserialiseQR(body.QRToken)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid QR token",
		})
	}

	var result models.Users
	err = db.Collection("users").FindOne(context.Background(), bson.D{{Key: "username", Value: username}}).Decode(&result)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "user does not exist"})
	}

	isRegistered := false
	for _, event := range result.RegisteredEvents {
		if event == body.EventName {
			isRegistered = true
			break
		}
	}

	if !isRegistered {
		return c.Status(401).JSON(fiber.Map{
			"message": "User not registered for this event",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":  "User successfully verified for event",
		"id":       result.ID,
		"name":     result.Name,
		"username": result.Username,
	})
}
