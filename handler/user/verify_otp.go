package user

import (
	"context"
	"net/http"
	"strings"
	"technexRegistration/database"
	"technexRegistration/models"
	"technexRegistration/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func VerifyOTP(c *fiber.Ctx) error {
	var body struct {
		Email   string `json:"email"`
		OTP     int    `json:"otp"`
		Purpose string `json:"purpose"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	db, err := database.Connect()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	// Check if user exists
	var user models.Users
	err = db.Collection("users").FindOne(context.Background(), bson.D{
		{Key: "email", Value: body.Email},
	}).Decode(&user)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "user does not exist"})
	}
	// user := c.Locals("user").(models.Users)

	// if user.Email != body.Email {
	//     return c.Status(403).JSON(fiber.Map{
	//         "message": "Email does not match logged in user",
	//     })
	// }

	// Now lookup the OTP in the "otps" collection
	var storedOtp models.Otp
	filter := bson.M{
		"userId":  user.ID,
		"purpose": body.Purpose,
		"code":    body.OTP,
		"used":    false,
	}
	err = db.Collection("otps").FindOne(context.Background(), filter).Decode(&storedOtp)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid or used OTP"})
	}

	if time.Now().After(storedOtp.ExpiresAt) {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "OTP has expired"})
	}

	_, err = db.Collection("otps").UpdateOne(
		context.Background(),
		bson.M{"_id": storedOtp.ID},
		bson.M{"$set": bson.M{"used": true}},
	)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error marking OTP as used",
		})
	}

	if body.Purpose == "verify" {
		parts := strings.Split(body.Email, "@")
		if len(parts) != 2 {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid email format",
			})
		}

		domain := parts[1]

		isInstitute := domain == "itbhu.ac.in" || domain == "iitbhu.ac.in"

		_, err = db.Collection("users").UpdateOne(
			context.Background(),
			bson.M{"_id": user.ID},
			bson.M{"$set": bson.M{
				"email_verified":    true,
				"email_verified_at": time.Now(),
				"is_institute":      isInstitute,
			}},
		)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Failed to update user verification",
			})
		}

		// Clear the cached user profile after update
		utils.DeleteUserProfile(user.Username)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "OTP verified successfully",
	})
}
