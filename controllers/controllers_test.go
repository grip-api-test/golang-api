package controllers

import (
	"testing"

	"example.com/go-crud/initializers"
)


func init() {
    // Setup code to initialize the database and Gin if necessary.
    initializers.LoadEnvVariables()
    initializers.ConnectToDB()
}

func TestAddition(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Expected 5, but got %d", result)
    }
}