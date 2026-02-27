package user

import (
	"context"
	"fmt"
	"log"
	"technexRegistration/database"
	"technexRegistration/models"
	"technexRegistration/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateUsers(c *fiber.Ctx) error {
	users := new(models.Users)
	var ctx = context.Background()
	db, err := database.Connect()

	if err != nil {
		log.Printf("Database connection error: %v", err)
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	users.CreatedAt = time.Now()
	users.UpdatedAt = time.Now()

	numCollection := db.Collection("num")
	var currentNum struct {
		Number int `bson:"number"`
	}

	// Use atomic findOneAndUpdate to increment counter
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After).SetUpsert(true)
	err = numCollection.FindOneAndUpdate(ctx, bson.M{}, bson.M{"$inc": bson.M{"number": 1}}, opts).Decode(&currentNum)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to update counter"})
	}

	currentNumber := currentNum.Number
	zeroPadding := 4 - len(fmt.Sprintf("%d", currentNumber))
	users.TechnexID = "TX26"
	for i := 0; i < zeroPadding; i++ {
		users.TechnexID += "0"
	}
	users.TechnexID += fmt.Sprintf("%d", currentNumber)
	// var registeredEvents []string
	users.RegisteredEvents = []string{}
	users.Tickets = []models.Ticket{}

	if err := c.BodyParser(users); err != nil {
		return utils.ResponseMsg(c, 400, err.Error(), nil)
	}

	users.Email = normalizeEmail(users.Email)
	users.Username = normalizeUsername(users.Username)

	if users.Email == "" || users.Username == "" {
		return utils.ResponseMsg(c, 400, "username and email are required", nil)
	}

	if users.Password == "" {
		return utils.ResponseMsg(c, 400, "password is required", nil)
	}

	var existingUser models.Users
	err = db.Collection("users").FindOne(ctx, bson.M{"email": users.Email}).Decode(&existingUser)
	if err == nil {
		return c.Status(409).JSON(fiber.Map{"message": "Email already exists"})
	}
	if err != nil && err != mongo.ErrNoDocuments {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to validate email uniqueness"})
	}

	usernameFilter := usernameCaseInsensitiveFilter(users.Username)
	if len(usernameFilter) == 0 {
		return utils.ResponseMsg(c, 400, "username cannot be empty", nil)
	}
	err = db.Collection("users").FindOne(ctx, usernameFilter).Decode(&existingUser)
	if err == nil {
		return c.Status(409).JSON(fiber.Map{"message": "Username already exists"})
	}
	if err != nil && err != mongo.ErrNoDocuments {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to validate username uniqueness"})
	}

	users.Password = utils.HashPassword(users.Password)
	if r, err := db.Collection("users").InsertOne(ctx, users); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	} else {
		return c.Status(201).JSON(fiber.Map{"id": r.InsertedID, "technexId": users.TechnexID})
	}
}
