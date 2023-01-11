package controllers

import (
	"JSON_CRUD_API_With_Gin_and_GORM/Initializers"
	"JSON_CRUD_API_With_Gin_and_GORM/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	//Get data off req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	//Create a post
	posts := models.Post{Title: body.Title, Body: body.Body}
	result := Initializers.DB.Create(&posts)
	if result.Error != nil {

		c.Status(400)
		return
	}

	//Return it
	c.JSON(200, gin.H{
		"Post": posts,
	})
}

func PostsIndex(c *gin.Context) {

	//Get the posts
	var post []models.Post
	Initializers.DB.Find(&post)

	//Respond with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsShow(c *gin.Context) {

	//Get id off url
	id := c.Param("id")

	// Get the posts
	var post models.Post
	Initializers.DB.First(&post, id)

	// Respond with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {

	// Get the id off the url
	id := c.Param("id")

	// Get the data off req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	// Find the post we're updating
	var post models.Post
	Initializers.DB.First(&post, id)

	// Update it
	Initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// Respond with it
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostsDelete(c *gin.Context) {

	// Get the id off the url
	id := c.Param("id")

	// Delete the posts
	Initializers.DB.Delete(&models.Post{}, id)

	// Respond
	c.Status(200)

}
