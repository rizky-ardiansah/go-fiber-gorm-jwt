package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizky-ardiansah/go-fiber-gorm-jwt/utils"
	"log"
)

func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Debug information
		log.Printf("Request Headers: %v", c.GetReqHeaders())
		log.Printf("JWT Cookie: %v", c.Cookies("jwt"))

		// Ambil token dari HTTP-only cookie
		tokenString := c.Cookies("jwt")
		if tokenString == "" {
			log.Printf("JWT cookie is empty or not found")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "error",
				"message": "Authentication required. No JWT cookie found.",
			})
		}

		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			log.Printf("JWT validation error: %v", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid or expired token",
				"data":    err.Error(),
			})
		}

		// Set informasi user ke context agar bisa diakses handler selanjutnya
		c.Locals("userID", claims.UserID)
		c.Locals("userEmail", claims.Email)

		return c.Next()
	}
}
