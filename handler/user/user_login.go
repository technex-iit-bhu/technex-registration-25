package user

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"technexRegistration/database"
	"technexRegistration/models"
	"technexRegistration/utils"
	"strings"
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
	c.BodyParser(&body)

	if body.Username == "" || body.Password == "" {
		return c.Status(400).JSON(fiber.Map{"message": "username/email and password are required"})
	}

	var result models.Users
	var filter bson.D

	if strings.Contains(body.Username, "@") {
		filter = bson.D{{Key: "email", Value: body.Username}}
	} else {
		filter = bson.D{{Key: "username", Value: body.Username}}
	}

	err = db.Collection("users").FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "invalid username or email"})
	}

	if !utils.CheckPassword(body.Password, result.Password) {
		return c.Status(404).JSON(fiber.Map{"message": "invalid password"})
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
	} else {
		token, _ := utils.SerialiseUser(result.Username)
		return c.Status(200).JSON(fiber.Map{"token": token})
	}
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

	email, _ := utils.DeserialiseGmailToken(body.GithubToken)
	var result models.Users
	err = db.Collection("users").FindOne(context.Background(), bson.D{{Key: "github", Value: email}}).Decode(&result)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "invalid github"})
	} else {
		token, _ := utils.SerialiseUser(result.Username)
		return c.Status(200).JSON(fiber.Map{"token": token})
	}
}
