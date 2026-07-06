package utils

import "github.com/gofiber/fiber/v2"

// SuccessResponse sends a success JSON response
func SuccessResponse(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

// ErrorResponse sends an error JSON response
func ErrorResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"success": false,
		"error":   message,
	})
}

// PaginatedResponse sends a paginated JSON response
func PaginatedResponse(c *fiber.Ctx, data interface{}, total, page, limit int) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    data,
		"pagination": fiber.Map{
			"total":        total,
			"page":         page,
			"limit":        limit,
			"total_pages":  (total + limit - 1) / limit,
		},
	})
}

// MessageResponse sends a simple message response
func MessageResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"success": true,
		"message": message,
	})
}
