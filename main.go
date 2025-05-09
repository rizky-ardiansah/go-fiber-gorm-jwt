package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/rizky-ardiansah/go-fiber-gorm-jwt/config" // Sesuaikan dengan path modul Anda
	"github.com/rizky-ardiansah/go-fiber-gorm-jwt/models"
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
