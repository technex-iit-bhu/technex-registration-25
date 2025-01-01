package events

import (
    "context"
    "fmt"

    "technexRegistration/database"
    "technexRegistration/models"
    "technexRegistration/utils"

    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

// GET /api/events/subevents?id=<eventID>&name=<subEventName?>
func GetSubEvents(c *fiber.Ctx) error {
    // 1) Validate Event ID
    id := c.Query("id")
    if id == "" {
        return utils.ResponseMsg(c, 400, "Event ID is required", nil)
    }
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return utils.ResponseMsg(c, 400, "Invalid Event ID", nil)
    }

    // 2) Connect to DB, fetch the parent event
    ctx := context.Background()
    db, err := database.Connect()
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"message": err.Error()})
    }
    var event models.Event
    err = db.Collection("events").FindOne(ctx, bson.M{"_id": objID}).Decode(&event)
    if err != nil {
        return utils.ResponseMsg(c, 404, "Event not found", nil)
    }

    // 3) Check if 'name' query param is provided
    subEventName := c.Query("name")
    if subEventName == "" {
        // No subEvent name => Return the entire subEvents array
        fmt.Println("Returning all subEvents")
        return c.Status(200).JSON(fiber.Map{"subEvents": event.SubEvents})
    }

    // 4) If name is provided => find that single subEvent
    for _, sub := range event.SubEvents {
        if sub.Name == subEventName {
            // Found the matching subEvent
            fmt.Println("Returning subEvent with name:", subEventName)
            return c.Status(200).JSON(fiber.Map{"subEvent": sub})
        }
    }

    // 5) Not found
    return utils.ResponseMsg(c, 404, "No subEvent found with that name", nil)
}
