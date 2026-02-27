package user

import (
	"context"
	"technexRegistration/database"
	"technexRegistration/models"
	"technexRegistration/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type Body struct {
	Name        string `json:"name"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
	Institute   string `json:"institute"`
	City        string `json:"city"`
	Year        int    `json:"year"`
	Branch      string `json:"branch"`
	Phone       string `json:"phone"`
}

func UpdateDetails(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if len(authHeader) < 8 || authHeader[:7] != "Bearer " {
		return c.Status(401).JSON(fiber.Map{"message": "authorization header missing"})
	}
	token := authHeader[7:]

	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	username, err := utils.DeserialiseAccessToken(token)

	var body Body

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request body"})
	}

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "invalid token"})
	}

	var user models.Users
	err = db.Collection("users").FindOne(context.Background(), bson.D{{Key: "username", Value: username}}).Decode(&user)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "user does not exist"})
	}
	if !utils.CheckPassword(body.OldPassword, user.Password) {
		return c.Status(404).JSON(fiber.Map{"message": "invalid password"})
	}

	body.OldPassword = utils.HashPassword(body.OldPassword)
	if body.NewPassword == "" {
		body.NewPassword = body.OldPassword
	} else {
		body.NewPassword = utils.HashPassword(body.NewPassword)
	}

	result, _ := db.Collection("users").UpdateOne(context.Background(), bson.D{{Key: "username", Value: username}},
		bson.D{{Key: "$set", Value: bson.D{
			{Key: "name", Value: body.Name},
			{Key: "password", Value: body.NewPassword},
			{Key: "institute", Value: body.Institute},
			{Key: "city", Value: body.City}, {Key: "year", Value: body.Year},
			{Key: "branch", Value: body.Branch},
			{Key: "phone", Value: body.Phone},
			{Key: "UpdatedAt", Value: time.Now()}}}})

	if result.MatchedCount == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "user does not exist"})
	}

	// Invalidate the cached profile so the next request fetches fresh data
	utils.DeleteUserProfile(username)
	return c.Status(200).JSON(fiber.Map{"message": "user updated successfully"})
}
