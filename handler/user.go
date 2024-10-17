package handler

import (
	"context"
	_ "fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"technexRegistration/database"
	"technexRegistration/models"
	"technexRegistration/utils"
	"time"
)

func Hello(c *fiber.Ctx) error {
	return utils.ResponseMsg(c, 200, "Api is running", nil)
}

func CreateUsers(c *fiber.Ctx) error {
	users := new(models.Users)
	var ctx = context.Background()
	db, err := database.Connect()

	if err != nil {
		log.Fatal(err.Error())
		return c.Status(400).JSON(fiber.Map{"message": err.Error()})
	}

	users.CreatedAt = time.Now()
	users.UpdatedAt = time.Now()

	if err := c.BodyParser(users); err != nil {
		return utils.ResponseMsg(c, 400, err.Error(), nil)
	} else {
		users.Password = utils.HashPassword(users.Password)
		if r, err := db.Collection("users").InsertOne(ctx, users); err != nil {
			return c.Status(500).JSON(fiber.Map{"message": err.Error()})
		} else {
			return c.Status(201).JSON(fiber.Map{"id": r.InsertedID})
		}
	}
}

func GetUserFromToken(c *fiber.Ctx) error {
	token := c.Get("Authorization")[7:]
	db, err := database.Connect()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": err.Error()})
	}
	username, err := utils.DeserialiseUser(token)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "invalid token"})
	}
	var result models.Users
	err = db.Collection("users").FindOne(context.Background(), bson.D{{Key: "username", Value: username}}).Decode(&result)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "user does not exist"})
	}
	return c.Status(200).JSON(fiber.Map{"data": result})
}

func LoginWithPassword(c *fiber.Ctx) error {
	db, err := database.Connect()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": err.Error()})
	}

	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	c.BodyParser(&body)

	var result models.Users
	err = db.Collection("users").FindOne(context.Background(), bson.D{{Key: "username", Value: body.Username}}).Decode(&result)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "invalid username"})
	}
	if !utils.CheckPassword(body.Password, result.Password) {
		return c.Status(404).JSON(fiber.Map{"message": "invalid password"})
	} else {
		token, _ := utils.SerialiseUser(result.Username)
		return c.Status(200).JSON(fiber.Map{"token": token})
	}
}

func LoginWithGoogle(c *fiber.Ctx) error {
	db, err := database.Connect()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": err.Error()})
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
		return c.Status(400).JSON(fiber.Map{"message": err.Error()})
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

func DeleteUser(c *fiber.Ctx) error {
	token := c.Get("Authorization")[7:]
	db, err := database.Connect()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": err.Error()})
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
func UpdateDetails(c *fiber.Ctx) error {
	token := c.Get("Authorization")[7:]
	db, err := database.Connect()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": err.Error()})
	}
	username, err := utils.DeserialiseUser(token)

	var body struct {
		Name        string `json:"name"`
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
		Institute   string `json:"institute"`
		City        string `json:"city"`
		Year        int    `json:"year"`
		Branch      string `json:"branch"`
		Phone       string `json:"phone"`
	}
	c.BodyParser(&body)
	body.OldPassword = utils.HashPassword(body.OldPassword)
	if body.NewPassword == "" {
		body.NewPassword = body.OldPassword
	} else {
		body.NewPassword = utils.HashPassword(body.NewPassword)
	}
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "invalid token"})
	}
	result, _ := db.Collection("users").UpdateOne(context.Background(), bson.D{{Key: "username", Value: username}, {Key: "password", Value: body.OldPassword}},
		bson.D{{Key: "$set", Value: bson.D{
			{Key: "name", Value: body.Name},
			{Key: "password", Value: body.NewPassword},
			{Key: "institute", Value: body.Institute},
			{Key: "city", Value: body.City}, {Key: "year", Value: body.Year},
			{Key: "branch", Value: body.Branch},
			{Key: "phone", Value: body.Phone}}}})
	if result.MatchedCount == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "user does not exist"})
	}
	return c.Status(200).JSON(fiber.Map{"message": "user updated successfully"})
}
