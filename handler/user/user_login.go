package user

import (
	"context"
	"strings"
	"technexRegistration/database"
	"technexRegistration/models"
	"technexRegistration/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func LoginWithPassword(c *fiber.Ctx) error {
	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid request body"})
	}
	body.Username = strings.TrimSpace(body.Username)
	body.Password = strings.TrimSpace(body.Password)

	if body.Username == "" || body.Password == "" {
		return c.Status(400).JSON(fiber.Map{"message": "username/email and password are required"})
	}

	var result models.Users
	var filter bson.M
	if strings.Contains(body.Username, "@") {
		filter = bson.M{"email": normalizeEmail(body.Username)}
	} else {
		filter = usernameCaseInsensitiveFilter(body.Username)
	}
	if len(filter) == 0 {
		return c.Status(400).JSON(fiber.Map{"message": "username/email is invalid"})
	}

	err = db.Collection("users").FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "invalid username or email"})
	}

	if !utils.CheckPassword(body.Password, result.Password) {
		return c.Status(404).JSON(fiber.Map{"message": "invalid password"})
	}

	// Check if email is verified
	if !result.EmailVerified {
		return c.Status(403).JSON(fiber.Map{"message": "Please verify your email before logging in"})
	}

	accessToken, err := utils.SerialiseAccessToken(result.Username)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "failed to generate access token"})
	}

	refreshToken, err := utils.SerialiseRefreshToken(result.Username)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "failed to generate refresh token"})
	}

	c.Cookie(&fiber.Cookie{
		Name:   "refresh_token",
		Value:  refreshToken,
		Path:   "/",
		MaxAge: 7 * 24 * 60 * 60, // 7 days
		// MaxAge:   120, //testing
		HTTPOnly: true,
		Secure:   true, // Set to true in production with HTTPS
		SameSite: "Lax",
	})

	return c.Status(200).JSON(fiber.Map{
		"access_token": accessToken,
		"token_type":   "Bearer",
		"expires_in":   7200, //2 hours
		// "expires_in":   30, //testing
	})
}

func LoginWithGoogle(c *fiber.Ctx) error {
	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	var body struct {
		GoogleToken string `json:"google_token"`
	}
	c.BodyParser(&body)

	email, _ := utils.DeserialiseGmailToken(body.GoogleToken)
	var result models.Users
	err = db.Collection("users").FindOne(context.Background(), bson.D{{Key: "email", Value: email}}).Decode(&result)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "invalid email"})
	}

	// Check if email is verified
	if !result.EmailVerified {
		return c.Status(403).JSON(fiber.Map{"message": "Please verify your email before logging in"})
	}

	accessToken, err := utils.SerialiseAccessToken(result.Username)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "failed to generate access token"})
	}

	refreshToken, err := utils.SerialiseRefreshToken(result.Username)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "failed to generate refresh token"})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/",
		MaxAge:   7 * 24 * 60 * 60, // 7 days
		HTTPOnly: true,
		Secure:   true, // Set to true in production with HTTPS
		SameSite: "Lax",
	})

	return c.Status(200).JSON(fiber.Map{
		"access_token": accessToken,
		"token_type":   "Bearer",
		"expires_in":   7200, //2 hours
	})
}
func LoginWithGithub(c *fiber.Ctx) error {
	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	var body struct {
		GithubToken string `json:"github_token"`
	}
	c.BodyParser(&body)

	email, err := utils.DeserialiseGithubToken(body.GithubToken)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid github token"})
	}
	var result models.Users
	err = db.Collection("users").FindOne(context.Background(), bson.D{{Key: "github", Value: email}}).Decode(&result)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "invalid github"})
	}

	// Check if email is verified
	if !result.EmailVerified {
		return c.Status(403).JSON(fiber.Map{"message": "Please verify your email before logging in"})
	}

	accessToken, err := utils.SerialiseAccessToken(result.Username)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "failed to generate access token"})
	}

	refreshToken, err := utils.SerialiseRefreshToken(result.Username)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "failed to generate refresh token"})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/",
		MaxAge:   7 * 24 * 60 * 60, // 7 days
		HTTPOnly: true,
		Secure:   true, // Set to true in production with HTTPS
		SameSite: "Lax",
	})

	return c.Status(200).JSON(fiber.Map{
		"access_token": accessToken,
		"token_type":   "Bearer",
		"expires_in":   7200, //2 hours
	})
}
