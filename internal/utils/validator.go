package utils

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

// ValidateRequired checks if required fields are present
func ValidateRequired(data map[string]string, fields ...string) []string {
	var missing []string
	for _, field := range fields {
		if strings.TrimSpace(data[field]) == "" {
			missing = append(missing, field)
		}
	}
	return missing
}

// ParseBody parses the request body into the given struct
func ParseBody(c *fiber.Ctx, out interface{}) error {
	if err := c.BodyParser(out); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}
	return nil
}
