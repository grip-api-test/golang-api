package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// global variable that will hold the database connection.
var DB *gorm.DB

func ConnectToDB() {
	var err error
	// Retrieve the database URL from environment variables.
	dsn := os.Getenv("DB_URL")
	// Open a connection to the database
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		// If there is an error connecting to the database, log the error.
		log.Fatal("Failed to connect to database!")
	}
}
