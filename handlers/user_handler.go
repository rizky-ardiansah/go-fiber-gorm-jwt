package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizky-ardiansah/go-fiber-gorm-jwt/config"
	"github.com/rizky-ardiansah/go-fiber-gorm-jwt/models"
)

// GetMyProfile adalah handler untuk mendapatkan profil user yang sedang login
func GetMyProfile(c *fiber.Ctx) error {
	// Ambil userID dari context yang sudah di-set oleh middleware
	userID := c.Locals("userID").(uint) // type assertion

	var user models.User
	result := config.DB.First(&user, userID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "User not found",
		})
	}

	// Jangan kirim password
	user.Password = ""

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Profile fetched successfully",
		"data":    user,
	})
}
