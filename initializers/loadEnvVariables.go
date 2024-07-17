package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnvVariables loads environment variables from a .env file.
func LoadEnvVariables() {
	// loads the environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		// If there's an error loading the .env file, log the error.
		log.Fatal("Error loading .env file")
	}
}
