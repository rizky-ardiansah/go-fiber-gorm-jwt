package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizky-ardiansah/go-fiber-gorm-jwt/handlers"
)

func SetupAuthRoutes(app *fiber.App) {
	api := app.Group("/api/auth") // Grup rute dengan prefix /api/auth

	api.Post("/register", handlers.RegisterUser)
	api.Post("/login", handlers.LoginUser)
	api.Post("/logout", handlers.LogoutUser)

}
