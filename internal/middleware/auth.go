package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/golang-jwt/jwt/v5"
	"hospital-app-backend/internal/config"
)

// Protected middleware for JWT authentication
func Protected(cfg *config.Config) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(cfg.JWTSecret)},
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"success": false,
		"error":   "Unauthorized: " + err.Error(),
	})
}

// GenerateToken creates a new JWT token
func GenerateToken(cfg *config.Config, userID, email, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = jwt.NewNumericDate(time.Now().AddDate(0, 0, cfg.JWTExpiryHours))

	return token.SignedString([]byte(cfg.JWTSecret))
}

// GetUserID extracts user ID from JWT token
func GetUserID(c *fiber.Ctx) string {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["user_id"].(string)
}

// GetUserRole extracts user role from JWT token
func GetUserRole(c *fiber.Ctx) string {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["role"].(string)
}

// AdminOnly restricts access to admin users
func AdminOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := GetUserRole(c)
		if role != "admin" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"success": false,
				"error":   "Admin access required",
			})
		}
		return c.Next()
	}
}
