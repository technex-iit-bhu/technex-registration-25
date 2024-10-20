package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"technexRegistration/handler"
	user_handler "technexRegistration/handler/user"
	event_handler "technexRegistration/handler/events"
)

func Route(app *fiber.App) {

	api := app.Group("/api", logger.New())
	app.Use(cors.New())
	api.Get("/", handler.Hello)

	user := api.Group("/user")
	user.Post("/register", user_handler.CreateUsers)
	user.Get("/profile", user_handler.GetUserFromToken)
	user.Post("/login/password", user_handler.LoginWithPassword)
	user.Post("/login/google", user_handler.LoginWithGoogle)
	user.Post("/login/github", user_handler.LoginWithGithub)
	user.Delete("/delete", user_handler.DeleteUser)
	user.Patch("/update", user_handler.UpdateDetails)
	user.Get("/recovery/:username", user_handler.SendRecoveryEmail)
	user.Post("/verify_recovery_and_update_password", user_handler.UpdatePassword)

	events := api.Group("/events")
	events.Get("/", event_handler.GetAllEvents)
	events.Get("/:name", event_handler.GetEventDetails)
	events.Post("/insertEvent", event_handler.InsertEvent)
}
