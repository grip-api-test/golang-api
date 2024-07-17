package models

import "gorm.io/gorm"

// defining a data model for posts.
type Post struct {
	gorm.Model 
	Title string
	Body string
}
