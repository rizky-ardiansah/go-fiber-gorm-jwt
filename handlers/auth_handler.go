package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizky-ardiansah/go-fiber-gorm-jwt/config"
	"github.com/rizky-ardiansah/go-fiber-gorm-jwt/models"
)

// RegisterUserInput defines the expected request body for user registration
type RegisterUserInput struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func RegisterUser(c *fiber.Ctx) error {
	// Parse the request body into RegisterUserInput
	input := new(RegisterUserInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "cannot parse request body",
			"data":    err.Error(),
		})
	}

	// Validate that Name, Email, and Password are provided
	if input.Name == "" || input.Email == "" || input.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "name, email, and password are required",
		})
	}

	// Create a models.User instance from the input
	user := &models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password, // Password from input is used here
	}

	// Hash password
	if err := user.HashPassword(user.Password); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "failed to hash password",
			"data":    err.Error(),
		})
	}

	// Save user to database
	result := config.DB.Create(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "failed to create user",
			"data":    result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "user registered successfully",
		"data":    user, // models.User is returned; Password field will be omitted due to json:"-"
	})
}
