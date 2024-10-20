package events

import (
	"context"
	"time"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"technexRegistration/database"
)

type Event struct {
	ID         string    `json:"id" bson:"_id,omitempty"`
	Name       string    `json:"name" bson:"name"`
	Desc       string    `json:"desc" bson:"description"`
	Start_Date time.Time `json:"sDate" bson:"startDate"`
	End_Date   time.Time `json:"eDate" bson:"endDate"`
}

func GetAllEvents(c *fiber.Ctx) error {
	var ctx = context.Background()

	token := c.Get("Authorization")[7:]
	if token == "" {
		return c.Status(401).JSON(fiber.Map{"message": "Unauthorized"})
	}

	db, err := database.Connect()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	cursor, err := db.Collection("events").Find(ctx, bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	defer cursor.Close(ctx)

	var events []Event

	for cursor.Next(ctx) {
		var event Event
		if err := cursor.Decode(&event); err != nil {
			return c.Status(500).JSON(fiber.Map{"message": err.Error()})
		}
		events = append(events, event)
	}

	if err := cursor.Err(); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"events": events})
}
