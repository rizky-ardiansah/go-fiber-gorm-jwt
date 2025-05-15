package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizky-ardiansah/go-fiber-gorm-jwt/handlers"
	"github.com/rizky-ardiansah/go-fiber-gorm-jwt/middlewares"
)

// SetupUserRoutes mengatur rute-rute yang berhubungan dengan data pengguna (terproteksi)
func SetupUserRoutes(app *fiber.App) {
	api := app.Group("/api/users")

	// Rute ini memerlukan autentikasi JWT
	api.Get("/me", middlewares.Protected(), handlers.GetMyProfile)
}
