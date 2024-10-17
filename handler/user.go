package handler

import (
	"context"
	"fmt"
	"log"
	"technexRegistration/database"
	"technexRegistration/helpers"
	"technexRegistration/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func Hello(c *fiber.Ctx) error {
	return helpers.ResponseMsg(c, 200, "Api is running", nil)
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
		return helpers.ResponseMsg(c, 400, err.Error(), nil)
	} else {
		users.Password = helpers.HashPassword(users.Password)
		if r, err := db.Collection("users").InsertOne(ctx, users); err != nil {
			return c.Status(500).JSON(fiber.Map{"message": err.Error()})
		} else {
			return c.Status(201).JSON(fiber.Map{"id": r.InsertedID})
		}
	}
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
	fmt.Println(body)

	var result models.Users
	err = db.Collection("users").FindOne(context.Background(), bson.D{{"username", body.Username}}).Decode(&result)
	fmt.Println(result)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "invalid username"})
	}
	if !helpers.CheckPassword(body.Password, result.Password) {
		return c.Status(404).JSON(fiber.Map{"message": "invalid password"})
	} else {
		token, _ := helpers.SerialiseUser(result.Username)
		return c.Status(200).JSON(fiber.Map{"token": token})
	}
}

func GetUserFromToken(c *fiber.Ctx) error {
	var body struct {
		Token string `json:"token"`
	}
	c.BodyParser(&body)
	db, err := database.Connect()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": err.Error()})
	}
	username,err:=helpers.DeserialiseUser(body.Token)
	fmt.Println(username)
	if err!=nil{
		return c.Status(404).JSON(fiber.Map{"message": "invalid token"})
	}
	// objectId, _ := primitive.ObjectIDFromHex(userID)
	var result models.Users
	err = db.Collection("users").FindOne(context.Background(), bson.D{{"username", username}}).Decode(&result)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"data":result})
}
