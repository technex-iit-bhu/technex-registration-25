package user

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"regexp"
	"strings"
	"technexRegistration/database"
	"technexRegistration/models"
	"technexRegistration/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SendOTP(c *fiber.Ctx) error {
	// accept either email or username in the body
	var body struct {
		EmailOrUsername string `json:"email"`
		Purpose         string `json:"purpose"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	ctx := context.Background()

	// Validate purpose
	if body.Purpose != "reset" && body.Purpose != "verify" {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid purpose"})
	}

	// Normalize the identifier (email or username)
	identifier := strings.TrimSpace(body.EmailOrUsername)
	if identifier == "" {
		return c.Status(400).JSON(fiber.Map{"message": "Email or username is required"})
	}

	// Check if user exists by email or username
	var user models.Users
	filter := bson.M{}
	if strings.Contains(identifier, "@") {
		// Use regex for case-insensitive email matching to support legacy data
		cleanEmail := normalizeEmail(identifier)
		pattern := fmt.Sprintf("^%s$", regexp.QuoteMeta(cleanEmail))
		filter = bson.M{"email": bson.M{"$regex": pattern, "$options": "i"}}
	} else {
		filter = usernameCaseInsensitiveFilter(identifier)
		if len(filter) == 0 {
			return c.Status(400).JSON(fiber.Map{"message": "Invalid username"})
		}
	}
	err = db.Collection("users").FindOne(ctx, filter).Decode(&user)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "User not found"})
	}

	if user.Email == "" {
		fmt.Printf("Data Error: User %s found but has no email\n", user.Username)
		return c.Status(500).JSON(fiber.Map{"message": "User record is invalid (missing email)"})
	}

	// Generate OTP using crypto/rand
	n, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to generate OTP"})
	}
	otpCode := fmt.Sprintf("%06d", n.Int64())

	// Delete previous OTPs using correct field name
	db.Collection("otps").DeleteMany(ctx, bson.M{
		"userId":  user.ID,
		"purpose": body.Purpose,
		"used":    false,
	})

	// Send OTP via Resend
	if body.Purpose == "reset" {
		err = utils.RecoveryMail(user.Email, user.Username, otpCode)
	}

	if body.Purpose == "verify" {
		err = utils.VerificationMail(user.Email, user.Username, otpCode)
	}

	if err != nil {
		fmt.Println("RESEND ERROR:", err)
		return c.Status(500).JSON(fiber.Map{"message": "Failed to send OTP"})
	}

	otpDoc := models.Otp{
		ID:        primitive.NewObjectID(),
		UserID:    user.ID,
		Code:      int(n.Int64()),
		Purpose:   body.Purpose,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(10 * time.Minute),
		Used:      false,
	}

	_, err = db.Collection("otps").InsertOne(ctx, otpDoc)
	if err != nil {
		fmt.Println("OTP DB ERROR:", err)
		return c.Status(500).JSON(fiber.Map{"message": "Failed to save OTP"})
	}

	return c.JSON(fiber.Map{
		"message":  "OTP sent successfully",
		"email":    user.Email,
		"username": user.Username,
	})
}
