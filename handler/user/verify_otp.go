package user

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
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

	identifier := strings.TrimSpace(body.Email)
	if identifier == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Email or Username is required"})
	}

	// Check if user exists by email or username
	var user models.Users
	filter := bson.M{}

	if strings.Contains(identifier, "@") {
		cleanEmail := normalizeEmail(identifier)
		// Use regex for case-insensitive email matching just like send_otp
		pattern := fmt.Sprintf("^%s$", regexp.QuoteMeta(cleanEmail))
		filter = bson.M{"email": bson.M{"$regex": pattern, "$options": "i"}}
	} else {
		// Try to find by username if input is not email
		pattern := fmt.Sprintf("^%s$", regexp.QuoteMeta(identifier))
		filter = bson.M{"username": bson.M{"$regex": pattern, "$options": "i"}}
	}

	err = db.Collection("users").FindOne(context.Background(), filter).Decode(&user)
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
	filter = bson.M{
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
		userEmail := strings.ToLower(user.Email)
		parts := strings.Split(userEmail, "@")
		if len(parts) != 2 {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid email format in stored user",
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

	if body.Purpose == "reset" {
		// Issue a short-lived reset token
		resetToken, err := utils.SerialiseRecovery(user.Username)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Failed to generate reset token",
			})
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message":     "OTP verified successfully",
			"reset_token": resetToken,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "OTP verified successfully",
	})
}
