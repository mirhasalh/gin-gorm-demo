package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mirhasalh/gin-gorm-demo/initializers"
	"github.com/mirhasalh/gin-gorm-demo/models"
)

func PostsCreate(c *gin.Context) {
	// Get data from request body
	var body struct {
		Body  string `json:"body" binding:"required"`
		Title string `json:"title" binding:"required"`
	}

	c.Bind(&body)

	// Create a new post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return the post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get id from URL
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get id from URL
	id := c.Param("id")

	// Get data from request body
	var body struct {
		Body  string `json:"body" binding:"required"`
		Title string `json:"title" binding:"required"`
	}

	c.Bind(&body)

	// Update the post
	var post models.Post
	initializers.DB.First(&post, id)

	post.Title = body.Title
	post.Body = body.Body
	initializers.DB.Save(&post)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get id from URL
	id := c.Param("id")

	// Delete the post
	var post models.Post
	initializers.DB.First(&post, id)
	initializers.DB.Delete(&post)

	c.Status(204)
}
