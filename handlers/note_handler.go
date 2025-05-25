package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rizky-ardiansah/go-fiber-gorm-jwt/config"
	"github.com/rizky-ardiansah/go-fiber-gorm-jwt/models"
)

func CreateNote(c *fiber.Ctx) error {
	userIDlocal := c.Locals("userID")
	if userIDlocal == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized: User ID not found in context"})
	}
	userID, ok := userIDlocal.(uint)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error: User ID is of an invalid type"})
	}

	note := new(models.Note)
	if err := c.BodyParser(note); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	note.UserID = userID

	if err := config.DB.Create(&note).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create note"})
	}

	return c.Status(fiber.StatusCreated).JSON(note)
}

func GetNotes(c *fiber.Ctx) error {
	userIDlocal := c.Locals("userID")
	if userIDlocal == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized: User ID not found in context"})
	}
	userID, ok := userIDlocal.(uint)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error: User ID is of an invalid type"})
	}

	var notes []models.Note
	if err := config.DB.Where("user_id = ?", userID).Find(&notes).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not retrieve notes"})
	}

	// Check if no notes found
	if len(notes) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "No notes found for this user"})
	}

	return c.JSON(notes)
}

func GetNote(c *fiber.Ctx) error {
	userIDlocal := c.Locals("userID")
	if userIDlocal == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized: User ID not found in context"})
	}
	userID, ok := userIDlocal.(uint)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error: User ID is of an invalid type"})
	}

	id := c.Params("id")
	noteID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid note ID"})
	}

	var note models.Note
	if err := config.DB.Where("id = ? AND user_id = ?", noteID, userID).First(&note).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Note not found"})
	}

	return c.JSON(note)
}

func UpdateNote(c *fiber.Ctx) error {
	userIDlocal := c.Locals("userID")
	if userIDlocal == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized: User ID not found in context"})
	}
	userID, ok := userIDlocal.(uint)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error: User ID is of an invalid type"})
	}

	id := c.Params("id")
	noteID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid note ID"})
	}

	var note models.Note
	if err := config.DB.Where("id = ? AND user_id = ?", noteID, userID).First(&note).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Note not found"})
	}

	updatedNote := new(models.Note)
	if err := c.BodyParser(updatedNote); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	note.Title = updatedNote.Title
	note.Content = updatedNote.Content

	if err := config.DB.Save(&note).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not update note"})
	}

	return c.JSON(note)
}

func DeleteNote(c *fiber.Ctx) error {
	userIDlocal := c.Locals("userID")
	if userIDlocal == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized: User ID not found in context"})
	}
	userID, ok := userIDlocal.(uint)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error: User ID is of an invalid type"})
	}

	id := c.Params("id")
	noteID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid note ID"})
	}

	var note models.Note
	if err := config.DB.Where("id = ? AND user_id = ?", noteID, userID).First(&note).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Note not found"})
	}

	if err := config.DB.Delete(&note).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not delete note"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
