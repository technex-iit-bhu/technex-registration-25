package user

import (
    "context"
    "technexRegistration/database"
    "technexRegistration/models"
    "net/http"

    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
    "golang.org/x/crypto/bcrypt"
)

func ResetPassword(c *fiber.Ctx) error {
    var body struct {
        Email       string `json:"email"`
        NewPassword string `json:"newPassword"`
    }

    if err := c.BodyParser(&body); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "message": "Invalid request body",
        })
    }

    // Connect to DB
    db, err := database.Connect()
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "message": err.Error(),
        })
    }

    // Check if user exists
    var user models.Users
    err = db.Collection("users").FindOne(context.Background(), bson.M{"email": body.Email}).Decode(&user)
    if err != nil {
        return c.Status(http.StatusNotFound).JSON(fiber.Map{
            "message": "User not found",
        })
    }

    // Hash the new password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.NewPassword), bcrypt.DefaultCost)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "message": "Failed to hash password",
        })
    }

    // Update user's password in DB
    _, err = db.Collection("users").UpdateOne(
        context.Background(),
        bson.M{"_id": user.ID},
        bson.M{"$set": bson.M{"password": string(hashedPassword)}},
    )
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "message": "Failed to update password",
        })
    }

    return c.Status(http.StatusOK).JSON(fiber.Map{
        "message": "Password reset successfully",
    })
}
