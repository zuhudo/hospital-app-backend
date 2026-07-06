package handlers

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"hospital-app-backend/internal/config"
	"hospital-app-backend/internal/middleware"
	"hospital-app-backend/internal/models"
	"hospital-app-backend/internal/utils"
)

// In-memory store — in production, use a database
var users []models.User

type AuthHandler struct {
	Config *config.Config
}

func NewAuthHandler(cfg *config.Config) *AuthHandler {
	h := &AuthHandler{Config: cfg}
	h.initAdminUser()
	return h
}

// initAdminUser creates default admin from env vars if configured
func (h *AuthHandler) initAdminUser() {
	adminEmail := os.Getenv("ADMIN_EMAIL")
	adminPassword := os.Getenv("ADMIN_PASSWORD")

	if adminEmail == "" || adminPassword == "" {
		log.Println("No ADMIN_EMAIL/ADMIN_PASSWORD set — no default admin user")
		return
	}

	// Check if admin already exists
	for _, u := range users {
		if u.Email == adminEmail {
			return
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to hash admin password: %v", err)
		return
	}

	admin := models.User{
		ID:        uuid.New().String(),
		FirstName: "Super",
		LastName:  "Admin",
		Email:     adminEmail,
		Phone:     os.Getenv("ADMIN_PHONE"),
		Password:  string(hashedPassword),
		Role:      "admin",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	users = append(users, admin)
	log.Printf("✅ Default admin user created: %s", adminEmail)
}

// Login handles user authentication
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req models.LoginRequest
	if err := utils.ParseBody(c, &req); err != nil {
		return err
	}

	// Find user by email
	var user *models.User
	for _, u := range users {
		if u.Email == req.Email {
			user = &u
			break
		}
	}

	if user == nil {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid email or password")
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid email or password")
	}

	// Generate token
	token, err := middleware.GenerateToken(h.Config, user.ID, user.Email, user.Role)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to generate token")
	}

	return utils.SuccessResponse(c, fiber.StatusOK, models.AuthResponse{
		Token: token,
		User:  *user,
	})
}

// Register creates a new user account
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req models.RegisterRequest
	if err := utils.ParseBody(c, &req); err != nil {
		return err
	}

	// Check if email exists
	for _, u := range users {
		if u.Email == req.Email {
			return utils.ErrorResponse(c, fiber.StatusConflict, "Email already registered")
		}
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to hash password")
	}

	role := req.Role
	if role == "" {
		role = "patient"
	}

	user := models.User{
		ID:        uuid.New().String(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
		Password:  string(hashedPassword),
		Role:      role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	users = append(users, user)

	// Generate token
	token, err := middleware.GenerateToken(h.Config, user.ID, user.Email, user.Role)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to generate token")
	}

	return utils.SuccessResponse(c, fiber.StatusCreated, models.AuthResponse{
		Token: token,
		User:  user,
	})
}
