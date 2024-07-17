package main

import (
	"example.com/go-crud/controllers"
	"example.com/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	// LoadEnvVariables loads environment variables from .env file
	initializers.LoadEnvVariables()

	// establishes a connection to the database.
	initializers.ConnectToDB()
}

func main() {
	// creates a Gin router
	r := gin.Default()

	// Create a new post
	r.POST("/posts", controllers.PostsCreate)

	// Retrieve all posts
	r.GET("/posts", controllers.PostIndex)

	// GET Retrieve a specific post by ID
	r.GET("/post/:id", controllers.PostShow)

	// PUT Update a specific post by ID
	r.PUT("/posts/:id", controllers.PostUpdate)

	// DELETE Delete a specific post by ID
	r.DELETE("/posts/:id", controllers.PostDelete)

	// starts the HTTP server
	r.Run()
}
