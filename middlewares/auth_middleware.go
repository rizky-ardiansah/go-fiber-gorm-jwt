package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizky-ardiansah/go-fiber-gorm-jwt/utils"
)

// Protected adalah middleware untuk melindungi rute yang memerlukan autentikasi
func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Ambil token dari HTTP-only cookie
		tokenString := c.Cookies("jwt")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "error",
				"message": "Authentication required. ",
			})
		}
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
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
