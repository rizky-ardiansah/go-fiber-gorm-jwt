package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/rizky-ardiansah/go-fiber-gorm-jwt/config"
	"github.com/rizky-ardiansah/go-fiber-gorm-jwt/models"
	"github.com/rizky-ardiansah/go-fiber-gorm-jwt/utils"
	"gorm.io/gorm"
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

	// http only cookie
	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "failed to generate authentication token",
			"data":    err.Error(),
		})
	}

	// Set cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Path:     "/",
		MaxAge:   86400, // 1 day in seconds
		Secure:   true,  // Use true in production with HTTPS
		HTTPOnly: true,
		SameSite: "Lax", // Lax for most use cases, Strict for higher security
	}
	c.Cookie(&cookie)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "user registered successfully",
		"data":    user, // models.User is returned; Password field will be omitted due to json:"-"
	})
}

type LoginUserInput struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func LoginUser(c *fiber.Ctx) error {
	input := new(LoginUserInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"data":    err.Error(),
		})
	}

	if input.Email == "" || input.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Email and Password are required",
		})
	}

	var user models.User
	// Cari user berdasarkan email
	result := config.DB.Where("email = ?", input.Email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid email or password",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Database error",
			"data":    result.Error.Error(),
		})
	}

	// Verifikasi password
	if err := user.CheckPassword(input.Password); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid email or password",
		})
	}

	// Generate JWT
	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not generate token",
			"data":    err.Error(),
		})
	}

	// Set cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Path:     "/",
		MaxAge:   86400, // 1 day in seconds
		Secure:   true,  // Use true in production with HTTPS
		HTTPOnly: true,
		SameSite: "Lax", // Lax for most use cases, Strict for higher security
	}
	c.Cookie(&cookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Login successful",
		"data": fiber.Map{
			"token": token,
		},
	})
}

// LogoutUser handles user logout by invalidating the JWT cookie
func LogoutUser(c *fiber.Ctx) error {
	// Create an expired cookie to replace the current one
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "", // Empty value
		Path:     "/",
		MaxAge:   -1,   // Delete the cookie
		Secure:   true, // Use true in production with HTTPS
		HTTPOnly: true,
		SameSite: "Lax", // Lax for most use cases, Strict for higher security
	}

	// Set the expired cookie
	c.Cookie(&cookie)

	// Catatan: Jika aplikasi front-end juga menyimpan token di localStorage,
	// token tersebut harus dihapus di sisi klien karena server tidak dapat
	// menghapus data localStorage secara langsung

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Logged out successfully",
	})
}
