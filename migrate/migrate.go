// Automatically creates or updates the database table for the Post model to ensure it matches the current definitions.
package main

import (
	"example.com/go-crud/initializers"
	"example.com/go-crud/models"
)

func init() {
	// LoadEnvVariables loads environment variables from a .env file
	initializers.LoadEnvVariables()

	// connection to the database.
	initializers.ConnectToDB()
}

func main() {
	// AutoMigrate creates or updates database tables based on the defined model structs.
	initializers.DB.AutoMigrate(&models.Post{})
}
