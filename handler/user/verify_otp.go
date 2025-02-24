package user

import (
    "context"
    "net/http"
    "time"
    "technexRegistration/database"
    "technexRegistration/models"
    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
)

func VerifyOTP(c *fiber.Ctx) error {
    var body struct {
        Email string `json:"email"`
        OTP   int    `json:"otp"`
    }

    if err := c.BodyParser(&body); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "Invalid request body",
        })
    }

    db, err := database.Connect()
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
    }

    // Check if user exists
    var user models.Users
    err = db.Collection("users").FindOne(context.Background(), bson.D{
        {Key: "email", Value: body.Email},
    }).Decode(&user)
    if err != nil {
        return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "user does not exist"})
    }

    // Now lookup the OTP in the "otps" collection
    var storedOtp models.Otp
    filter := bson.M{
        "userId": user.ID,
        "code":   body.OTP,
        "used":   false,
    }
    err = db.Collection("otps").FindOne(context.Background(), filter).Decode(&storedOtp)
    if err != nil {
        return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid or used OTP"})
    }

    if time.Now().After(storedOtp.ExpiresAt) {
        return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "OTP has expired"})
    }

    _, err = db.Collection("otps").UpdateOne(
        context.Background(),
        bson.M{"_id": storedOtp.ID},
        bson.M{"$set": bson.M{"used": true}},
    )
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "message": "Error marking OTP as used",
        })
    }

    return c.Status(http.StatusOK).JSON(fiber.Map{
        "message": "OTP verified successfully",
    })
}
