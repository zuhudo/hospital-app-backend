package routes

import (
	"github.com/gofiber/fiber/v2"
	"hospital-app-backend/internal/config"
	"hospital-app-backend/internal/handlers"
	"hospital-app-backend/internal/middleware"
)

func SetupAuthRoutes(api fiber.Router, cfg *config.Config) {
	auth := api.Group("/auth")
	handler := handlers.NewAuthHandler(cfg)

	auth.Post("/login", handler.Login)
	auth.Post("/register", handler.Register)
}

func SetupPatientRoutes(api fiber.Router, cfg *config.Config) {
	patients := api.Group("/patients", middleware.Protected(cfg))
	handler := handlers.NewPatientHandler()

	patients.Get("/", handler.GetAll)
	patients.Get("/:id", handler.GetByID)
	patients.Post("/", handler.Create)
	patients.Put("/:id", handler.Update)
	patients.Delete("/:id", handler.Delete)
}

func SetupDoctorRoutes(api fiber.Router, cfg *config.Config) {
	doctors := api.Group("/doctors", middleware.Protected(cfg))
	handler := handlers.NewDoctorHandler()

	doctors.Get("/", handler.GetAll)
	doctors.Get("/:id", handler.GetByID)
	doctors.Post("/", handler.Create)
}

func SetupAppointmentRoutes(api fiber.Router, cfg *config.Config) {
	appointments := api.Group("/appointments", middleware.Protected(cfg))
	handler := handlers.NewAppointmentHandler()

	appointments.Get("/", handler.GetAll)
	appointments.Get("/:id", handler.GetByID)
	appointments.Post("/", handler.Create)
	appointments.Put("/:id/cancel", handler.Cancel)
}

func SetupRecordRoutes(api fiber.Router, cfg *config.Config) {
	records := api.Group("/records", middleware.Protected(cfg))
	handler := handlers.NewRecordHandler()

	records.Get("/:patientId", handler.GetByPatientID)
	records.Post("/", handler.Create)
}
