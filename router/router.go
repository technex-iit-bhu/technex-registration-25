package router

import (
	"technexRegistration/handler"
	event_handler "technexRegistration/handler/events"
	payments_handler "technexRegistration/handler/payments"
	user_handler "technexRegistration/handler/user"
	workshop_handler "technexRegistration/handler/workshops"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Route(app *fiber.App) {

	app.Use(cors.New())
	app.Get("/", handler.Hello)
	api := app.Group("/api", logger.New())
	api.Get("/", handler.HelloAPI)

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
	user.Post("/verify-qr", user_handler.VerifyQR)
	user.Post("/send-otp", user_handler.SendOTP)
    user.Post("/verify-otp", user_handler.VerifyOTP)
	user.Post("/reset-password", user_handler.ResetPassword)

	events := api.Group("/events")
	events.Get("/", event_handler.GetAllEvents)
	events.Get("/getEvent", event_handler.GetEventDetails)
	events.Get("/getEventByID", event_handler.GetEventByID)
	events.Post("/insertEvent", event_handler.InsertEvent)
	events.Post("/insertEvents", event_handler.BulkInsertEvents)
	events.Delete("/deleteEvent", event_handler.DeleteEvent)
	events.Patch("/updateEvent", event_handler.UpdateEvent)
	events.Patch("/updateSubEvents", event_handler.UpdateSubEvents)
	events.Get("/subevents", event_handler.GetSubEvents)
	events.Get("/subevent-by-name", event_handler.GetSubEventByName)

	workshops := api.Group("/workshops")
	workshops.Get("/", workshop_handler.GetAllWorkshops)
	workshops.Get("/getWorkshop", workshop_handler.GetWorkshopDetails)
	workshops.Get("/getWorkshopByID", workshop_handler.GetWorkshopByID)
	workshops.Post("/insertWorkshop", workshop_handler.InsertWorkshop)
	workshops.Post("/insertWorkshops", workshop_handler.BulkInsertWorkshops)
	workshops.Delete("/deleteWorkshop", workshop_handler.DeleteWorkshop)
	workshops.Patch("/updateWorkshop", workshop_handler.UpdateWorkshop)
	workshops.Patch("/updateSubWorkshops/", workshop_handler.UpdateSubWorkshops)
	workshops.Get("/subworkshops", workshop_handler.GetSubWorkshops)

	payments := api.Group("/payments")
	payments.Post("/", payments_handler.CapturePayments)
}
