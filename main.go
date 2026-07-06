package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/joho/godotenv"
	"hospital-app-backend/internal/config"
	"hospital-app-backend/internal/routes"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using defaults")
	}

	// Load config
	cfg := config.Load()

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName:      "Hospital API v1.0",
		ErrorHandler: customErrorHandler,
	})

	// Global middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.AllowedOrigins,
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: cfg.AllowedOrigins != "*",
	}))

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"service": "Hospital API",
			"version": "1.0.0",
		})
	})

	// Setup routes
	api := app.Group("/api")
	routes.SetupAuthRoutes(api, cfg)
	routes.SetupPatientRoutes(api, cfg)
	routes.SetupDoctorRoutes(api, cfg)
	routes.SetupAppointmentRoutes(api, cfg)
	routes.SetupRecordRoutes(api, cfg)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("🏥 Hospital API running on port %s", port)
	log.Fatal(app.Listen(":" + port))
}

func customErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return c.Status(code).JSON(fiber.Map{
		"success": false,
		"error":   err.Error(),
	})
}
