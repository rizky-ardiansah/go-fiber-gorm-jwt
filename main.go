package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rizky-ardiansah/go-fiber-gorm-jwt/config" // Sesuaikan dengan path modul Anda
	"github.com/rizky-ardiansah/go-fiber-gorm-jwt/models"
	"github.com/rizky-ardiansah/go-fiber-gorm-jwt/routes"
)

func main() {
	// Muat konfigurasi environment
	config.LoadEnv()

	// Koneksi ke database
	config.ConnectDB()

	err := config.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migrated successfully")

	app := fiber.New()

	app.Use(logger.New())

	// Setup routes
	routes.SetupAuthRoutes(app)
	routes.SetupUserRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World from Fiber!")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Printf("Server is starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
