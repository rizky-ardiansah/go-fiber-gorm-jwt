package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/rizky-ardiansah/go-fiber-gorm-jwt/config"
	"github.com/rizky-ardiansah/go-fiber-gorm-jwt/handlers"
	"github.com/rizky-ardiansah/go-fiber-gorm-jwt/models"
	"github.com/rizky-ardiansah/go-fiber-gorm-jwt/routes"
	"github.com/stretchr/testify/assert"
)

// Helper function to setup the Fiber app for testing
func setupApp() *fiber.App {
	// Load environment variables for test
	os.Setenv("DB_HOST", "localhost") // Or your test DB config
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "postgres")
	os.Setenv("DB_NAME", "fiber_api_db_test")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("JWT_SECRET_KEY", "kmzway87aa")
	os.Setenv("DB_SSLMODE", "disable")
	os.Setenv("DB_TIMEZONE", "Asia/Jakarta")
	os.Setenv("JWT_EXPIRES_IN", "1h")

	config.LoadEnv()
	config.ConnectDB()

	// Drop existing tables and migrate for a clean test environment
	config.DB.Migrator().DropTable(&models.User{}, &models.Note{})
	err := config.DB.AutoMigrate(&models.User{}, &models.Note{})
	if err != nil {
		panic("Failed to migrate database for testing: " + err.Error())
	}

	app := fiber.New()
	routes.SetupAuthRoutes(app)
	return app
}

func TestLoginE2E(t *testing.T) {
	app := setupApp()

	registerInput := handlers.RegisterUserInput{
		Name:     "Test User Login",
		Email:    "testlogin@example.com",
		Password: "password123",
	}
	registerBody, _ := json.Marshal(registerInput)

	reqRegister := httptest.NewRequest(http.MethodPost, "/api/auth/register", bytes.NewBuffer(registerBody))
	reqRegister.Header.Set("Content-Type", "application/json")
	respRegister, errRegister := app.Test(reqRegister, -1)
	assert.NoError(t, errRegister)
	assert.Equal(t, http.StatusCreated, respRegister.StatusCode)

	loginInput := handlers.LoginUserInput{
		Email:    "testlogin@example.com",
		Password: "password123",
	}
	loginBody, _ := json.Marshal(loginInput)

	reqLogin := httptest.NewRequest(http.MethodPost, "/api/auth/login", bytes.NewBuffer(loginBody))
	reqLogin.Header.Set("Content-Type", "application/json")

	respLogin, errLogin := app.Test(reqLogin, -1)
	assert.NoError(t, errLogin)

	assert.Equal(t, http.StatusOK, respLogin.StatusCode)

	var loginResponseBody map[string]interface{}
	errDecode := json.NewDecoder(respLogin.Body).Decode(&loginResponseBody)
	assert.NoError(t, errDecode)

	assert.Equal(t, "success", loginResponseBody["status"])
	assert.Equal(t, "Login successful", loginResponseBody["message"])

	// Check for token in response data
	responseData, ok := loginResponseBody["data"].(map[string]interface{})
	assert.True(t, ok, "Response data should be a map")
	assert.NotEmpty(t, responseData["token"], "Token should not be empty")

	// Check for JWT cookie
	foundCookie := false
	for _, cookie := range respLogin.Cookies() {
		if cookie.Name == "jwt" {
			foundCookie = true
			assert.NotEmpty(t, cookie.Value, "JWT cookie value should not be empty")
			break
		}
	}
	assert.True(t, foundCookie, "JWT cookie should be set")

	sqlDB, _ := config.DB.DB()
	sqlDB.Close()
}

func TestLoginE2E_InvalidCredentials(t *testing.T) {
	app := setupApp()

	registerInput := handlers.RegisterUserInput{
		Name:     "Test User Invalid Login",
		Email:    "testinvalidlogin@example.com",
		Password: "password123",
	}
	registerBody, _ := json.Marshal(registerInput)
	reqRegister := httptest.NewRequest(http.MethodPost, "/api/auth/register", bytes.NewBuffer(registerBody))
	reqRegister.Header.Set("Content-Type", "application/json")
	respRegister, _ := app.Test(reqRegister, -1)
	assert.Equal(t, http.StatusCreated, respRegister.StatusCode)

	loginInput := handlers.LoginUserInput{
		Email:    "testinvalidlogin@example.com",
		Password: "password125",
	}
	loginBody, _ := json.Marshal(loginInput)

	reqLogin := httptest.NewRequest(http.MethodPost, "/api/auth/login", bytes.NewBuffer(loginBody))
	reqLogin.Header.Set("Content-Type", "application/json")

	respLogin, errLogin := app.Test(reqLogin, -1)
	assert.NoError(t, errLogin)

	assert.Equal(t, http.StatusUnauthorized, respLogin.StatusCode)

	var responseBody map[string]interface{}
	errDecode := json.NewDecoder(respLogin.Body).Decode(&responseBody)
	assert.NoError(t, errDecode)

	assert.Equal(t, "error", responseBody["status"])
	assert.Equal(t, "Invalid email or password", responseBody["message"])

	sqlDB, _ := config.DB.DB()
	sqlDB.Close()
}

func TestLoginE2E_UserNotFound(t *testing.T) {
	app := setupApp()

	loginInput := handlers.LoginUserInput{
		Email:    "nonexistentuser@example.com",
		Password: "password123",
	}
	loginBody, _ := json.Marshal(loginInput)

	reqLogin := httptest.NewRequest(http.MethodPost, "/api/auth/login", bytes.NewBuffer(loginBody))
	reqLogin.Header.Set("Content-Type", "application/json")

	respLogin, errLogin := app.Test(reqLogin, -1)
	assert.NoError(t, errLogin)

	assert.Equal(t, http.StatusUnauthorized, respLogin.StatusCode)

	var responseBody map[string]interface{}
	errDecode := json.NewDecoder(respLogin.Body).Decode(&responseBody)
	assert.NoError(t, errDecode)

	assert.Equal(t, "error", responseBody["status"])
	assert.Equal(t, "Invalid email or password", responseBody["message"])

	sqlDB, _ := config.DB.DB()
	sqlDB.Close()
}
