package controllers

import (
	"example.com/go-crud/initializers"
	"example.com/go-crud/models"
	"github.com/gin-gonic/gin"
)

// just for test
func Add(a, b int) int {
    return a + b
}

func PostsCreate(c *gin.Context) {
	// Get data from the request body.
	var body struct {
		Body  string
		Title string
	}
	
	// Bind the request body data to the body struct.
	c.Bind(&body)

	// Create a new post with the data.
	post := models.Post{Title: body.Title, Body: body.Body}

	// Save the new post to the database.
	result := initializers.DB.Create(&post)

	// If there was an error, respond with status 400.
	if result.Error != nil {
		c.Status(400)
		return
	}

	// Respond with the created post.
	c.JSON(200, gin.H{
		"message": post,
	})
}

// PostIndex handles fetching all posts.
func PostIndex(c *gin.Context) {
	// Retrieve all posts from the database.
	var posts []models.Post
	initializers.DB.Find(&posts)
	
	// Respond with the list of posts.
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

// PostShow handles fetching a single post by its ID.
func PostShow(c *gin.Context) {
	// Get the post ID from the URL parameters.
	id := c.Param("id")
	
	// Retrieve the post from the database.
	var post models.Post
	initializers.DB.First(&post, id)
	
	// Respond with the post.
	c.JSON(200, gin.H{
		"posts": post,
	})
}

// PostUpdate handles updating an existing post.
func PostUpdate(c *gin.Context) {
	// Get the post ID from the URL parameters.
	id := c.Param("id")
	
	// Get data from the request body.
	var body struct {
		Body  string
		Title string
	}

	// Bind the request body data to the body struct.
	c.Bind(&body)
	
	// Retrieve the post to be updated from the database.
	var post models.Post
	initializers.DB.First(&post, id)
	
	// Update the post with the new data.
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title, 
		Body: body.Body,
	})
	
	// Respond with the updated post.
	c.JSON(200, gin.H{
		"message": post,
	})
}

// PostDelete handles deleting a post by its ID.
func PostDelete(c *gin.Context) {
	// Get the post ID from the URL parameters.
	id := c.Param("id")
	
	// Delete the post from the database.
	initializers.DB.Delete(&models.Post{}, id)
	
	// Respond with status 200.
	c.Status(200)
}
