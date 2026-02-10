package user

import (
    "context"
    "net/http"
    "technexRegistration/database"
    "technexRegistration/models"
    "technexRegistration/utils"
    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
)

type InstituteIdRegisterBody struct {
    Events []string `json:"events"`
}

func InstituteIdRegisterEvent(c *fiber.Ctx) error {

    token := c.Get("Authorization")
    if len(token) < 7 || token[:7] != "Bearer " {
        return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
            "message": "Missing or invalid authorization token",
        })
    }
    token = token[7:]

    username, err := utils.DeserialiseAccessToken(token)
    if err != nil {
        return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
            "message": "Invalid token",
        })
    }

    db, err := database.Connect()
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "message": "Database connection failed",
        })
    }

    var user models.Users
    err = db.Collection("users").FindOne(context.Background(), bson.D{
        {Key: "username", Value: username},
    }).Decode(&user)
    if err != nil {
        return c.Status(http.StatusNotFound).JSON(fiber.Map{
            "message": "User not found",
        })
    }
    if !user.IsInstitute {
        return c.Status(http.StatusForbidden).JSON(fiber.Map{
            "message": "This endpoint is only for institute students. Please use the payment gateway.",
        })
    }

    if !user.EmailVerified {
        return c.Status(http.StatusForbidden).JSON(fiber.Map{
            "message": "Please verify your email first",
        })
    }

    var body InstituteIdRegisterBody
    if err := c.BodyParser(&body); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "message": "Invalid request body",
        })
    }

    if len(body.Events) == 0 {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "message": "No events provided",
        })
    }

    ticket := models.Ticket{
        Name:          "Institute Student - Direct Registration",
        TicketURL:     "",
        InvoiceURL:    "",
        Accommodation: false,
    }

    update := bson.M{
        "$addToSet": bson.M{
            "registeredEvents": bson.M{
                "$each": body.Events,
            },
            "tickets": ticket,
        },
    }

    result, err := db.Collection("users").UpdateOne(
        context.Background(),
        bson.M{"_id": user.ID},
        update,
    )

    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "message": "Failed to register for events",
            "error":   err.Error(),
        })
    }

    if result.MatchedCount == 0 {
        return c.Status(http.StatusNotFound).JSON(fiber.Map{
            "message": "User not found",
        })
    }

    return c.Status(http.StatusOK).JSON(fiber.Map{
        "message":         "Successfully registered for events",
        "events":          body.Events,
        "accommodation":   false,
        "is_institute":    true,
    })
}