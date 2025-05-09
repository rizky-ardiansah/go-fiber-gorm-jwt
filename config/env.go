package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// GetEnv mengambil nilai dari environment variable atau mengembalikan nilai default jika tidak ada.
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// LoadEnv memuat variabel dari file .env
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}
}
