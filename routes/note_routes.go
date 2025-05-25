package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizky-ardiansah/go-fiber-gorm-jwt/handlers"
	"github.com/rizky-ardiansah/go-fiber-gorm-jwt/middlewares"
)

func SetupNoteRoutes(app *fiber.App) {
	note := app.Group("/notes")

	note.Use(middlewares.Protected())

	note.Post("/", handlers.CreateNote)
	note.Get("/", handlers.GetNotes)
	note.Get("/:id", handlers.GetNote)
	note.Put("/:id", handlers.UpdateNote)
	note.Delete("/:id", handlers.DeleteNote)
}
