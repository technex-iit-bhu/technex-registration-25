package events

import (
    "context"
    "fmt"

    "technexRegistration/database"
    "technexRegistration/models"
    "technexRegistration/utils"

    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
)

// GET /api/events/subevent-by-name?name=<subEventName>
func GetSubEventByName(c *fiber.Ctx) error {
    // 1) Get subEvent name from query params
    subEventName := c.Query("name")
    if subEventName == "" {
        return utils.ResponseMsg(c, 400, "SubEvent name is required", nil)
    }

    // 2) Connect to DB
    ctx := context.Background()
    db, err := database.Connect()
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"message": err.Error()})
    }

    // 3) Get all events
    cursor, err := db.Collection("events").Find(ctx, bson.M{})
    if err != nil {
        return utils.ResponseMsg(c, 500, "Database error", nil)
    }
    defer cursor.Close(ctx)

    // 4) Iterate through all events and their subEvents
    var events []models.Event
    if err = cursor.All(ctx, &events); err != nil {
        return utils.ResponseMsg(c, 500, "Error parsing events", nil)
    }

    // 5) Search for the subEvent
    for _, event := range events {
        for _, sub := range event.SubEvents {
            if sub.Name == subEventName {
                fmt.Printf("Found subEvent '%s' in event '%s'\n", subEventName, event.Name)
                return c.Status(200).JSON(fiber.Map{
                    "subEvent": sub,
                    "parentEvent": fiber.Map{
                        "id":   event.ID,
                        "name": event.Name,
                    },
                })
            }
        }
    }

    // 6) Not found
    return utils.ResponseMsg(c, 404, "No subEvent found with that name", nil)
}
