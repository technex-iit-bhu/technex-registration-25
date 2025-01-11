package user

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"technexRegistration/database"
	"technexRegistration/models"
	"technexRegistration/utils"
	"time"
)

func CreateUsers(c *fiber.Ctx) error {
	users := new(models.Users)
	var ctx = context.Background()
	db, err := database.Connect()

	if err != nil {
		log.Fatal(err.Error())
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	users.CreatedAt = time.Now()
	users.UpdatedAt = time.Now()

	numCollection := db.Collection("num")
	var currentNum struct {
		Number int `bson:"number"`
	}

	err = numCollection.FindOne(ctx, bson.M{}).Decode(&currentNum)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(500).JSON(fiber.Map{"message": "No document found in 'num' collection"})
		}
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	currentNumber := currentNum.Number + 1
	zeroPadding := 5 - len(fmt.Sprintf("%d", currentNumber))
	users.TechnexID = fmt.Sprintf("TX%s%d", string(make([]byte, zeroPadding)), currentNumber)

	// Update the number in the 'num' collection
	_, err = numCollection.UpdateOne(ctx, bson.M{}, bson.M{"$set": bson.M{"number": currentNumber}})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to update 'num' collection"})
	}

	if err := c.BodyParser(users); err != nil {
		return utils.ResponseMsg(c, 400, err.Error(), nil)
	} else {
		users.Password = utils.HashPassword(users.Password)
		if r, err := db.Collection("users").InsertOne(ctx, users); err != nil {
			return c.Status(500).JSON(fiber.Map{"message": err.Error()})
		} else {
			return c.Status(201).JSON(fiber.Map{"id": r.InsertedID, "technexId": users.TechnexID})
		}
	}
}
