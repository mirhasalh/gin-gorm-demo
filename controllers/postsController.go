package controllers

import "github.com/gin-gonic/gin"

func PostsCreate(c *gin.Context) {
	// Get data from request body

	// Create a new post

	// Return the post
	c.JSON(200, gin.H{ "message": "Post created" })
}