package controllers

import (
	"github.com/alighamgosar/golang_test_rest_api/initializers"
	"github.com/alighamgosar/golang_test_rest_api/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	// Create
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {

	// Get posts
	var posts []models.Post

	initializers.DB.Find(&posts)
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {

	id := c.Param("id")
	var post models.Post

	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {

	// id of url
	id := c.Param("id")

	// get the data off struct
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	// find the post we're updating
	var post models.Post
	initializers.DB.First(&post, id)

	// update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})
	// respond it
	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostsDelete(c *gin.Context) {

	// id of post
	id := c.Param("id")

	// delete it
	initializers.DB.Delete(&models.Post{}, id)

	// respond
	c.Status(200)

}
