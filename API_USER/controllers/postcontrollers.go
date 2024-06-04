package controllers

import (
	"api/initializers"
	"api/models"
	"api/repositories"

	"github.com/gin-gonic/gin"
)

var postRepository = repositories.NewPostRepository(initializers.DB)

func Create(c *gin.Context) {
	// create a post
	var post models.Post
	if err := c.Bind(&post); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	// Use repository to create post
	if err := postRepository.Create(&post); err != nil {
		c.JSON(400, gin.H{"error": "Failed to create post"})
		return
	}
	c.JSON(200, gin.H{"post": post})
}

func ReadAll(c *gin.Context) {
	// get all posts
	var posts []models.Post
	if err := postRepository.FindAll(&posts); err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve posts"})
		return
	}
	c.JSON(200, gin.H{"posts": posts})
}

func ReadOne(c *gin.Context) {
	// get id off url
	id := c.Param("id")
	// get the post of particular id
	var post models.Post
	if err := postRepository.FindByID(id, &post); err != nil {
		c.JSON(404, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(200, gin.H{"post": post})
}

func Update(c *gin.Context) {
	// get id
	id := c.Param("id")
	// get the existing post
	var post models.Post
	if err := postRepository.FindByID(id, &post); err != nil {
		c.JSON(404, gin.H{"error": "Post not found"})
		return
	}
	// update the post with request body data
	if err := c.Bind(&post); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	// use repository to update the post
	if err := postRepository.Update(&post); err != nil {
		c.JSON(400, gin.H{"error": "Failed to update post"})
		return
	}
	c.JSON(200, gin.H{"post": post})
}

func Deletes(c *gin.Context) {
	// get id
	id := c.Param("id")
	// use repository to delete the post
	if err := postRepository.Delete(id); err != nil {
		c.JSON(400, gin.H{"error": "Failed to delete post"})
		return
	}
	c.Status(200)
}
